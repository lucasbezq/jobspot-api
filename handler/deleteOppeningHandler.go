package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOppeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opennings",
	})
}
