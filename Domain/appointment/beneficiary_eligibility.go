package appointment

import (
	"fmt"
	"vaccinationdrive/Datasources/mysql/users_db"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

//CheckBeneficiaryInDBAdd - Check the no of entries of Beneficiary Id in the DataBase
func (Bapp *BookAppointment) CheckBeneficiaryInDBAdd(query string) (int,*apierrors.RestErr) {

	//Variable to store the count
	var beneficiary_Entry int

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Execute the query,which will give the pointer to first row
	SelectResult, err := stmt.Query(Bapp.BeneficiaryID)
	
	//Error-Handling
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Add:%s", err.Error()))
	} else {

	//Loop to get the table data with the pointer
	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}
	fmt.Println("Beneficiary Count",beneficiary_Entry)
	return beneficiary_Entry,nil

}
}

//CheckBeneficiaryInDBUpdate - Check the no of entries of Beneficiary Id in the DataBase
func (Bapp *BookAppointment) CheckBeneficiaryInDBUpdate(query string) (int,*apierrors.RestErr) {

	//Variable to store the count
	var beneficiary_Entry int

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Execute the query,which will give the pointer to first row
	SelectResult, err := stmt.Query(Bapp.BeneficiaryID,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	
	//Error-Handling
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Update:%s", err.Error()))
	} else {

	//Loop to get the table data with the pointer
	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}

	return beneficiary_Entry,nil

}
}

//CheckBeneficiaryInDBDelete - Check the no of entries of Beneficiary Id in the DataBase
func (Bapp *BookAppointment) CheckBeneficiaryInDBDelete(query string) (int,*apierrors.RestErr) {

	//Variable to store the count
	var beneficiary_Entry int

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return 0,apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Execute the query,which will give the pointer to first row
	SelectResult, err := stmt.Query(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	
	if err != nil {
		return 0,apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Beneficiary Count in DB for Delete:%s", err.Error()))
	} 
	
	//Loop to get the table data with the pointer
	for SelectResult.Next() {
		SelectResult.Scan(&beneficiary_Entry)
		
	}

	return beneficiary_Entry,nil

}


//GetVaccineDetailsInDB - to get the Beneficiary Vaccination history in DB
func (Bapp *BookAppointment) GetVaccineDetailsInDB(query string) (*apierrors.RestErr) {
	//variable to store DB data
	var beneficiary_DB_Data BookAppointment

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(query)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Execute the query,which will give the pointer to first row
	SelectResult:= stmt.QueryRow(Bapp.BeneficiaryID)

	if err:=SelectResult.Scan(&beneficiary_DB_Data.DoseType,&beneficiary_DB_Data.BookingDate);err!=nil{		
		return apierrors.NewInternalServerError(
			fmt.Sprintf("error When trying to get user with Beneficiary Id %s:%s",Bapp.BeneficiaryID,err.Error()))
		}
	//Check the booking vaccine dose is eligible for the beneficiary
	if doseErr:=Bapp.ValidateDoseCheck(beneficiary_DB_Data.DoseType,beneficiary_DB_Data.BookingDate);doseErr!=nil{
		return doseErr
	}	
	return nil
}

//ValidateDoseCheck - check whether the booking  vaccine dose is eligible or not based on the DB vaccine dose 
//for the given beneficiary Id
func (b *BookAppointment) ValidateDoseCheck(dosetype int, dbDate string) *apierrors.RestErr {

	//Booking dosetype is ineligible when it is already present in the DB
	if b.DoseType == dosetype {
		return apierrors.NewInternalServerError("Vaccination already taken for the given dose type")

	} else if b.DoseType > dosetype {
		//no of dyas between First dose and Second dose should be 15 days
		gap := date.CheckTimeGap(date.ConvertStringtoTime(b.BookingDate),date.ConvertStringtoTime(dbDate))
		fmt.Println("No of gap Days ",gap)
		if gap < 15 {
			return (apierrors.NewBadRequestError("Second Dose Should be taken after 15 days"))
		}
	}

	return nil
}