package services

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
)

type GetQuestionResponse struct {
	Question *models.Question
}

func NewQuestionService(questionRepository QuestionRepository) (*QuestionService, error) {
	return &QuestionService{
		questionRepository: questionRepository,
	}, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, id int) (*GetQuestionResponse, error) {
	q, err := s.questionRepository.FindQuestion(ctx, id)
	if err != nil {
		return nil, err
	}
	return &GetQuestionResponse{
		Question: q,
	}, nil
}

func (s *QuestionService) GetAllQuestions(ctx context.Context) ([]models.Question, error) {
	qs, err := s.questionRepository.GetAllQuestions(ctx)
	if err != nil {
		return nil, err
	}
	return qs, nil
}
