package users

import (
	"fmt"
	"strings"
	"time"
	"vaccinationdrive/Datasources/mysql/users_db"
	apierrors "vaccinationdrive/utills/error"
)

const (

//Insert Query to insert user data into users_DB
	QueryInsetUser = "INSERT INTO users(Name,DOB,Aadhaar_Number,Phone_Number,Beneficiary_Id,Date_Created) VALUES (?,?,?,?,?,?);"

//Select Query to get the user data from the DB
	QueryGetId = "SELECT Name,DOB,Aadhaar_Number,Phone_Number,Beneficiary_Id FROM users WHERE Beneficiary_Id = ?"

//Variable to replace common SQL errors to reasonable messages
	errorNoRows  ="no rows in result set"
	Duplicate="Duplicate"
	
)

//Save - to insert user data into the DB
func (user *UserRegistration) Save() *apierrors.RestErr {

	//Check the query for any errors
	stmt,err:=users_db.Client.Prepare(QueryInsetUser)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Exec the query
	insertResult,err :=stmt.Exec(user.Name,user.DOB,user.AadhaarNumber,user.PhoneNumber,user.BeneficiaryID,time.Now())
	

	if err!=nil{
		//Replace SQL-Error Duplicate with reasonable Message
		if strings.Contains(err.Error(),Duplicate){			
			return apierrors.NewBadRequestError(fmt.Sprintf("Aadhaar Number %s already exists",user.AadhaarNumber))
	
		}
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to save user:%s",err.Error()))
	}

	//to get the Last Inserted details,if the insert query fails
	_,err= insertResult.LastInsertId()
	if err!=nil{
		return apierrors.NewInternalServerError(fmt.Sprintf("error when trying to save user:%s",err.Error()))
	}

	return nil
}

//Get - To get the user data from the DB
func (user *UserRegistration) Get() *apierrors.RestErr{

	//Check the query for any errors
	stmt,err:=users_db.Client.Prepare(QueryGetId)
	if err!=nil{
		return apierrors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	//Query Single Row
	result:=stmt.QueryRow(user.BeneficiaryID)
	if err:=result.Scan(&user.Name,&user.DOB,&user.AadhaarNumber,&user.PhoneNumber,&user.BeneficiaryID);err!=nil{
		//Replace SQL-Error "no rows in result set" with reasonable Message
		if strings.Contains(err.Error(),errorNoRows){
			return apierrors.NewNotFoundError(fmt.Sprintf("Beneficiary Id %s not registered",user.BeneficiaryID))
		}
		return apierrors.NewInternalServerError(
			fmt.Sprintf("error When trying to get user with Beneficiary Id %s:%s",user.BeneficiaryID,err.Error()))
		}

	return nil

}