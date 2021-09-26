package controllers

import (
	"net/http"
	"vaccinationdrive/Domain/users"
	services "vaccinationdrive/Services"
	apierrors "vaccinationdrive/utills/error"

	"github.com/gin-gonic/gin"
)

//RegisterUser - - Valid JSON for its data type and route to Services-Register User
func RegisterUser(c *gin.Context) {

	//Create variable of User-Registration type
	var beneficiary users.UserRegistration

	//Validate JSON Body
	if err:=c.ShouldBindJSON(&beneficiary);err!=nil{
		resterr:=apierrors.NewBadRequestError("Invalid json Body")
		c.JSON(resterr.Code,resterr)
		return
	}

	//Call Service -To  register User
	result,saveErr:= services.RegisterUser(beneficiary)
	if saveErr!=nil{
		c.JSON(saveErr.Code,saveErr)
		return
	}	

	//Status
	c.JSON(http.StatusAccepted,result)

}