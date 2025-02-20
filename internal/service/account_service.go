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

type AccountService interface {
	SignUp(input *dto.RegisterBody) (*dto.RegisterResponse, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
	profileRepo repository.ProfileRepository
	hasher      hasher.Hasher
	jwtProvider tokenprovider.JWTTokenProvider
}

type AccountServiceConfig struct {
	AccountRepo repository.AccountRepository
	ProfileRepo repository.ProfileRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewAccountService(config AccountServiceConfig) AccountService {
	return &accountService{
		accountRepo: config.AccountRepo,
		profileRepo: config.ProfileRepo,
		hasher:      config.Hasher,
		jwtProvider: config.JwtProvider,
	}
}

func (s *accountService) SignUp(input *dto.RegisterBody) (*dto.RegisterResponse, error) {
	logger.Info("accountService", "Starting SignUp process", "username", input.UserName, "email", input.Email)

	// Hashing password
	hashed, err := s.hasher.Hash(input.Password)
	if err != nil {
		logger.Error("accountService", "Failed to hash password", "error", err)
		return nil, errs.ErrorGeneratingHashedPassword
	}
	logger.Info("accountService", "Password hashed successfully", "username", input.UserName)

	// Creating user account
	res, err := s.accountRepo.CreateUser(&dto.RegisterBody{
		UserName: input.UserName,
		Password: hashed,
		Email:    input.Email,
		Role:     input.Role,
	})
	if err != nil {
		logger.Error("accountService", "Failed to create user account", "username", input.UserName, "error", err)
		return nil, errs.ErrorCreatingAccount
	}
	logger.Info("accountService", "User account created", "user_id", res.ID, "username", res.Username)

	// Creating profile
	resProfile, err := s.profileRepo.CreateProfile(&model.Profile{AccountId: res.ID, Name: res.Username})
	if err != nil {
		logger.Error("accountService", "Failed to create user profile", "user_id", res.ID, "error", err)
		return nil, errs.ErrorOnCreatingNewProfile
	}
	logger.Info("accountService", "User profile created", "profile_id", resProfile.ID, "user_id", res.ID)

	// Preparing response
	respBody := &dto.RegisterResponse{
		UserID:   res.ID,
		Username: res.Username,
		Role:     res.Role,
		ProfileResponse: &dto.ProfileResponse{
			ID:      resProfile.ID,
			Name:    resProfile.Name,
			Address: resProfile.Address,
		},
	}

	logger.Info("accountService", "SignUp process completed successfully", "user_id", res.ID, "username", res.Username)
	return respBody, nil
}
