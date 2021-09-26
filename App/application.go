package app

import (
	"github.com/gin-gonic/gin"
)

//Creating Router to handle different http Rquests
var (
	router = gin.Default()
)

//Application - Start listening on port 8080
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}