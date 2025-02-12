package repository

import (
	"time"

	errs "github.com/Ayasibp/be-smart-farming-hydroponic/internal/errors"
	"github.com/Ayasibp/be-smart-farming-hydroponic/internal/model"
	"gorm.io/gorm"
)

type FarmRepository interface {
	CreateFarm(inputModel *model.Farm) (*model.Farm, error)
	GetFarms() ([]*model.Farm, error)
	GetFarmById(inputModel *model.Farm) (*model.Farm, error)
	UpdateFarm(inputModel *model.Farm) (*model.Farm, error)
	DeleteFarm(inputModel *model.Farm) (*model.Farm, error)
}

type farmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) FarmRepository {
	return &farmRepository{
		db: db,
	}
}

func (r *farmRepository) CreateFarm(inputModel *model.Farm) (*model.Farm, error) {

	sqlScript := `INSERT INTO hydroponic_system.farms (profile_id , name , address, created_at) 
				VALUES (?,?,?,?) 
				RETURNING profile_id, name, address;`

	res := r.db.Raw(sqlScript,
		inputModel.ProfileId,
		inputModel.Name,
		inputModel.Address,
		time.Now()).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	return inputModel, nil

}

func (r *farmRepository) GetFarms() ([]*model.Farm, error) {

	var farms []*model.Farm

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE deleted_at IS NULL`

	res := r.db.Raw(sqlScript).Scan(&farms)

	if res.Error != nil {
		return nil, res.Error
	}

	return farms, nil

}

func (r *farmRepository) GetFarmById(inputModel *model.Farm) (*model.Farm, error) {

	sqlScript := `SELECT * FROM hydroponic_system.farms 
				WHERE id = ?`

	res := r.db.Raw(sqlScript, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidFarmID
	}
	return inputModel, nil
}

func (r *farmRepository) UpdateFarm(inputModel *model.Farm) (*model.Farm, error) {

	sqlScript := `UPDATE hydroponic_system.farms 
				SET updated_at = ?, name = ?, address = ?  
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.Name, inputModel.Address, inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidFarmID
	}
	return inputModel, nil
}

func (r *farmRepository) DeleteFarm(inputModel *model.Farm) (*model.Farm, error) {

	sqlScript := `UPDATE hydroponic_system.farms 
				SET deleted_at = ? 
				WHERE id = ? 
				RETURNING *`

	res := r.db.Raw(sqlScript, time.Now(), inputModel.ID).Scan(&inputModel)

	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errs.InvalidFarmID
	}
	return inputModel, nil
}
