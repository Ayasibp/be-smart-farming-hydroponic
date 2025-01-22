package service

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/repository"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/hasher"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/tokenprovider"
)

type AccountService interface {
	SignUp(input dto.RegisterBody) (*dto.RegisterBody, error)
}

type accountService struct {
	accountRepo repository.AccountRepository
	hasher      hasher.Hasher
	jwtProvider tokenprovider.JWTTokenProvider
}

type AccountServiceConfig struct {
	AccountRepo repository.AccountRepository
	Hasher      hasher.Hasher
	JwtProvider tokenprovider.JWTTokenProvider
}

func NewAccountService(config AccountServiceConfig) AccountService {
	return &accountService{
		accountRepo: config.AccountRepo,
		hasher:      config.Hasher,
		jwtProvider: config.JwtProvider,
	}
}

func (ts accountService) SignUp(input dto.RegisterBody) (*dto.RegisterBody, error) {
	// hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	// usernameLower := strings.ToLower(input.UserName)

	resp := &dto.RegisterBody{}

	return resp, nil
}
