package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST opennings",
	})
}

func FindOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opennings",
	})
}

func ListOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opennings",
	})
}

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opennings",
	})
}

func UpdateOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT opennings",
	})
}
