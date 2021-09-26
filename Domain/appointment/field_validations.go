package appointment

import (
	"fmt"
	"regexp"
	apierrors "vaccinationdrive/utills/error"
)
var(

//Available Vaccination Centers
AvailableVaccinationCenter = []string{"Nungambakkam", "Tambaram", "Velachery", "Shozhinganallur"}

//Available Time Slots
AvailableTimeSlot          = []string{"9.30 am to 11.30 am","2pm to 4pm","6pm to 8pm"}

//Available Dose Types | 1 - First dose , 2 - Second dose
AvailableDoseType          =[]int {1,2}

)

//vaccineCenterValidation - Check whether the booking Vaccine Center present in the available vaccine center or not
func (a *BookAppointment) VaccineCenterValidation() *apierrors.RestErr{

	//Loop to check the booking VC in available VC array
	for _, value := range AvailableVaccinationCenter {
		if a.UserVC == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf("Vaccination is not available %s,Please choose between 1.Nungambakkam 2.Tambaram 3.Velachery 4.Shozhinganallur",a.UserVC))

}

//TimeSlotValidation - Check whether the booking time slot present in the available vaccine center or not
func (a *BookAppointment) TimeSlotValidation() *apierrors.RestErr{

	//Loop to check the booking timeslot in available timeslot array
	for _, value := range AvailableTimeSlot {
		if a.TimeSlot == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf("Selected Slot is not available for the time %s,Please choose between 9.30 am to 11.30 am /2pm to 4pm /6pm to 8pm",a.TimeSlot))

}

//DoseTypeValidation - Check whether the booking Dose type present in the available Dose Type or Not
func (a *BookAppointment) DoseTypeValidation() *apierrors.RestErr{

	//Loop to check the booking dosetype in available dosetype array
	for _, value := range AvailableDoseType {
		if a.DoseType == value {
			return nil
		}
	}

	return apierrors.NewBadRequestError(
		fmt.Sprintf ("Selected % d Dose is not applicable,Please choose between 1 (first Dose) / 2 (second Dose)",a.DoseType))

}

//FieldValidation - Check whether the input fields satisfies the given length and Pattern
func FieldValidation(length int,RegEx string,field string,fieldName string,Pattern string) *apierrors.RestErr{
	
	//Compare the input field length with the given length
	if len(field)!=length{
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Length should be %d",fieldName,length))

	}

	//compare the input field pattern with the given pattern
	inputPattern:=regexp.MustCompile(RegEx)
	if !inputPattern.MatchString(field){
		return apierrors.NewBadRequestError(fmt.Sprintf("Invalid %s - Format should be %s",fieldName,Pattern))

	}
	return nil
}