package appointment

import (
	"fmt"
	"time"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)
const (
//Insert query to insert all the Booking details into DB
	QueryInsertUser = "INSERT INTO booking_appointment(Beneficiary_Id,Booking_Date,Time_Slot,Dose_type,Vaccination_Center,Date_Created) VALUES (?,?,?,?,?,?);"
	
//Update Query to update time-Slot and Booking date	in the DB
	QueryUpdateUser ="UPDATE booking_appointment set Booking_Date = ? Time_Slot = ? WHERE Beneficiary_Id = ?  AND Dose_Type = ? AND Vaccination_Center =? ;"

//Delete Beneficiary details from the DB
	QueryDeleteUser ="DELETE FROM booking_appointment WHERE Beneficiary_Id = ?  AND Booking_Date = ? and Time_Slot = ? AND Dose_Type = ? AND Vaccination_Center =? "
)

//AddBeneficiary-Add data to DB
func (Bapp *BookAppointment) AddBeneficiary() *apierrors.RestErr {

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(QueryInsertUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	insertResult, err := stmt.Exec(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC,time.Now())

	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to Add user in Booking_Appointment DB:%s", err.Error()))
	}

	//to get the Last Inserted details,if the insert query fails
	_, err = insertResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to Add user in Booking_Appointment DB:%s", err.Error()))
	}

	return nil
}

//UpdateBeneficiary-Update DB
func (Bapp *BookAppointment) UpdateBeneficiary() *apierrors.RestErr {

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(QueryUpdateUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	updateResult, err := stmt.Exec(Bapp.BookingDate,Bapp.TimeSlot,Bapp.BeneficiaryID,
		Bapp.DoseType,Bapp.UserVC)

	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to update user in Booking_Appointment DB:%s", err.Error()))
	}

	//to get the Last updated details,if the update query fails
	_, err = updateResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to update user in Booking_Appointment DB:%s", err.Error()))
	}

	return nil
}

//DeleteBeneficiary-Delete data from DB
func (Bapp *BookAppointment) DeleteBeneficiary() *apierrors.RestErr {

	//Check the query for any errors
	stmt, err := users_db.Client.Prepare(QueryDeleteUser)
	if err != nil {
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	deleteResult, err := stmt.Exec(Bapp.BeneficiaryID,Bapp.BookingDate,Bapp.TimeSlot,Bapp.DoseType,Bapp.UserVC)
	

	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to delete user in Booking_Appointment DB:%s", err.Error()))
	}

	//to get the Last deleted data details,if the delete query fails
	_, err = deleteResult.LastInsertId()
	if err != nil {
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to delete user in Booking_Appointment DB:%s", err.Error()))
	}
	return nil
}