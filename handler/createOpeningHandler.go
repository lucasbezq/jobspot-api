package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbezq/jobspot-api/schemas"
)

// @Tags Openings
// @Summary Create opening
// @Description Create a new job opening
// @BasePath /api/v1
// @Produce json
// @Accept json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 201 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
