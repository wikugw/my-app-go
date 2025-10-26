package handlers

import (
	"fmt"
	"net/http"
	"time"

	"my-app/helper"
	"my-app/services"
	handlerTypes "my-app/types/handlers"

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
	deptID, err := helper.StringToUintPtr(req.Position)
	if err != nil {
		fmt.Println("❌ Invalid position ID:", err)
		helper.JSONResponse(c, http.StatusBadRequest, "error", "Invalid department ID", nil)
		return
	}

	employmentTypeID, err := helper.StringToUintPtr(req.EmploymentType)
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
		CreatedByID:          &req.CreatedByID,
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
