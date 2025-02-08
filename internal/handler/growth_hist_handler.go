package handler

import (
	"encoding/hex"
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

func (h GrowthHistHandler) GetGrowthHistAggregationByFilter(c *gin.Context) {

	period := c.Query("period")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	farmId := c.Query("farm_id")
	systemId := c.Query("system_id")

	var startDateVal time.Time
	var endDateVal time.Time

	checkerFlag, err := getGrowthHistQueryParamsValidator(&period, &farmId, &systemId, &startDate, &endDate, &startDateVal, &endDateVal)

	if !checkerFlag {
		response.Error(c, 400, err.Error())
		return
	}

	resp, err := h.growthHistService.GetGrowthHistAggregationByFilter(&dto.GetGrowthFilter{
		FarmId:    farmId,
		SystemId:  systemId,
		StartDate: startDateVal,
		EndDate:   endDateVal,
		Period:    period,
	})
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Get "+resp.Period+" Aggregate Growth History Success", resp.AggregateData)
}
func (h GrowthHistHandler) GetGrowthHistByFilter(c *gin.Context) {

	period := "custom"
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	farmId := c.Query("farm_id")
	systemId := c.Query("system_id")

	var startDateVal time.Time
	var endDateVal time.Time

	checkerFlag, err := getGrowthHistQueryParamsValidator(&period, &farmId, &systemId, &startDate, &endDate, &startDateVal, &endDateVal)

	if !checkerFlag {
		response.Error(c, 400, err.Error())
		return
	}

	resp, err := h.growthHistService.GetGrowthHistByFilter(&dto.GetGrowthFilter{
		FarmId:    farmId,
		SystemId:  systemId,
		StartDate: startDateVal,
		EndDate:   endDateVal,
		Period:    period,
	})
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, "Get Growth History Success", resp)
}

func (h GrowthHistHandler) GenerateDummyData(c *gin.Context) {
	var createGrowthHistBody *dto.GrowthHistDummyDataBody

	if err := c.ShouldBindJSON(&createGrowthHistBody); err != nil {
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}
	resp, err := h.growthHistService.GenerateDummyData(createGrowthHistBody)
	if err != nil {
		response.Error(c, 400, err.Error())
		return
	}
	response.JSON(c, 200, " Success Generating random data", resp)
}

func getGrowthHistQueryParamsValidator(period *string, farmId *string, systemId *string, startDate *string, endDate *string, startDateVal *time.Time, endDateVal *time.Time) (bool, error) {

	if *farmId == "" {
		return false, errs.EmptyFarmIdParams
	}

	if *systemId == "" {
		return false, errs.EmptySystemIdParams
	}

	if *period == "" {
		return false, errs.EmptyPeriodQueryParams
	}
	if !(*period == "today" || *period == "last_3_days" || *period == "last_30_days" || *period == "custom") {
		return false, errs.InvalidValuePeriodQueryParams
	}

	if *period == "custom" {
		if *startDate == "" {
			return false, errs.EmptyStartDateQueryParams
		}
		if *endDate == "" {
			return false, errs.EmptyEndDateQueryParams
		}
		*startDateVal, _ = time.Parse("2006-01-02", *startDate)
		*endDateVal, _ = time.Parse("2006-01-02", *endDate)

		if startDateVal.Unix() >= endDateVal.Unix() {
			return false, errs.StartDateExceedEndDate
		}
	}

	return true, nil
}
