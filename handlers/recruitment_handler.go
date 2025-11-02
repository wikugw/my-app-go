package handlers

import (
	"fmt"
	"net/http"
	"time"

	"my-app/helper"
	"my-app/services"
	handlerTypes "my-app/types/handlers"

	"github.com/araddon/dateparse"
	"github.com/gin-gonic/gin"
)

// CreateRecruitmentHandler godoc
// @Summary Create a new recruitment
// @Description Add a new recruitment posting with requirements
// @Tags recruitment
// @Accept json
// @Produce json
// @Param recruitment body handlerTypes.CreateRecruitmentRequest true "Recruitment data"
// @Success 201 {object} types.Response
// @Failure 400 {object} types.Response
// @Failure 500 {object} types.Response
// @Router /recruitments [post]
func CreateRecruitmentHandler(c *gin.Context) {
	var req handlerTypes.CreateRecruitmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.JSONResponse(c, http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	fmt.Printf("📥 Incoming request: %+v\n", req)

	// Konversi dan validasi ID
	deptID, err := helper.StringToUint(req.Position)
	if err != nil {
		fmt.Println("❌ Invalid position ID:", err)
		helper.JSONResponse(c, http.StatusBadRequest, "error", "Invalid department ID", nil)
		return
	}

	employmentTypeID, err := helper.StringToUint(req.EmploymentType)
	if err != nil {
		fmt.Println("❌ Invalid employment type ID:", err)
		helper.JSONResponse(c, http.StatusBadRequest, "error", "Invalid employment type ID", nil)
		return
	}

	// Parse tanggal aplikasi (jika ada)
	var startDate, endDate *time.Time
	if len(req.ApplicationDates) == 2 {
		fmt.Println("🕓 Raw application dates:", req.ApplicationDates)

		start, err1 := time.Parse(time.RFC3339, req.ApplicationDates[0])
		end, err2 := time.Parse(time.RFC3339, req.ApplicationDates[1])

		fmt.Println("Parsed start:", start, "err:", err1)
		fmt.Println("Parsed end:", end, "err:", err2)

		if err1 == nil {
			startDate = &start
		}
		if err2 == nil {
			endDate = &end
		}
	} else {
		fmt.Println("⚠️ ApplicationDates length:", len(req.ApplicationDates))
	}

	fmt.Println("✅ startDate:", startDate)
	fmt.Println("✅ endDate:", endDate)

	recruitment := handlerTypes.RecruitmentParam{
		DepartmentID:         deptID,
		Salary:               req.Salary,
		EmploymentTypeID:     employmentTypeID,
		ApplicationStartDate: startDate,
		ApplicationEndDate:   endDate,
		CreatedByID:          req.CreatedByID,
		Requirements:         req.Requirements,
	}

	if err := services.CreateRecruitment(&recruitment); err != nil {
		fmt.Println("❌ Service error:", err)
		helper.JSONResponse(c, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	fmt.Println("✅ Recruitment created successfully:", recruitment)
	helper.JSONResponse(c, http.StatusCreated, "success", "Recruitment created successfully", recruitment)
}

// GetActiveRecruitmentsHandler godoc
// @Summary      Get active recruitments by current date
// @Description  Returns a list of active recruitments where current date is between application_start_date and application_end_date.
// @Tags         Recruitments
// @Accept       json
// @Produce      json
// @Param        currentDate   query     string  true  "Current date in RFC3339 format (e.g. 2025-10-28T00:00:00Z)"
// @Success      200  {object}  types.Response{data=[]handlers.RecruitmentResponse}  "List of active recruitments"
// @Failure      400  {object}  types.Response  "Missing or invalid currentDate parameter"
// @Failure      500  {object}  types.Response  "Internal server error"
// @Router       /recruitments/active [get]
func GetActiveRecruitmentsHandler(c *gin.Context) {
	dateRequest := c.Query("currentDate")
	if dateRequest == "" {
		helper.JSONResponse(c, http.StatusBadRequest, "error", "currentDate query param is required", nil)
		return
	}

	currentDate, dateError := dateparse.ParseAny(dateRequest)

	if dateError != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch recruitments", dateError)
		return
	}

	recruitments, err := services.GetActiveRecruitments(currentDate)
	if err != nil {
		helper.JSONResponse(c, http.StatusInternalServerError, "error", "Failed to fetch recruitments", err)
		return
	}
	helper.JSONResponse(c, http.StatusOK, "success", "", recruitments)
}

// GetRecruitmentById godoc
// @Summary      Get recruitment by ID
// @Description  Retrieve a single recruitment record based on its ID.
// @Tags         Recruitments
// @Accept       json
// @Produce      json
// @Param        id   query     uint   true  "Recruitment ID"
// @Success      200  {object}  types.Response{data=handlers.RecruitmentResponse}  "Recruitment detail"
// @Failure      400  {object}  types.Response  "Missing or invalid ID parameter"
// @Failure      404  {object}  types.Response  "Recruitment not found"
// @Failure      500  {object}  types.Response  "Internal server error"
// @Router       /recruitments [get]
func GetRecruitmentById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		helper.JSONResponse(c, http.StatusBadRequest, "error", "id query param is required", nil)
		return
	}

	formatId, err := helper.StringToUint(id)
	if err != nil {
		helper.JSONResponse(c, http.StatusBadRequest, "error", "invalid id format", nil)
		return
	}

	recruitment, err := services.GetRecruitmentById(formatId)
	if err != nil {
		helper.JSONResponse(c, http.StatusNotFound, "error", err.Error(), nil)
		return
	}

	helper.JSONResponse(c, http.StatusOK, "success", "", recruitment)
}
