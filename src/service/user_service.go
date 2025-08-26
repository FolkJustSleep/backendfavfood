package service

import (
	"time"
	"fmt"
	"github.com/google/uuid"

	"go-template/data/model"
	"go-template/data/repository"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository repository.IUserRepository
	LogsRepository repository.ILogsRepository
}

type IUserService interface {
	CreateUser(user model.User) (*model.User, error)
	GetAllUser() (*[]model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUser(user model.User) (*model.User, error)
	DeleteUser(id string) error
}

func NewUserService(userRepository repository.IUserRepository, logsRepository repository.ILogsRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
		LogsRepository: logsRepository,
	}
}

func (sv *UserService) CreateUser(user model.User) (*model.User, error) {
	hashedpassword , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fiberlog.Error("Error hashing password: ", err)
		return nil, err
	}
	user.Password = string(hashedpassword)
	time := time.Now()
	user.CreatedAt = time
	user.ID = uuid.New().String()
	fiberlog.Info("Creating user with ID: ", user.ID)
	resp, err := sv.UserRepository.CreateUser(user)
	if err != nil {
		fiberlog.Error(err)
		return nil, err
	}
	return resp, nil
}

func (sv *UserService) GetAllUser() (*[]model.User, error) {
	data, err := sv.UserRepository.GetAllUser()
	if err != nil {
		fiberlog.Error("Error getting all users: ", err)
		return nil, err
	}
	return data, nil
}

func (sv *UserService) GetUserByID(id string) (*model.User, error) {
	fiberlog.Info("[DEBUG] service.GetUserByID called with id:", id)
	data, err := sv.UserRepository.GetUserByID(id)
	if err != nil {
		fiberlog.Error("Error getting user by ID: ", err)
		return nil, err
	}
	return data, nil
}

func (sv *UserService) UpdateUser(user model.User) (*model.User, error) {
	if user.ID == "" {
		fiberlog.Error("Error updating user: ID is empty")
		return nil, fmt.Errorf("ID is empty")
	}
	basedata , err := sv.UserRepository.GetUserByID(user.ID)
	if err != nil {
		fiberlog.Error("Error updating user: failed to get user :", err)
		return nil, err
	}
	if user.Name == "" {
		user.Name = basedata.Name
	}
	if user.Email == "" {
		user.Email = basedata.Email
	}
	data, err := sv.UserRepository.UpdateUser(user)
	if err != nil {
		fiberlog.Error("Error updating user: ", err)
		return nil, err
	}
	return data, nil
}

func (sv *UserService) DeleteUser(id string) error {
	err := sv.UserRepository.DeleteUser(id)
	if err != nil {
		fiberlog.Error("Error deleting user: ", err)
		return err
	}
	return nil
}
