package handler

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/service"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/response"
	"github.com/gin-gonic/gin"
)

type GrowthHistHandler struct {
	growthHistService service.GrowthHistService
	systemLogService  service.SystemLogService
}

type GrowthHistHandlerConfig struct {
	GrowthHistService service.GrowthHistService
	SystemLogService  service.SystemLogService
}

func NewGrowthHistHandler(config GrowthHistHandlerConfig) *GrowthHistHandler {
	return &GrowthHistHandler{
		growthHistService: config.GrowthHistService,
		systemLogService:  config.SystemLogService,
	}
}

func (h GrowthHistHandler) CreateGrowthHist(c *gin.Context) {
	var createGrowthHistBody *dto.GrowthHist

	if err := c.ShouldBindJSON(&createGrowthHistBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	resp, err := h.growthHistService.CreateGrowthHist(createGrowthHistBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	err = h.systemLogService.CreateSystemLog("Create Growth History: " + "{ID:" + hex.EncodeToString(resp.ID[:]) + "}")
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Create Growth History Success", resp)
}

func (h GrowthHistHandler) GetGrowthHistByFilter(c *gin.Context) {
	
	period:=c.Query("period")
    startDate:=c.Query("start_date")
	endDate:=c.Query("end_date")

	if period =="" {
		response.Error(c, 400, errs.EmptyPeriodQueryParams.Error())
		return
	}
	if !(period == "today" || period == "last_3_days" || period == "last_30_days"|| period == "custom"){
		response.Error(c, 400, errs.InvalidValuePeriodQueryParams.Error())
		return
	}

	if period=="custom"{
		if startDate =="" {
			response.Error(c, 400, errs.EmptyStartDateQueryParams.Error())
			return
		}
		if endDate =="" {
			response.Error(c, 400, errs.EmptyEndDateQueryParams.Error())
			return
		}
		startDateVal, _ := time.Parse("2006-01-02", startDate)
		endDateVal, _ := time.Parse("2006-01-02", endDate)
		
		if startDateVal.Unix() >= endDateVal.Unix(){
			response.Error(c, 400, errs.StartDateExceedEndDate.Error())
			return
		}
		fmt.Println(startDateVal)
		fmt.Println(endDateVal)
	}

	resp,err:=h.growthHistService.GetGrowthHistByFilter(&period)
	if err !=nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Get Growth History Success",resp)
}

func (h GrowthHistHandler) GenerateDummyData(c *gin.Context) {
	h.growthHistService.GenerateDummyData()
	response.JSON(c, 200, " Success Generating random data", "")
}
