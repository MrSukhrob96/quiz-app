package services

import "quiz-app/internal/repositories"

type UserService interface {
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	userRepository repositories.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) GetUsers() {

}

func (service *userService) GetUserById() {

}

func (service *userService) CreateUser() {

}

func (service *userService) UpdateUser() {

}

func (service *userService) DeleteUser() {

}

func (service *userService) GetUserByPhoneNumber() {

}
