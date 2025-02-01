package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
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

func (s accountService) SignUp(input *dto.RegisterBody) (*dto.RegisterResponse, error) {

	hashed, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, errs.ErrorGeneratingHashedPassword
	}

	res, err := s.accountRepo.CreateUser(&dto.RegisterBody{
		UserName: input.UserName,
		Password: hashed,
		Email:    input.Email,
		Role:     input.Role,
	})
	if err != nil {
		return nil, errs.ErrorCreatingAccount
	}

	resProfile, err := s.profileRepo.CreateProfile(&model.Profile{AccountId: res.ID, Name: res.Username})
	if err != nil {
		return nil, errs.ErrorOnCreatingNewProfile
	}

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

	return respBody, err
}
