package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
)

type AccountService interface {
	SignUp(input dto.RegisterBody) (*dto.RegisterResponse, error)
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

func (as accountService) SignUp(input dto.RegisterBody) (*dto.RegisterResponse, error) {

	hashed, err := as.hasher.Hash(input.Password)
	if err != nil {
		return nil, errs.ErrorGeneratingHashedPassword
	}

	res, err := as.accountRepo.CreateUser(&dto.RegisterBody{
		UserName: input.UserName,
		Password: hashed,
		Email:    input.Email,
		Role:     input.Role,
	})
	resCreatedProfile, err := as.profileRepo.CreateProfile(&dto.CreateProfile{
		AccountID: res.ID,
		Name:      res.Username,
		Address:   "",
	})
	respBody := &dto.RegisterResponse{
		UserID:   res.ID,
		Username: res.Username,
		Role:     res.Role,
		Profile: &dto.ProfileResponse{
			ID:      resCreatedProfile.ID,
			Name:    resCreatedProfile.Name,
			Address: resCreatedProfile.Address,
		},
	}

	return respBody, err
}
