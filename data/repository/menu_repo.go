package repository

import (
	"fmt"

	"go-template/data/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

type IMenuRepository interface {
	CreateMenu(menu model.Menu) (model.Menu, error)
	GetAllMenus() (*[]model.Menu, error)
	GetMenuByID(id string) (*model.Menu, error)
	UpdateMenu(menu model.Menu) (*model.Menu, error)
	DeleteMenu(id string) error
}

func NewMenuRepository(databaseConnection *gorm.DB) *MenuRepository {
	return &MenuRepository{
		db: databaseConnection,
	}
}
func (repo *MenuRepository) CreateMenu(menu model.Menu) (model.Menu, error) {
	if menu.Items != nil {
		for i := range menu.Items {
			if menu.Items[i].ID == "" {
				menu.Items[i].ID = uuid.New().String()
			}
		}
	}
	if err := repo.db.Create(&menu).Error; err != nil {
		return model.Menu{}, fmt.Errorf("failed to create menu: %w", err)
	}
	return menu, nil
}

func (repo *MenuRepository) GetAllMenus() (*[]model.Menu, error) {
	var menus []model.Menu
	if err := repo.db.Find(&menus).Error; err != nil {
		return nil, fmt.Errorf("failed to get all menus: %w", err)
	}
	return &menus, nil
}
func (repo *MenuRepository) GetMenuByID(id string) (*model.Menu, error) {
	var menu model.Menu
	if err := repo.db.Where("id = ?", id).First(&menu).Error; err != nil {
		return nil, fmt.Errorf("failed to get menu by ID: %w", err)
	}
	return &menu, nil
}
func (repo *MenuRepository) UpdateMenu(menu model.Menu) (*model.Menu, error) {
	if err := repo.db.Save(&menu).Error; err != nil {
		return nil, fmt.Errorf("failed to update menu: %w", err)
	}
	return &menu, nil
}
func (repo *MenuRepository) DeleteMenu(id string) error {
	if err := repo.db.Delete(&model.Menu{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete menu: %w", err)
	}
	return nil
}
