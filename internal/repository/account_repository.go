package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Begin() *gorm.DB
	CreateUser(input *dto.RegisterBody) (*model.User, error)
	GetUserById(accountID uuid.UUID)(*model.User, error)
}
type accountRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r accountRepository) Begin() *gorm.DB {
	return r.db.Begin()
}

func (r accountRepository) CreateUser(input *dto.RegisterBody) (*model.User, error) {
	var inputModel = &model.User{
		Username: input.UserName,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	res := r.db.Raw("INSERT INTO hydroponic_system.accounts (username , email , password, role, created_at) VALUES (?,?,?,?,?) RETURNING *;", input.UserName, input.Email, input.Password, input.Role, time.Now()).Scan(inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil
}

func (r accountRepository) GetUserById(accountID uuid.UUID)(*model.User, error) {
	var inputModel  *model.User

	res := r.db.Raw("SELECT id FROM hydroponic_system.accounts where id = ?", accountID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}

	return inputModel ,nil
}
