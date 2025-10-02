package routs

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB reading"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusOK, *event)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB reading"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request (event) data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB writing"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB reading"})
		fmt.Println(err)
		return
	}
	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request (event) data"})
		return
	}
	err = event.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB updating"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated", "event": event})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB reading"})
		fmt.Println(err)
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during event deletion from DB"})
		fmt.Println(err)
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})

}
