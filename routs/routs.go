package routs

import "github.com/gin-gonic/gin"

func RegisterRouts(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)

}
