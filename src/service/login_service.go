package service

import (
	// "go-template/data/model"
	"go-template/data/repository"
	"go-template/src/middleware"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepository repository.IUserRepository
	LogsRepository  repository.ILogsRepository
}

type ILoginService interface {
	Login(email string, password string) (string, error)
}

func NewLoginService(userRepository repository.IUserRepository, logsRepository repository.ILogsRepository) ILoginService {
	return &LoginService{
		UserRepository: userRepository,
		LogsRepository:  logsRepository,
	}
}

func (sv *LoginService) Login(email string, password string) (string, error) {
	user, err := sv.UserRepository.GetUserByEmail(email)
	if err != nil {
		fiberlog.Error("Error getting user by email: ", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fiberlog.Error("Invalid password: ", err)
		return "", err
	}
	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		fiberlog.Error("Error generating token: ", err)
		return "", err
	}
	return *token.Token, nil
}
