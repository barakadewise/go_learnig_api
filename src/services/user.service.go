package services

import (
	"fmt"
	"sample-api/src/entities"
)

//user service interface and implementation
type UserServiceInterface interface {
	SaveUser(user entities.User) (entities.User, error)
	GetUserById(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	UpdateUser(id int, user entities.User) error
	DeleteUser(id int) error
}


type UserService struct {
	users []entities.User
	nextID int

}


func NewUserService() UserServiceInterface {
	return &UserService{
		users:  []entities.User{},
		nextID: 1,
	}
}

func (s *UserService) SaveUser(user entities.User) (entities.User, error) {
	user.ID = s.nextID
	s.nextID++
	s.users = append(s.users, user)
	return user, nil
}

func (s *UserService) GetUserById(id int) (entities.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return entities.User{}, fmt.Errorf("user not found")
}

func (s *UserService) GetAllUsers() ([]entities.User, error) {
	return s.users, nil
}

func (s *UserService) UpdateUser(id int, updated entities.User) error {
	for i, u := range s.users {
		if u.ID == id {
			updated.ID = id
			s.users[i] = updated
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (s *UserService) DeleteUser(id int) error {
	for i, u := range s.users {
		if u.ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
