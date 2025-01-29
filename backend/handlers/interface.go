package handlers

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
	"github.com/empelt/web-tech-dojo/services"
)

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type AuthService interface {
	AuthorizeAsUser(ctx context.Context, idToken string) (string, error)
}

type AnswerService interface {
	GetPreviousAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error)
	PostQuestionAnswer(ctx context.Context, uid string, qid int, message string) (*services.PostQuestionAnswerResponse, error)
}

type QuestionService interface {
	GetQuestion(ctx context.Context, id int) (*services.GetQuestionResponse, error)
}

type AnswerHandler struct {
	authService   AuthService
	answerService AnswerService
}

type QuestionHandler struct {
	questionService QuestionService
}
