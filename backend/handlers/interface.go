package handlers

import (
	"context"

	"github.com/empelt/web-tech-dojo/services"
)

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type ChatService interface {
	PostQuestionAnswer(ctx context.Context, message string) (string, error)
}

type QuestionService interface {
	GetQuestion(ctx context.Context, id int) (*services.GetQuestionResponse, error)
}

type ChatHandler struct {
	chatService ChatService
}

type QuestionHandler struct {
	questionService QuestionService
}
