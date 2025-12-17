package service

import (
	"go-template/data/model"
	"go-template/data/repository"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type MenuService struct {
	MenuRepo repository.IMenuRepository
}

type IMenuService interface {
	CreateMenu(menu model.Menu) (model.Menu, error)
	GetAllMenus() (*[]model.Menu, error)
	GetMenuByID(id string) (*model.Menu, error)
	UpdateMenu(menu model.Menu) (*model.Menu, error)
	DeleteMenu(id string) error
}

func NewMenuService(menuRepo repository.IMenuRepository) IMenuService {
	return &MenuService{
		MenuRepo: menuRepo,
	}
}

func (sv *MenuService) CreateMenu(menu model.Menu) (model.Menu, error) {
	menu.ID = uuid.New().String()
	resp, err := sv.MenuRepo.CreateMenu(menu)
	if err != nil {
		fiberlog.Error("Error menu service (CreateMenu): ", err)
		return model.Menu{}, err
	}
	return resp, nil
}

func (sv *MenuService) GetAllMenus() (*[]model.Menu, error) {
	resp, err := sv.MenuRepo.GetAllMenus()
	if err != nil {
		fiberlog.Error("Error menu service (GetAllMenus): ", err)
		return nil, err
	}
	return resp, nil
}

func (sv *MenuService) GetMenuByID(id string) (*model.Menu, error) {
	resp, err := sv.MenuRepo.GetMenuByID(id)
	if err != nil {
		fiberlog.Error("Error menu service (GetMenuByID): ", err)
		return nil, err
	}
	return resp, nil
}

func (sv *MenuService) UpdateMenu(menu model.Menu) (*model.Menu, error) {
	resp, err := sv.MenuRepo.UpdateMenu(menu)
	if err != nil {
		fiberlog.Error("Error menu service (UpdateMenu): ", err)
		return nil, err
	}
	return resp, nil
}

func (sv *MenuService) DeleteMenu(id string) error {
	err := sv.MenuRepo.DeleteMenu(id)
	if err != nil {
		fiberlog.Error("Error menu service (DeleteMenu): ", err)
		return err
	}
	return nil
}	