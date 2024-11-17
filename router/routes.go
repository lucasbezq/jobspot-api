package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasbezq/jobspot-api/handler"
)

func initializeRoutes(router *gin.Engine) {
	handler.IniializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.GET("/openings/opening", handler.FindOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
