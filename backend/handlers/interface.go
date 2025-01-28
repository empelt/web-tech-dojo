package handlers

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/empelt/web-tech-dojo/services"
)

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type AnswerService interface {
	PostQuestionAnswer(ctx context.Context, uid string, qid int, message string) (*services.PostQuestionAnswerResponse, error)
}

type QuestionService interface {
	GetQuestion(ctx context.Context, id int) (*services.GetQuestionResponse, error)
}

type AnswerHandler struct {
	authClient    *auth.Client
	answerService AnswerService
}

type QuestionHandler struct {
	questionService QuestionService
}
