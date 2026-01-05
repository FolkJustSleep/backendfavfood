package repository

import (
	"fmt"

	"go-template/data/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	GetAllUser() (*[]model.User, error)
	CreateUser(user model.User) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(user model.User) (*model.User, error)
	DeleteUser(id string) error
}

func NewUserRepository(databaseConnection *gorm.DB) IUserRepository {
	return &UserRepository{
		db: databaseConnection,
	}
}

func (repo *UserRepository) CreateUser(user model.User) (*model.User, error) {
	result := repo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Successfully created user: %v\n", user)
	return &user, nil
}


func (repo *UserRepository) GetAllUser() (*[]model.User, error) {
	var users []model.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		fmt.Printf("failed to get all user: %s\n", result.Error)
		return nil, result.Error
	}
	return &users, nil
}

func (repo *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	result := repo.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user by id: %s", result.Error)
	}
	return &user, nil
}

func (repo *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get user by email: %s", result.Error)
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user model.User) (*model.User, error) {
	result := repo.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Printf("Successfully updated user: %v\n", user)
	return &user, nil
}
//softdeleted
func (repo *UserRepository) DeleteUser(id string) error {
	result := repo.db.Where("id = ? ", id).Delete(&model.User{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete user: %s", result.Error)
	}
	fmt.Printf("Successfully deleted user with ID: %s\n", id)
	return nil
}
