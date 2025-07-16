package repository

import (
	"fmt"
	"time"

	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/dto"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/util/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Begin() *gorm.DB
	CreateUser(input *dto.RegisterBody) (*model.User, error)
	GetUserById(accountID uuid.UUID) (*model.User, error)
	GetUserByName(name *string) (*model.User, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{
		db: db,
	}
}

func (r *accountRepository) Begin() *gorm.DB {
	logger.Info("accountRepository", "Starting a new transaction", nil)
	return r.db.Begin()
}

func (r *accountRepository) CreateUser(input *dto.RegisterBody) (*model.User, error) {
	logger.Info("accountRepository", "Creating a new user", map[string]string{
		"username": input.UserName,
		"email":    input.Email,
	})

	var inputModel = &model.User{
		Username: input.UserName,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
	}

	sqlScript := `INSERT INTO hydroponic_system.accounts (username, email, password, role, created_at) 
				VALUES (?,?,?,?,?) 
				RETURNING id, username, email, password, role;`

	res := r.db.Raw(sqlScript, input.UserName, input.Email, input.Password, input.Role, time.Now()).Scan(inputModel)

	if res.Error != nil {
		logger.Error("accountRepository", "Failed to create user", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	logger.Info("accountRepository", "User created successfully", map[string]string{
		"username": inputModel.Username,
		"email":    inputModel.Email,
	})
	return inputModel, nil
}

func (r *accountRepository) GetUserById(accountID uuid.UUID) (*model.User, error) {
	logger.Info("accountRepository", "Fetching user by ID", map[string]string{
		"userID": accountID.String(),
	})

	var inputModel *model.User
	sqlScript := `SELECT id, username, email, role FROM hydroponic_system.accounts WHERE id = ?`

	res := r.db.Raw(sqlScript, accountID).Scan(&inputModel)

	if res.Error != nil {
		logger.Error("accountRepository", "Failed to fetch user by ID", map[string]string{
			"error":  res.Error.Error(),
			"userID": accountID.String(),
		})
		return nil, res.Error
	}

	logger.Info("accountRepository", "User fetched successfully", map[string]string{
		"userID":   inputModel.ID.String(),
		"username": inputModel.Username,
	})
	return inputModel, nil
}

func (r *accountRepository) GetUserByName(name *string) (*model.User, error) {
	logger.Info("accountRepository", "Fetching user by name", nil)

	var user *model.User

	sqlScript := `SELECT id, username, password 
				  FROM hydroponic_system.accounts 
				  WHERE 
				  	username = '%s' AND
				 	deleted_at IS NULL`

	sqlScript = fmt.Sprintf(sqlScript, *name)

	res := r.db.Raw(sqlScript).Scan(&user)

	if res.Error != nil {
		logger.Error("accountRepository", "Failed to fetch account", map[string]string{
			"error": res.Error.Error(),
		})
		return nil, res.Error
	}

	if user == nil {
		return nil, nil
	}

	logger.Info("accountRepository", "account fetched successfully", map[string]string{
		"username": user.Username,
	})
	return user, nil
}
