package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/lucasbezq/jobspot-api/docs"
	"github.com/lucasbezq/jobspot-api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.IniializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.GET("/openings/opening", handler.FindOpeningHandler)
		v1.PUT("/openings", handler.UpdateOpeningHandler)
		v1.DELETE("/openings", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)
	}

	router.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))
}
