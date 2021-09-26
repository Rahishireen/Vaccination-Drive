package users

//Type UserRegistration Defined with all the mandatory fields to register a user
type UserRegistration struct {
	Name          string `json:"name"`
	DOB           string `json:"dob"`
	AadhaarNumber string  `json:"aadharnumber"`
	PhoneNumber   string  `json:"phonenumber"`
	BeneficiaryID string  `json:"beneficiaryid"`
	DateCreated    string `json:"datecreated"`
}


