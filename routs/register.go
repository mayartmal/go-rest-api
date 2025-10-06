package routs

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	actorId := context.GetInt64("userId")
	targetEventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	targetEvent, err := models.GetEventById(targetEventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = targetEvent.Register(actorId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register"})
		return
	}

	fmt.Println("HERE")
	context.JSON(http.StatusOK, gin.H{"message": "Registered"})

}

func cancelRegistrationForEvent(context *gin.Context) {
	actorId := context.GetInt64("userId")
	targetEventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	targetEvent, err := models.GetEventById(targetEventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = targetEvent.CancelRegistration(actorId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel the registration"})
		return
	}

	fmt.Println("HERE")
	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled"})

}
