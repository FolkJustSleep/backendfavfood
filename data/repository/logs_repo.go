package repository

import (
	"fmt"

	"go-template/data/model"

	"gorm.io/gorm"
)

type LogsRepository struct {
	db *gorm.DB
}

type ILogsRepository interface {
	CreateLog(log model.Logs) (model.Logs, error)
	GetAllLogs() (*[]model.Logs, error)
	GetLogByID(id string) (*model.Logs, error)
	GetLogByUserID(id string) (*model.Logs, error)
	UpdateLog(log model.Logs) (*model.Logs, error)
	DeleteLog(id string) error
}

func NewLogsRepository(databaseConnection *gorm.DB) *LogsRepository {
	return &LogsRepository{
		db: databaseConnection,
	}
}


func (repo *LogsRepository) CreateLog(log model.Logs) (model.Logs, error) {
	result := repo.db.Create(&log)
	if result.Error != nil {
		return model.Logs{}, result.Error
	}
	fmt.Printf("Successfully created log: %v\n", log)
	return log, nil
}

func (repo *LogsRepository) GetAllLogs() (*[]model.Logs, error) {
	var logs []model.Logs
	result := repo.db.Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return &logs, nil
}

func (repo *LogsRepository) GetLogByID(id string) (*model.Logs, error) {
	var log model.Logs
	result := repo.db.Where("user_id = ?", id).First(&log)
	if result.Error != nil {
		return nil, result.Error
	}
	return &log, nil
}


func (repo *LogsRepository) GetLogByUserID(id string) (*model.Logs, error) {
	var log model.Logs
	result := repo.db.First(&log).Where("user_id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &log, nil
}



func (repo *LogsRepository) UpdateLog(log model.Logs) (*model.Logs, error) {
	result := repo.db.Save(&log)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Successfully updated log: %v\n", log)
	return &log, nil
}

func (repo *LogsRepository) DeleteLog(id string) error {
	result := repo.db.Delete(&model.Logs{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete log: %s", result.Error)
	}
	fmt.Printf("Successfully deleted log with ID: %s\n", id)
	return nil
}
