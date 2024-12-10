package main

import (
	"REST-API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST,PUT,PATCH,DELETE
	server.POST("/events", createEvents)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context) {

	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}
func createEvents(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Validating error"})
		return
	}
	event.ID = 1
	event.UserID = 1

	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
