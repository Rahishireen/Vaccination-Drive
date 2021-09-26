package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping to check the connection to send http request
func Ping(c *gin.Context){
	c.String(http.StatusOK,"pong")
}