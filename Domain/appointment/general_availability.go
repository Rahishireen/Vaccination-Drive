package appointment

import (
	"fmt"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)

//VaccinationAvailabilityCheck- Check whether the vaccination available for the given day
//30 Vaccines available in each Vaccination center per day
func (Bapp *BookAppointment) VaccinationAvailabilityCheck(query string)*apierrors.RestErr{

	//variable to store booked Vaccine count
	var booked_count int	

	//Check query for any errors
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	
	//Exec Query
	SelectResult,err :=stmt.Query(Bapp.BookingDate,Bapp.UserVC)	

	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked Vaccine count :%s",err.Error()))
	}
	
	//Check the count
	for SelectResult.Next(){
		SelectResult.Scan(&booked_count)
	}
	
	//Throw error if the count is greater than or equal to 30
	if booked_count >=30 {
		return apierrors.NewBadRequestError(fmt.Sprintf("All slots are booked for the date %s",Bapp.BookingDate))
	}
	return nil
	
}

//DoseAvailabilityCheck- Check whether the Dose available for the given day
//15 Vaccines available for First dose & 15 Vaccines available for second dose in each Vaccination center per day
func (Bapp *BookAppointment) DoseAvailabilityCheck(query string)*apierrors.RestErr{
	//Variable to store booked dose count
	var dose_count int

	//Check the query for any errors
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	SelectResult,err :=stmt.Query(Bapp.BookingDate,Bapp.DoseType,Bapp.UserVC)	

	
	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked dose count:%s",err.Error()))
	}
	
	//Check the count
	for SelectResult.Next(){
		SelectResult.Scan(&dose_count)
	}
	
	//Throw error if the count is greater than or equal 15
	if dose_count >=15 {
		return apierrors.NewInternalServerError(fmt.Sprintf("All slots are booked for the %d dose",Bapp.DoseType))
	}
	return nil
	
}

//TimeSlotAvailabilityCheck- Check whether the TimeSlot available for the given day
//10 users can register in each time per day
func (Bapp *BookAppointment) TimeslotAvailabilityCheck(query string)*apierrors.RestErr{
	//variable to store the time slot count
	var timeSlot_count int

	//Check the query for any errors
	stmt,err:=users_db.Client.Prepare(query)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	SelectResult,err :=stmt.Query(Bapp.BookingStatus,Bapp.TimeSlot,Bapp.UserVC)	

	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to get Booked TimeSlot count:%s",err.Error()))
	}
	
	//Check the count
	for SelectResult.Next(){
			SelectResult.Scan(&timeSlot_count)
	}
	
	//Throw error if the count is greater than or equal 10
	if timeSlot_count >=10 {
		return apierrors.NewInternalServerError(fmt.Sprintf("All slots are booked for the %s Time Slot",Bapp.TimeSlot))
	}
	return nil
	
}



