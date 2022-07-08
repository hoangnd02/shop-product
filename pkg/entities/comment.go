package entities

import "github.com/hoanggggg5/shopproduct/models"

type CommentEntity struct {
	Content string `json:"content"`
}

func CommentToEntity(comment models.Comment) CommentEntity {
	return CommentEntity{
		Content: comment.Content,
	}
}
