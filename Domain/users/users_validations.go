package users

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"vaccinationdrive/utills/age"
	"vaccinationdrive/utills/date"
	apierrors "vaccinationdrive/utills/error"
)

//CalculateAge - To calculate the age for the user DOB
func (u *UserRegistration) CalculateAge() int {

	//Split string Date into day,month,year of Int type
	day, month, year := splitDate(u.DOB)

	//convert string into time format
	dob := getDOB(year, month, day)

	//calculate age with DOB
	return age.Age(dob)

}

//Convert string into time format
func getDOB(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}

//Split string date into day,month,year
func splitDate(date string) (int, int, int) {

	//Split the given string by "-"
	datestr := strings.Split(date, "-")

	//convert the splited string date into int
	day, _ := strconv.Atoi(datestr[0])
	month, _ := strconv.Atoi(datestr[1])
	year, _ := strconv.Atoi(datestr[2])

	return day, month, year

}

//FieldValidation - Check whether the input fields satisfies the given length and Pattern
func FieldValidation(length int, RegEx string, field string, fieldName string, Pattern string) *apierrors.RestErr {

	//Compare the input field length with the given length
	if len(field) != length {
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Length should be %d", fieldName, length))

	}

	//Compare the input field pattern with the given pattern
	inputPattern := regexp.MustCompile(RegEx)
	if !inputPattern.MatchString(field) {
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Format should be %s", fieldName, Pattern))

	}
	return nil
}

//AgeValidation - Check whether the user is 45+
//users with 45+ of age are allowed to register for Vaccinaton
func (u *UserRegistration) AgeValidation(inputAge int) *apierrors.RestErr {

	//function to calculate Age
	age := u.CalculateAge()

	//check the age is <= inputage
	if age <= inputAge {
		return apierrors.NewBadRequestError(fmt.Sprintf("Member (Age - %d) not eligible - Age Should be greater than 45", age))

	}
	return nil
}

//CreateBeneficiaryID - Create 14 digit unique ID based on the current Time Stamp
func CreateBeneficiaryId() string {
	var beneficiaryid string

	//get the current time stamp in string format
	currenttime := date.GetTimeString()

	//Remove special characters in the TimeStamp
	str1 := strings.ReplaceAll(currenttime, "-", "")
	str2 := strings.ReplaceAll(str1, " ", "")
	beneficiaryid = strings.ReplaceAll(str2, ":", "")
	
	return beneficiaryid
}