package controllers

import "github.com/gin-gonic/gin"

func GetEvents(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetEvents",
	})
}

func CreateEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "CreateEvent",
	})
}

func GetEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetEvent",
	})
}

func UpdateEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "UpdateEvent",
	})
}

func DeleteEvent(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DeleteEvent",
	})
}

func EventsController(r *gin.Engine) {
	events := r.Group("/events")
	{
		events.GET("/", GetEvents)
		events.POST("/", CreateEvent)
		events.GET("/:id", GetEvent)
		events.PUT("/:id", UpdateEvent)
		events.DELETE("/:id", DeleteEvent)
	}
}
