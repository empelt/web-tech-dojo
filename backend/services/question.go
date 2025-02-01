package services

import (
	"context"
	"strings"

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

func (s *QuestionService) SearchQuestions(ctx context.Context, keyword string, tags []string) ([]models.Question, error) {
	qs, err := s.questionRepository.FilterQuestionsByTags(ctx, tags)
	if err != nil {
		return nil, err
	}
	if keyword != "" {
		rtv := []models.Question{}
		for _, q := range qs {
			if strings.Contains(q.Title, keyword) {
				rtv = append(rtv, q)
			}
		}
		return rtv, nil
	} else {
		return qs, nil
	}
}
