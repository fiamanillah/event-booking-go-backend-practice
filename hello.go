package main

import (
	"gin/db"
	"gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	server.Static("/static", "./static")

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	routes.RegisterRoutes(server)

	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.Run(":1090")
}
