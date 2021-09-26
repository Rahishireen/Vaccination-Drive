package services

import (
	"vaccinationdrive/Domain/appointment"
	apierrors "vaccinationdrive/utills/error"
)

//Appointment - Services to check Basic Field Validation,General Availability Validation,UserEligibiliy Validation and
//then Add/Update/Delete the Booking details based on the user Input
func Appointment(Bapp appointment.BookAppointment) (*appointment.BookAppointment,*apierrors.RestErr){

	//Validate all the input fields for length and Format
	if FieldErr:=Bapp.BasicFieldValidation();FieldErr!=nil{
		return nil,FieldErr
	}

	//Booking Status
	// 1 - Add/Book
	// 2 - Update/Reschedule
	// 3 - Delete/Cancel
	switch Bapp.BookingStatus{
	// 1 - Add/Book
	case 1:

		//Check the availability of vaccine,dose,timeslot
		if generalErr:=Bapp.GeneralAvailabityValidation();generalErr!=nil{
		return nil,generalErr
		}

		//Check the user eligibility based on the DB data
		if beneficiaryErr:=Bapp.UserEligibleValidationAdd();beneficiaryErr!=nil{
		return nil,beneficiaryErr
		}

		//Add Beneficiary to the Booking_Appointment DB
		if DBAdderr := Bapp.AddBeneficiary();DBAdderr!=nil{
			return nil,DBAdderr
		}
	    return &Bapp, nil
	// 2 - Update/Reschedule
	case 2:

		//Check the availability of vaccine,dose,timeslot
		if generalErr:=Bapp.GeneralAvailabityValidation();generalErr!=nil{
			return nil,generalErr
			}

		//Check the user eligibility based on the DB data
		if beneficiaryErr:=Bapp.UserEligibleValidationUpdate();beneficiaryErr!=nil{
				return nil,beneficiaryErr		
			}

		//Update Booking date or Time Slot
		if DBUpdateerr := Bapp.UpdateBeneficiary();DBUpdateerr!=nil{
				return nil,DBUpdateerr
			}
			return &Bapp, nil
	// 3 - Delete/Cancel	
	case 3:

		//Check the user eligibility based on the DB data
		if beneficiaryErr:=Bapp.UserEligibleValidationDelete();beneficiaryErr!=nil{
			return nil,beneficiaryErr
			}

		//Delete Beneficiary Data from the DB
		if DBUpdateerr := Bapp.DeleteBeneficiary();DBUpdateerr!=nil{
			return nil,DBUpdateerr
		}
		return &Bapp, nil

	}

	return nil,apierrors.NewInternalServerError("Please provide valid Action,It Should be 1 (Book), 2 (Reschedule), 3 (Cancel)")
}