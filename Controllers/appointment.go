package controllers

import (
	"net/http"
	"vaccinationdrive/Domain/appointment"
	services "vaccinationdrive/Services"
	apierrors "vaccinationdrive/utills/error"

	"github.com/gin-gonic/gin"
)

//BookAppointment - Valid JSON for its data type and route to Services-Book Appointment
func BookAppointment(c *gin.Context){
	//Create variable of BookingAppointment type
	var beneficiary appointment.BookAppointment
	if err:=c.ShouldBindJSON(&beneficiary);err!=nil{
		resterr:=apierrors.NewBadRequestError("Invalid json Body")
		c.JSON(resterr.Code,resterr)
		return
	}
	result,saveErr:= services.Appointment(beneficiary)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	

	c.JSON(http.StatusAccepted,result)

}