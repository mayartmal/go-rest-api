package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routs"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()
	routs.RegisterRouts(server)
	server.Run(":8080") // localhost:8080
}
