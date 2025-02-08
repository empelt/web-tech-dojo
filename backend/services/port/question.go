package port

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
)

type QuestionRepository interface {
	FindQuestion(ctx context.Context, id int) (*models.Question, error)
	GetAllQuestions(ctx context.Context) ([]models.Question, error)
	UpsertQuestion(ctx context.Context, question models.Question) (string, error)
	UpsertQuestionWithTx(ctx context.Context, question models.Question) error
}

type TxExecutor interface {
	ExecQuestionTx(ctx context.Context, f func(ctx context.Context, repo QuestionRepository) error) error
}
