package port

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
)

type AnswerRepository interface {
	FindAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error)
	UpsertAnswer(ctx context.Context, answer *models.Answer) (string, error)
}
