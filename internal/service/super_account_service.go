package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
)

type SuperAccountService interface {
	CreateSuperUser(input *dto.RegisterSuperUserBody) (*dto.RegisterSuperUserResponse, error)
}

type superAccountService struct {
	superAccountRepo repository.SuperAccountRepository
	hasher           hasher.Hasher
	jwtProvider      tokenprovider.JWTTokenProvider
}

type SuperAccountServiceConfig struct {
	SuperAccountRepo repository.SuperAccountRepository
	Hasher           hasher.Hasher
	JwtProvider      tokenprovider.JWTTokenProvider
}

func NewSuperAccountService(config SuperAccountServiceConfig) SuperAccountService {
	return &superAccountService{
		superAccountRepo: config.SuperAccountRepo,
		hasher:           config.Hasher,
		jwtProvider:      config.JwtProvider,
	}
}

func (s *superAccountService) CreateSuperUser(input *dto.RegisterSuperUserBody) (*dto.RegisterSuperUserResponse, error) {
	logger.Info("superAccountService", "Creating Super User", map[string]string{
		"username": input.UserName,
	})

	hashed, err := s.hasher.Hash(input.Password)
	if err != nil {
		logger.Error("superAccountService", "Error hashing password", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorGeneratingHashedPassword
	}

	res, err := s.superAccountRepo.CreateSuperUser(&model.SuperUser{
		Username: input.UserName,
		Password: hashed,
	})
	if err != nil {
		logger.Error("superAccountService", "Error creating Super User", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorCreatingAccount
	}

	logger.Info("superAccountService", "Super User created successfully", map[string]string{
		"userId": res.ID.String(),
	})

	return &dto.RegisterSuperUserResponse{
		UserID:   res.ID,
		Username: res.Username,
	}, nil
}
