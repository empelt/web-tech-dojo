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

type AnswerRepository interface {
	FindAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error)
	BulkUpsertAnswer(ctx context.Context, answer *models.Answer, newMessages []models.Message) error
}

type AnswerService struct {
	genaiClient        GenaiClient
	questionRepository QuestionRepository
	answerRepository   AnswerRepository
}

type QuestionService struct {
	questionRepository QuestionRepository
}
