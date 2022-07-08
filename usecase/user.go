package usecase

import (
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/repositories"
)

type UserService interface {
	Get(string) (*models.User, error)
	Create(*models.User) (*models.User, error)
	Update(*models.User) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (s userService) Get(uid string) (*models.User, error) {
	return s.userRepository.GetUser(uid)
}

func (s userService) Create(User *models.User) (*models.User, error) {
	return s.userRepository.CreateUser(User)
}

func (s userService) Update(User *models.User) (*models.User, error) {
	return s.userRepository.UpdateUser(User)
}
