package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"go-template/data/model"
	"go-template/data/repository"
	"go-template/src/middleware"

	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository repository.IUserRepository
	LogsRepository repository.ILogsRepository
}

type IAuthService interface {
	Login(email string, password string) (string, error)
	Register(user model.User) (*model.User, error)
	Logout(ctx *fiber.Ctx) error
}

func NewAuthService(userRepository repository.IUserRepository, logsRepository repository.ILogsRepository) IAuthService {
	return &AuthService{
		UserRepository: userRepository,
		LogsRepository: logsRepository,
	}
}

func (sv *AuthService) Login(email string, password string) (string, error) {
	var log model.Logs
	log.Action = "User Login"
	log.Status = "complete"
	log.ID = uuid.New().String()
	user, err := sv.UserRepository.GetUserByEmail(email)
	if err != nil {
		fiberlog.Error("Error getting user by email: ", err)
		log.Status = "failed" + err.Error()
		log.UserID = ""
		_, err = sv.LogsRepository.CreateLog(log)
		if err != nil {
			fiberlog.Error("Error creating log: ", err)
		}
		return "", err
	}
	log.UserID = user.ID
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fiberlog.Error("Invalid password: ", err)
		log.Status = "failed" + err.Error()
		_, err = sv.LogsRepository.CreateLog(log)
		if err != nil {
			fiberlog.Error("Error creating log: ", err)
		}
		return "", err
	}
	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		fiberlog.Error("Error generating token: ", err)
		log.Status = "failed" + err.Error()
		_, err = sv.LogsRepository.CreateLog(log)
		if err != nil {
			fiberlog.Error("Error creating log: ", err)
		}
		return "", err
	}
	_ , err = sv.LogsRepository.CreateLog(log)
	if err != nil {
		fiberlog.Error("Error creating log: ", err)
	}
	return *token.Token, nil
}

func (sv *AuthService) Register(user model.User) (*model.User, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
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

func (sv *AuthService) Logout(ctx *fiber.Ctx) error {
	Token , err := middleware.DecodeCookie(ctx)
	if err != nil {
		return err
	}

	if Token == nil {
		return fiber.ErrUnauthorized
	}
	ctx.ClearCookie("token")
	cookies := ctx.Cookies("token")
	fmt.Println("Cookies after clearing:", cookies)
	// var log model.Logs
	// log.UserID = Token.UserID
	// log.Action = "User Logout"
	// log.Status = "complete"
	// log.ID = uuid.New().String()
	// _ , err = sv.LogsRepository.CreateLog(log)
	// if err != nil {
	// 	fiberlog.Error("Error creating log: ", err)
	// 	return err
	// }
	return nil
}
