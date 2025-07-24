package main

import (
	"gin/db"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./static/favicon.ico")
	})

	router.GET("/events", getEvents)
	router.POST("/events", createEvent)

	router.Run(":8080")
}
func getEvents(context *gin.Context) {
	events, _ := models.GetEvents()

	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})

}

func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		event.Save()
		context.JSON(http.StatusCreated, event)
	}
}
