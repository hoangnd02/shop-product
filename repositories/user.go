package repositories

import (
	"log"

	"github.com/hoanggggg5/shopproduct/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Migrate() error
	GetUser(string) (*models.User, error)
	CreateUser(*models.User) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (p userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return p.DB.AutoMigrate(&models.User{})
}

func (p userRepository) GetUser(uid string) (*models.User, error) {
	var Users *models.User
	if result := p.DB.Where("uid = ?", uid).Find(&Users); result.Error != nil {
		return Users, result.Error
	}
	return Users, nil
}

func (p userRepository) CreateUser(User *models.User) (*models.User, error) {
	if result := p.DB.Create(&User); result.Error != nil {
		return User, result.Error
	}
	return User, nil
}

func (p userRepository) UpdateUser(User *models.User) (*models.User, error) {
	if result := p.DB.Where("uid = ?", User.UID).Save(&User); result.Error != nil {
		return User, result.Error
	}
	return User, nil
}
