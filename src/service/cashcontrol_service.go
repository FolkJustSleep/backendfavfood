package service

import (
	"go-template/data/model"
	"go-template/data/repository"

	"github.com/google/uuid"
)

type CashControlService struct {
	CashControlRepo repository.ICashControlRepository
}

type ICashControlService interface {
	CreateCashControl(cashcontrol model.CashControl) (model.CashControl, error)
	GetCashControlByID(id string) (model.CashControl, error)
	GetCashControlByUserID(Userid string) (model.CashControl, error)
	GetAllCashControls() ([]model.CashControl, error)
	UpdateCashControl(cashcontrol model.CashControl) (model.CashControl, error)
	DeleteCashControl(id string) error
}

func NewCashControlService(CashControlRepo repository.ICashControlRepository) *CashControlService {
	return &CashControlService{
		CashControlRepo: CashControlRepo,
	}
}

func (sv *CashControlService) CreateCashControl(cashcontrol model.CashControl) (model.CashControl, error) {
	cashcontrol.ID = uuid.New().String()
	cashcontrol, err := sv.CashControlRepo.CreateCashControl(cashcontrol)
	if err != nil {
		return model.CashControl{}, err
	}
	return cashcontrol, nil
}

func (sv *CashControlService) GetCashControlByID(id string) (model.CashControl, error) {
	cashcontrol, err := sv.CashControlRepo.GetCashControlByID(id)
	if err != nil {
		return model.CashControl{}, err
	}
	return cashcontrol, nil
}

func (sv *CashControlService) GetCashControlByUserID(id string) (model.CashControl, error) {
	cashcontrol, err := sv.CashControlRepo.GetCashControlByID(id)
	if err != nil {
		return model.CashControl{}, err
	}
	return cashcontrol, nil
}

func (sv *CashControlService) GetAllCashControls() ([]model.CashControl, error) {
	cashcontrols, err := sv.CashControlRepo.GetAllCashControls()
	if err != nil {
		return nil, err
	}
	return cashcontrols, nil
}

func (sv *CashControlService) UpdateCashControl(cashcontrol model.CashControl) (model.CashControl, error) {
	cashcontrol, err := sv.CashControlRepo.UpdateCashControl(cashcontrol)
	if err != nil {
		return model.CashControl{}, err
	}
	return cashcontrol, nil
}

func (sv *CashControlService) DeleteCashControl(id string) error {
	if err := sv.CashControlRepo.DeleteCashControl(id); err != nil {
		return err
	}
	return nil
}
