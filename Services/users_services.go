package services

import (
	"strings"
	users "vaccinationdrive/Domain/users"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

//RegisterUser - Check Basic field Validation and Add the user into DB
func RegisterUser(user users.UserRegistration) (*users.UserRegistration,*apierrors.RestErr) {

	//Name Validation- Should be greater than spaces
	if len(strings.TrimSpace(user.Name))==0{
		return nil,apierrors.NewBadRequestError("Name Should be valid value")

	}

	//Date Validation- Check field Length and Format
	if dateerr := users.FieldValidation(10,`\d{2}-\d{2}-\d{4}`,user.DOB,"Date","DD-MM-YYYY"); dateerr != nil {
		return nil, dateerr
	}

	//Aadhar Validation- Check field Length and Format
	if aadharerr:=users.FieldValidation(12,`\d{12}`,user.AadhaarNumber,"Aadhaar Number","12 Digit");aadharerr!=nil{
		return nil,aadharerr
	}

	//PhoneNumber Validation- Check field Length and Format
	if phoneNumErr:=users.FieldValidation(10,`\d{10}`,user.PhoneNumber,"Phone Number","10 Digit");phoneNumErr!=nil{
		return nil,phoneNumErr
	}

	//age Should be 45+
	age:=45
	if ageErr:= user.AgeValidation(age);ageErr!=nil{
		return nil,ageErr
	}

	//Create random unique beneficiary ID based on the current Time Stamp
	user.BeneficiaryID=users.CreateBeneficiaryId()

	//Get the current date to store in DB
	user.DateCreated=date.GetTimeString()

	//Insert the user details into the DB
	if DBerr := user.Save();DBerr!=nil{
		return nil,DBerr
	}
	
	return &user, nil
}