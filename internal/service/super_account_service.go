package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
)

type SuperAccountService interface {
	
}

type superAccountService struct {
	superAccountRepo repository.SuperAccountRepository
	hasher      hasher.Hasher
	jwtProvider tokenprovider.JWTTokenProvider
}

type AccountSuperServiceConfig struct {
	SuperAccountRepo repository.SuperAccountRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewAccountSuperService(config AccountSuperServiceConfig) SuperAccountService {
	return &superAccountService{
		superAccountRepo: config.SuperAccountRepo,
		hasher:      config.Hasher,
		jwtProvider: config.JwtProvider,
	}
}

func (s superAccountService) CreateUserSuper(input *dto.RegisterSuperUserBody) (*dto.RegisterSuperUserResponse, error) {

	hashed, err := s.hasher.Hash(input.Password)
	if err != nil {
		return nil, errs.ErrorGeneratingHashedPassword
	}

	res, err := s.superAccountRepo.CreateSuperUser(&model.SuperUser{
		Username: input.UserName,
		Password: hashed,
	})
	if err != nil {
		return nil, errs.ErrorCreatingAccount
	}

	return &dto.RegisterSuperUserResponse{
		UserID: res.ID,
		Username: res.Username,
	}, err
}
