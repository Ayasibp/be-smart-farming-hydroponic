package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type SuperAccountRepository interface {
	CreateSuperUser(input *model.SuperUser) (*model.SuperUser, error)
}
type superAccountRepository struct {
	db *gorm.DB
}

func NewSuperAccountRepository(db *gorm.DB) SuperAccountRepository {
	return &superAccountRepository{
		db: db,
	}
}

func (r superAccountRepository) CreateSuperUser(input *model.SuperUser) (*model.SuperUser, error) {

	sqlScript:=`INSERT INTO super_admin.accounts (username , password, created_at) 
				VALUES (?,?,?) 
				RETURNING *;`

	res := r.db.Raw(sqlScript, input.Username, input.Password, time.Now()).Scan(&input)

	if res.Error != nil {
		return nil, res.Error
	}
	return input, nil
}
