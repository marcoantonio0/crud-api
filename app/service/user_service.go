package service

import (
	"github.com/marcoantonio63/crud-api/app/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(user repository.User) (any, error) {
	result, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) ListAllUsers() ([]repository.User, error) {
	result, err := s.userRepository.List()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserService) FindById(id string) (*repository.User, error) {
	result, err := s.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
