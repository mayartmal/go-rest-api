package routs

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
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
	userId := context.GetInt64("userId")
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request (event) data"})
		return
	}

	event.UserID = userId
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
	fmt.Println(eventId)
	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB reading"})
		fmt.Println(err)
		return
	}
	var event models.Event
	event.ID = eventId
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

	context.JSON(http.StatusOK, gin.H{"message": "Event Updated"})
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

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request (user) data"})
		return
	}

	// user.ID = 2
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error during DB writing"})
		fmt.Println(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Created", "user": user})
	print(err)
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request (user) data"})
		return
	}
	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Error during login"})
		fmt.Println(err)
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error during token generation"})
		fmt.Println(err)
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}
