package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.POST("/events", createEvent)
	router.GET("/events/:id", getEvent)
}
