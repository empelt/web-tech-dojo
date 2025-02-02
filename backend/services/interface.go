package services

import (
	"context"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type GenaiClient interface {
	GenerateContentFromText(ctx context.Context, message string) (string, error)
}

type QuestionRepository interface {
	FindQuestion(ctx context.Context, id int) (*models.Question, error)
	GetAllQuestions(ctx context.Context) ([]models.Question, error)
}

type AnswerRepository interface {
	FindAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error)
	UpsertAnswer(ctx context.Context, answer *models.Answer, newMessages []models.Message) (string, error)
}

type UserRepository interface {
	GetUser(ctx context.Context, uid string) (*models.User, error)
	UpsertUser(ctx context.Context, uid string, u *models.User) (string, error)
}

type AuthService struct {
	firebaseAuth *infrastructures.FirebaseAuth
}

type UserService struct {
	userRepository UserRepository
}

type AnswerService struct {
	genaiClient        GenaiClient
	userRepository     UserRepository
	questionRepository QuestionRepository
	answerRepository   AnswerRepository
}

type QuestionService struct {
	questionRepository QuestionRepository
	userService        *UserService
}
