package app

import controllers "vaccinationdrive/Controllers"

//mapUrls - Map different requests into its corresponding handler functions
func mapUrls() {

	//Ping to check the connection
	router.GET("/ping", controllers.Ping)

	//user_registration - Route to User Registration controller
	router.POST("/user_registration",controllers.RegisterUser)

	//booking_appointment - Route to Booking Appointment controller
	router.GET("/booking_appointment",controllers.BookAppointment)
}