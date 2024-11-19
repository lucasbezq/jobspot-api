package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lucasbezq/jobspot-api/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":    msg,
		"errorCode:": code,
		"timestamp":  time.Now(),
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
