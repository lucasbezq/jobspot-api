package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbezq/jobspot-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/opennings", handler.CreateOppeningHandler)
		v1.GET("/opennings/oppening", handler.FindOppeningHandler)
		v1.PUT("/opennings", handler.UpdateOppeningHandler)
		v1.DELETE("/opennings", handler.DeleteOppeningHandler)
		v1.GET("/opennings", handler.ListOppeningHandler)
	}
}
