package repositories

import (
	"quiz-app/internal/dto"
	"quiz-app/internal/entities"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() (entities.Users, error)
	GetById(id *uuid.UUID) (entities.User, error)
	Create(userDTO dto.CreateUserDTO) (entities.User, error)
	Update(id *uuid.UUID, userDTO dto.UpdateUserDTO) (entities.User, error)
	Delete(id *uuid.UUID) error
	GetByPhoneNumber(phone string) (entities.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (repository *userRepository) GetAll() (entities.Users, error) {
	return nil, nil
}

func (repository *userRepository) GetById(id *uuid.UUID) (entities.User, error) {
	return entities.User{}, nil
}

func (repository *userRepository) Create(userDTO dto.CreateUserDTO) (entities.User, error) {
	return entities.User{}, nil
}

func (repository *userRepository) Update(id *uuid.UUID, userDTO dto.UpdateUserDTO) (entities.User, error) {
	return entities.User{}, nil
}

func (repository *userRepository) Delete(id *uuid.UUID) error {
	return nil
}

func (repository *userRepository) GetByPhoneNumber(phone string) (entities.User, error) {
	return entities.User{}, nil
}
