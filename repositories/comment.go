package repositories

import (
	"log"

	"github.com/hoanggggg5/shopproduct/models"
	"gorm.io/gorm"
)

type commentRepository struct {
	DB *gorm.DB
}

type CommentRepository interface {
	Migrate() error
	GetComment(int64) (*models.Comment, error)
	GetComments(int64) ([]*models.Comment, error)
	CreateComment(*models.Comment) (*models.Comment, error)
	UpdateComment(*models.Comment) (*models.Comment, error)
	DeleteComment(id int64) error
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return commentRepository{
		DB: db,
	}
}

func (p commentRepository) Migrate() error {
	log.Print("[CommentRepository]...Migrate")
	return p.DB.AutoMigrate(&models.Comment{})
}

func (p commentRepository) GetComment(id int64) (*models.Comment, error) {
	var comments *models.Comment
	if result := p.DB.Where("id = ?", id).Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}

func (p commentRepository) GetComments(idProduct int64) ([]*models.Comment, error) {
	var comments []*models.Comment
	if result := p.DB.Where("product_id = ?", idProduct).Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}

func (p commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if result := p.DB.Create(&comment); result.Error != nil {
		return comment, result.Error
	}
	return comment, nil
}

func (p commentRepository) UpdateComment(Comment *models.Comment) (*models.Comment, error) {
	if result := p.DB.Save(&Comment); result.Error != nil {
		return Comment, result.Error
	}
	return Comment, nil
}

func (p commentRepository) DeleteComment(id int64) error {
	if result := p.DB.Delete(&models.Comment{}, id); result != nil {
		return result.Error
	}
	return nil
}
