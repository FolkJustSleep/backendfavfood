package repository

import (
	"go-template/data/model"

	"gorm.io/gorm"
)

type TimestampRepository struct {
	db *gorm.DB
}

type ITimestampRepository interface {
	CreateTimestamp(timestamp model.TimestampModel) (model.TimestampModel, error)
	GetTimestampByUserID(id string) ([]model.TimestampModel, error)
	GetAllTimestamps() ([]model.TimestampModel, error)
	UpdateTimestamp(timestamp model.TimestampModel) (model.TimestampModel, error)
	DeleteTimestamp(id string) error
}

func NewTimestampRepository(databaseConnection *gorm.DB) *TimestampRepository {
	return &TimestampRepository{
		db: databaseConnection,
	}
}

func (repo *TimestampRepository) CreateTimestamp(timestamp model.TimestampModel) (model.TimestampModel, error) {
	if err := repo.db.Create(&timestamp).Error; err != nil {
		return model.TimestampModel{}, err
	}
	return timestamp, nil
}

func (repo *TimestampRepository) GetTimestampByUserID(UserID string) ([]model.TimestampModel, error) {
	var timestamp []model.TimestampModel
	if err := repo.db.Where("user_id = ?", UserID).Find(&timestamp).Error; err != nil {
		return nil, err
	}
	return timestamp, nil
}


func (repo *TimestampRepository) GetAllTimestamps() ([]model.TimestampModel, error) {
	var timestamps []model.TimestampModel
	if err := repo.db.Find(&timestamps).Error; err != nil {
		return nil, err
	}
	return timestamps, nil
}
func (repo *TimestampRepository) UpdateTimestamp(timestamp model.TimestampModel) (model.TimestampModel, error) {
	if err := repo.db.Save(&timestamp).Error; err != nil {
		return model.TimestampModel{}, err
	}
	return timestamp, nil
}

func (repo *TimestampRepository) DeleteTimestamp(id string) error {
	if err := repo.db.Delete(&model.TimestampModel{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}