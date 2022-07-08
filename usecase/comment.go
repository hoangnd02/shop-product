package usecase

import (
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/repositories"
)

type CommentService interface {
	GetComment(int64) (*models.Comment, error)
	GetComments(int64) ([]*models.Comment, error)
	CreateComment(*models.Comment) (*models.Comment, error)
	UpdateComment(*models.Comment) (*models.Comment, error)
	DeleteComment(int64) error
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(r repositories.CommentRepository) CommentService {
	return commentService{
		commentRepository: r,
	}
}

func (s commentService) GetComment(id int64) (*models.Comment, error) {
	return s.commentRepository.GetComment(id)
}

func (s commentService) GetComments(idProduct int64) ([]*models.Comment, error) {
	return s.commentRepository.GetComments(idProduct)
}

func (s commentService) CreateComment(comment *models.Comment) (*models.Comment, error) {
	return s.commentRepository.CreateComment(comment)
}

func (s commentService) UpdateComment(comment *models.Comment) (*models.Comment, error) {
	return s.commentRepository.UpdateComment(comment)
}

func (s commentService) DeleteComment(id int64) error {
	return s.commentRepository.DeleteComment(id)
}
