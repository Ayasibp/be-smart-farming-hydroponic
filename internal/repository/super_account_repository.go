package repository

import (
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
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

func (r *superAccountRepository) CreateSuperUser(input *model.SuperUser) (*model.SuperUser, error) {
	logger.Info("superAccountRepository", "Creating a new super user", "username", input.Username)

	sqlScript := `INSERT INTO super_admin.accounts (username, password, created_at) 
				  VALUES (?, ?, ?) 
				  RETURNING id, username, password;`

	res := r.db.Raw(sqlScript, input.Username, input.Password, time.Now()).Scan(input)

	if res.Error != nil {
		logger.Error("superAccountRepository", "Failed to create super user", "username", input.Username, "error", res.Error)
		return nil, res.Error
	}

	logger.Info("superAccountRepository", "Successfully created super user", "id", input.ID, "username", input.Username)
	return input, nil
}
