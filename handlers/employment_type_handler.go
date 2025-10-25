package handlers

import (
	"my-app/helper"
	"my-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEmploymentTypesHandler godoc
// @Summary Get all employment types
// @Description Retrieve a list of all employment types
// @Tags employment-types
// @Produce json
// @Success 200 {object} types.Response
// @Failure 500 {object} types.Response
// @Router /employment-types [get]
func GetEmploymentTypesHandler(c *gin.Context) {
	employmentTypes, err := services.GetAllEmploymentTypes()
	if err != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch employment types", nil)
		return
	}
	helper.JSONResponse(c, http.StatusOK, "success", "", employmentTypes)
}
