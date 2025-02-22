package service

import (
	"strconv"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/google/uuid"
)

type UnitIdService interface {
	CreateUnitId() (*dto.UnitIdResponse, error)
	GetUnitIds() ([]*dto.UnitIdResponse, error)
	DeleteUnitIdbyId(unitId *uuid.UUID) (*dto.UnitIdResponse, error)
}

type unitIdService struct {
	unitIdRepo repository.UnitIdRepository
}

type UnitIdServiceConfig struct {
	UnitIdRepo repository.UnitIdRepository
}

func NewUnitIdService(config UnitIdServiceConfig) UnitIdService {
	return &unitIdService{
		unitIdRepo: config.UnitIdRepo,
	}
}

func (s *unitIdService) CreateUnitId() (*dto.UnitIdResponse, error) {
	logger.Info("unitIdService", "Creating new unit ID", nil)

	res, err := s.unitIdRepo.CreateUnitId()
	if err != nil {
		logger.Error("unitIdService", "Error creating unit ID", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorCreatingAccount
	}

	logger.Info("unitIdService", "Successfully created unit ID", map[string]string{
		"unit_id": res.ID.String(),
	})
	return &dto.UnitIdResponse{
		ID: res.ID,
	}, err
}

func (s *unitIdService) GetUnitIds() ([]*dto.UnitIdResponse, error) {
	logger.Info("unitIdService", "Fetching all unit IDs", nil)

	var unitIdRes []*dto.UnitIdResponse
	res, err := s.unitIdRepo.GetUnitIds()
	if err != nil {
		logger.Error("unitIdService", "Error fetching unit IDs", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorCreatingAccount
	}

	for i := 0; i < len(res); i++ {
		unitIdRes = append(unitIdRes, &dto.UnitIdResponse{
			ID: res[i].ID,
		})
	}

	logger.Info("unitIdService", "Successfully fetched unit IDs", map[string]string{
		"count": strconv.Itoa(len(unitIdRes)),
	})
	return unitIdRes, err
}

func (s *unitIdService) DeleteUnitIdbyId(unitId *uuid.UUID) (*dto.UnitIdResponse, error) {
	logger.Info("unitIdService", "Deleting unit ID", map[string]string{
		"unit_id": unitId.String(),
	})

	res, err := s.unitIdRepo.DeleteUnitIdById(&model.UnitId{ID: *unitId})
	if err != nil {
		logger.Error("unitIdService", "Error deleting unit ID", map[string]string{
			"unit_id": unitId.String(),
			"error":   err.Error(),
		})
		return nil, err
	}

	logger.Info("unitIdService", "Successfully deleted unit ID", map[string]string{
		"unit_id": unitId.String(),
	})
	return &dto.UnitIdResponse{
		ID: res.ID,
	}, err
}
