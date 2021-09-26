package appointment

import (
	"time"
	"vaccinationdrive/Domain/users"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

//Query's to handle different database Operations
const (
//SELECT Queries
	//To get the count of No of Booked users for the given date in the given Vaccination Center
	QuerySelectVaccination = "SELECT count(*) as No_Booked_users FROM booking_appointment WHERE booking_date in (?) AND Vaccination_Center in (?);"

	//To get the count of Booked Dose for the given date in the given Vaccination Center
	QuerySelectDose        = "SELECT count(*) as Dose_Count FROM booking_appointment WHERE booking_date in (?) AND Dose_Type in (?) AND Vaccination_Center in (?);"

	//To get the Count of Booked Slots for the given date and time slot in the given Vaccination Center
	QuerySelectTimeSlot    = "SELECT count(*) as Time_Slot FROM booking_appointment WHERE booking_date in (?) AND Time_Slot in (?) AND Vaccination_Center in (?);"
	
	//To get the No of Rows present in the DB for the given Beneficiary_Id for ADD operation
	QueryBeneficiaryCountAdd  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?)"

	//To get the No of Rows present in the DB for the given Beneficiary_Id for UPDATE operation
	QueryBeneficiaryCountUpdate  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?) AND  Time_Slot in (?) AND Dose_Type in (?) AND Vaccination_Center in (?)"
	
	//To get the No of Rows present in the DB for the given Beneficiary_Id for DELETE operation
	QueryBeneficiaryCountDelete  = "SELECT count(*) as Beneficiary_Count FROM booking_appointment WHERE Beneficiary_Id in (?) AND booking_date in (?) AND  Time_Slot in (?) AND Dose_Type in (?) AND Vaccination_Center in (?)"
	
	//To get the Dose Type,Booking date from DB for the given Beneficiary_Id
	QueryGetId             ="SELECT Dose_Type,Booking_date FROM booking_appointment WHERE Beneficiary_Id = ?"

//Variables to replace the common SQL errors with more reasonable Error Messages 
	errorNoRows            = "no rows in result set"
	Duplicate              = "Duplicate"
)


//BasicFieldValidation - Check for the Field Length,Patterns,Date Validations
func (Bapp *BookAppointment) BasicFieldValidation() *apierrors.RestErr{

	//Create Variable of Type UserRegistration
	var user users.UserRegistration

	//Map UserRegistration-Beneficiary Id to Booking Appointment Beneficiary to get Data from UsersRegistration DB 
	user.BeneficiaryID=Bapp.BeneficiaryID

//Beneficiary Id Validation- Check for Beneficiary Id in Users DB,Check the pattern
	//1.Pattern Validation
	if beneficiaryErr := FieldValidation(14,`\d{14}`,Bapp.BeneficiaryID,"Beneficiary Id","14 digit"); beneficiaryErr != nil {
		return beneficiaryErr
	}
	//2.Check Entry in Users DB
	if beneficiaryEntryErr:= user.Get();beneficiaryEntryErr!=nil{
		return beneficiaryEntryErr
	}

//Vaccination Center Validation
	if vcErr:=Bapp.VaccineCenterValidation();vcErr!=nil{
		return vcErr
	}

//Booking_date Validation   
	//1.Check Booking Date > current Date

	
	//2.Pattern Validation
	if dateErr := FieldValidation(10,`\d{2}-\d{2}-\d{4}`,Bapp.BookingDate,"Booking Date","DD-MM-YYYY"); dateErr != nil {
		return dateErr
	}
	//3.No of days gap Check - Should be greater than or equal to 90 days

	//No need to check the below validation for Delete Operation
	if Bapp.BookingStatus!=3{
	gap := date.CheckTimeGap(date.ConvertStringtoTime(Bapp.BookingDate),time.Now())
		if gap < 90 {
			return (apierrors.NewBadRequestError("Slot can't be booked before 90 days"))
		}
	}

//TimeSlot Validation
	if tsErr:=Bapp.TimeSlotValidation();tsErr!=nil{
		return tsErr
	}

//DoseType Validation
	if dtErr:=Bapp.DoseTypeValidation();dtErr!=nil{
		return dtErr
	}

	return nil
}


//GeneralAvailability Validation - Check the Availability of Vaccines,Dose,Time-Slot
func (Bapp *BookAppointment) GeneralAvailabityValidation() *apierrors.RestErr {

	//Availability of Vaccination in the Booked Center on the Booked Date
	//Vaccines - 30 vaccines available in Each Vaccine Center Per Day
	if vaccinationErr := Bapp.VaccinationAvailabilityCheck(QuerySelectVaccination); vaccinationErr != nil {
		return vaccinationErr
	}

	//Availability of Dose in the Booked Center on the Booked Date
	//Dose- In each vaccine center 15 available for First Dose,15 available for Second Dose
	if doseErr:=Bapp.DoseAvailabilityCheck(QuerySelectDose);doseErr!=nil{
		return doseErr
	}	

	//Availability of Vaccination in the booked time slot on the booked date,booked vaccination center
	//Time-Slot-Each Time Slot allow only 10 users to register
	if timeSlotErr:=Bapp.TimeslotAvailabilityCheck(QuerySelectTimeSlot);timeSlotErr!=nil{
		return timeSlotErr
	}

	return nil
}


//userEligibleValidationAdd - Check the users is eligible to Book Appointment or not
func (Bapp *BookAppointment) UserEligibleValidationAdd() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBAdd(QueryBeneficiaryCountAdd)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}

	switch beneficiarycount {
	//No Entries in Database
	case 0:
		if Bapp.DoseType!=1{			
			return apierrors.NewBadRequestError("Second Dose Should be taken after First Dose")
		}
	//1 Entry in Database
	case 1:
		  if err:= Bapp.GetVaccineDetailsInDB(QueryGetId);err!=nil{
			  return err
		  }
	//2 Entries in database
	case 2:
		return apierrors.NewBadRequestError("Booking limit Exhausted,Users allowed to book upto 2 appointments")
	//No than 2 Entries- invalid Scenario
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	return nil
}

//userEligibleValidationUpdate - Check the users is eligible to Book Appointment or not
func (Bapp *BookAppointment) UserEligibleValidationUpdate() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBUpdate(QueryBeneficiaryCountUpdate)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}

	switch beneficiarycount {
	//No Entries in Data Base
	case 0:		
			return apierrors.NewNotFoundError("No entry for the given data")	
	//1 Entry in DataBase
	case 1:
		  return nil
	//2 Entries in Data Base
	case 2:
		if err:= Bapp.GetVaccineDetailsInDB(QueryGetId);err!=nil{
			return err
		}
	//More than 2 Entries - Invalid Scenario
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	return nil
}

//userEligibleValidationDelete - Check the users is eligible to Book Appointment or not
func (Bapp *BookAppointment) UserEligibleValidationDelete() *apierrors.RestErr{

	//Check for no of entries of Beneficiary id in Booking Appointment dataBase
	beneficiarycount,beneficiaryErr:=Bapp.CheckBeneficiaryInDBDelete(QueryBeneficiaryCountDelete)
	if beneficiaryErr!=nil{
		return beneficiaryErr
	}
	switch beneficiarycount {
	//No Entry in Data Base
	case 0:		
			return apierrors.NewNotFoundError("No entry for the given data")	
	//1 Entry in Data Base	
	case 1:
		  return nil
	//2 Entries in Data Base
	case 2:
			return nil
	//More than 2 entries in Data Base,Invalid Scenario
	default:
		return apierrors.NewInternalServerError("DB contains More than 2 entries,which is Invalid")		
	}
	//return nil
}