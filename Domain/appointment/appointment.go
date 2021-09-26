package appointment

//type BookAppointment defined with mandatory fields for booking an appointment
type BookAppointment struct {
	BeneficiaryID string `json:"beneficiaryid"`
	BookingDate   string `json:"bookingdate"`
	TimeSlot      string `json:"timeslot"`
	DoseType      int    `json:"dose"`
	UserVC        string `json:"vaccinationcenter"`
	BookingStatus int    `json:"action - 1.Book 2.Reschedule 3.Cancel"`
}
