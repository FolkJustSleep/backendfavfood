package service

import (
	"fmt"
	"time"
	"github.com/google/uuid"

	"go-template/data/model"
	"go-template/data/repository"

	fiberlog "github.com/gofiber/fiber/v2/log"
)

type TimestampService struct {
	TimestampRepository repository.ITimestampRepository
	UserRepository      repository.IUserRepository
}

type ITimestampService interface {
	CheckIn(timestamp model.TimestampModel) (model.TimestampModel, error)
	CheckOut(UserID string) (model.TimestampModel, error)
	GetCheckedInEmployee() (*[]*model.User, error)
	// GetTimestampByID(id string) (model.TimestampModel, error)
	// GetAllTimestamps() ([]model.TimestampModel, error)
	// UpdateTimestamp(timestamp model.TimestampModel) (model.TimestampModel, error)
	// DeleteTimestamp(id string) error
}

func NewTimestampService(timestampRepository repository.ITimestampRepository, userRepository repository.IUserRepository) ITimestampService {
	return &TimestampService{
		TimestampRepository: timestampRepository,
		UserRepository:      userRepository,
	}
}

func (sv *TimestampService) CheckIn(timestamp model.TimestampModel) (model.TimestampModel, error) {
	Allstamp , err := sv.TimestampRepository.GetTimestampByUserID(timestamp.UserID)
	if err != nil{
		fiberlog.Error("Error timestamp service (CheckIn - GetTimestampByUserID): ", err)
		return model.TimestampModel{}, err
	}
	Year , Month , Day := time.Now().Date()
	var TodayStamp model.TimestampModel
	for _, stamp := range Allstamp {
		if stamp.CheckIn.Day() == Day && stamp.CheckIn.Month() == Month && stamp.CheckIn.Year() == Year {
			fmt.Println("Stamp", stamp)
			TodayStamp = stamp
			break
		}
	}
	fmt.Println("TodayStamp:", TodayStamp)
	fmt.Printf("Day :%d, Month: %d, Year: %d\n", Day, Month, Year)
	fmt.Printf("Day :%d, Month: %d, Year: %d\n", TodayStamp.CheckIn.Day(), TodayStamp.CheckIn.Month(), TodayStamp.CheckIn.Year())
	if TodayStamp.ID != "" {
		fiberlog.Error("Error timestamp service (CheckIn): Already checked in for today")
		return model.TimestampModel{}, fmt.Errorf("already checked in for today")
	}
	timestamp.ID = uuid.New().String()
	timestamp.CheckIn = time.Now()
	timestamp.CheckOut = time.Time{}
	resp, err := sv.TimestampRepository.CreateTimestamp(timestamp)
	if err != nil {
		fiberlog.Error("Error timestamp service (CheckIn): ", err)
		return model.TimestampModel{}, err
	}
	return resp, nil
}

func (sv *TimestampService) CheckOut(UserID string) (model.TimestampModel, error) {
	Year , Month , Day := time.Now().Date()
	timestamp, err := sv.TimestampRepository.GetTimestampByUserID(UserID)
	if err != nil {
		fiberlog.Error("Error timestamp service (CheckOut - GetTimestampByUserID): ", err)
		return model.TimestampModel{}, err
	}
	var CheckOutStamp model.TimestampModel
	for _, stamp := range timestamp {
		if stamp.CheckIn.Day() != Day && stamp.CheckIn.Month() != Month && stamp.CheckIn.Year() != Year {
			continue
		}else {
			CheckOutStamp = stamp
			break
		}
	}
	if CheckOutStamp.ID == "" {
		fiberlog.Error("Error timestamp service (CheckOut): No check-in record found for today")
		return model.TimestampModel{},  fmt.Errorf("no check-in record found for today")
	}
	if !CheckOutStamp.CheckOut.IsZero() {
		fiberlog.Error("Error timestamp service (CheckOut): Already checked out for today")
		return model.TimestampModel{}, fmt.Errorf("already checked out for today")
	}
	CheckOutStamp.CheckOut = time.Now()
	resp, err := sv.TimestampRepository.UpdateTimestamp(CheckOutStamp)
	if err != nil {
		fiberlog.Error("Error timestamp service (CheckOut): ", err)
		return model.TimestampModel{}, err
	}
	return resp, nil
}

func (sv *TimestampService) GetCheckedInEmployee () (*[]*model.User, error) {
	Allstamp , err := sv.TimestampRepository.GetAllTimestamps()
	if err != nil{
		fiberlog.Error("Error timestamp service (GetCheckedInEmployee - GetAllTimestamps): ", err)
		return nil, err
	}
	Year , Month , Day := time.Now().Date()
	var CheckedInEmployees []model.TimestampModel
	for _, stamp := range Allstamp {
		if stamp.CheckIn.Day() == Day && stamp.CheckIn.Month() == Month && stamp.CheckIn.Year() == Year && stamp.CheckOut.IsZero() {
			CheckedInEmployees = append(CheckedInEmployees, stamp)
		}
	}
	var users []*model.User
	for _, emp := range CheckedInEmployees {
		fmt.Println("CheckedInEmployees:", emp.UserID)
		user, err := sv.UserRepository.GetUserByID(emp.UserID)
		if err != nil {
			fiberlog.Error("Error getting user by ID: ", err)
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}