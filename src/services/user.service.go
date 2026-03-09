package services

import (
	"sample-api/src/entities"
	"sample-api/src/repositories"
)

type UserServiceInterface interface {
	SaveUser(user entities.User) (entities.User, error)
	GetUserById(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(id int, user entities.User) error
	DeleteUser(id int) error
}

type UserService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) SaveUser(user entities.User) (entities.User, error) {
	return s.userRepo.SaveUser(user)
}

func (s *UserService) GetUserById(id int) (entities.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserService) UpdateUser(id int, user entities.User) error {
	return s.userRepo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepo.DeleteUser(id)
}
