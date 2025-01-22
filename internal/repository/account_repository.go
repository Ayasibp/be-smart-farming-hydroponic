package repository

import (
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Begin() *gorm.DB
	CreateUser(input dto.RegisterBody) (*model.User, error)
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

func (r accountRepository) CreateUser(input dto.RegisterBody) (*model.User, error) {

	var inputModel *model.User

	res := r.db.Raw("INSERT INTO accounts (username , email , password, role, created_at) VALUES (?,?,?,?,?) RETURNING *;", input.UserName, input.Email, input.Password, input.Role).Scan(inputModel)
	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}
