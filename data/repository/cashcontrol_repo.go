package repository

import (
	"fmt"

	"go-template/data/model"

	"gorm.io/gorm"
)

type CashControlRepository struct {
	db *gorm.DB
}

type ICashControlRepository interface {
	CreateCashControl(cashcontrol model.CashControl) (model.CashControl, error)
	GetCashControlByID(id string) (model.CashControl, error)
	GetCashControlByUserID(Userid string) (model.CashControl, error)
	GetAllCashControls() ([]model.CashControl, error)
	UpdateCashControl(cashcontrol model.CashControl) (model.CashControl, error)
	DeleteCashControl(id string) error
}
func NewCashControlRepository(databaseConnection *gorm.DB) *CashControlRepository {
	return &CashControlRepository{
		db: databaseConnection,
	}
}

func (repo *CashControlRepository) CreateCashControl(cashcontrol model.CashControl) (model.CashControl, error) {
	if err := repo.db.Create(&cashcontrol).Error; err != nil {
		return model.CashControl{}, fmt.Errorf("failed to create cash control: %w", err)
	}
	return cashcontrol, nil
}

func (repo *CashControlRepository) GetCashControlByID(id string) (model.CashControl, error) {
	var cashcontrol model.CashControl
	if err := repo.db.Where("id = ?", id).First(&cashcontrol).Error; err != nil {
		return model.CashControl{}, fmt.Errorf("failed to get cash control by ID: %w", err)
	}
	return cashcontrol, nil
}

func (repo *CashControlRepository) GetCashControlByUserID(Userid string) (model.CashControl, error) {
	var cashcontrol model.CashControl
	if err := repo.db.Where("user_id = ?", Userid).First(&cashcontrol).Error; err != nil {
		return model.CashControl{}, fmt.Errorf("failed to get cash control by user ID: %w", err)
	}
	return cashcontrol, nil
}

func (repo *CashControlRepository) GetAllCashControls() ([]model.CashControl, error) {
	var cashcontrols []model.CashControl
	if err := repo.db.Find(&cashcontrols).Error; err != nil {
		return nil, fmt.Errorf("failed to get all cash controls: %w", err)
	}
	return cashcontrols, nil
}
func (repo *CashControlRepository) UpdateCashControl(cashcontrol model.CashControl) (model.CashControl, error) {
	if err := repo.db.Save(&cashcontrol).Error; err != nil {
		return model.CashControl{}, fmt.Errorf("failed to update cash control: %w", err)
	}
	return cashcontrol, nil
}
func (repo *CashControlRepository) DeleteCashControl(id string) error {
	if err := repo.db.Where("id = ?", id).Delete(&model.CashControl{}).Error; err != nil {
		return fmt.Errorf("failed to delete cash control: %w", err)
	}
	return nil
}