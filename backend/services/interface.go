package services

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
)

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type GenaiClient interface {
	GenerateContentFromText(ctx context.Context, message string) (string, error)
}

type QuestionRepository interface {
	FindQuestion(ctx context.Context, id int) (*models.Question, error)
}

type ChatService struct {
	genaiClient GenaiClient
}

type QuestionService struct {
	questionRepository QuestionRepository
}
