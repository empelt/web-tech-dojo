package services

import (
	"context"
	"fmt"
	"slices"
	"time"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
	"github.com/empelt/web-tech-dojo/services/port"
)

type GetQuestionResponse struct {
	Id           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Tags         []string  `json:"tags"`
	IsBookmarked bool      `json:"isBookmarked"`
	CreatedAt    time.Time `json:"createdAt"`
}

type QuestionSummary struct {
	Id           int      `json:"id"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	IsBookmarked bool     `json:"isBookmarked"`
	Progress     int      `json:"progress"`
}

type QuestionService struct {
	questionRepository port.QuestionRepository
	userService        *UserService
	tx                 port.TxExecutor
}

func NewQuestionService(questionRepository port.QuestionRepository, userService *UserService, tx port.TxExecutor) (*QuestionService, error) {
	return &QuestionService{
		questionRepository: questionRepository,
		userService:        userService,
		tx:                 tx,
	}, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, uid string, qid int) (*GetQuestionResponse, error) {
	q, err := s.questionRepository.FindQuestion(ctx, qid)
	if err != nil {
		return nil, err
	}
	u, err := s.userService.userRepository.GetUser(ctx, uid)
	if err != nil {
		if err == models.EntityNotFoundError {
			u = &models.User{
				UserId:      uid,
				QuestionIds: []int{},
				Progresses:  []models.Progress{},
			}
		} else {
			return nil, err
		}
	}

	return &GetQuestionResponse{
		Id:           q.Id,
		Title:        q.Title,
		Content:      q.Content,
		Tags:         q.Tags,
		IsBookmarked: slices.Contains(u.QuestionIds, qid),
		CreatedAt:    q.CreatedAt,
	}, nil
}

func (s *QuestionService) GetAllQuestions(ctx context.Context, uid string) ([]QuestionSummary, error) {
	qs, err := s.questionRepository.GetAllQuestions(ctx)
	if err != nil {
		return nil, err
	}
	u, err := s.userService.GetUser(ctx, uid)
	if err != nil {
		if err == models.EntityNotFoundError {
			u = &models.User{
				UserId:      uid,
				QuestionIds: []int{},
				Progresses:  []models.Progress{},
			}
		} else {
			return nil, err
		}
	}

	qss := []QuestionSummary{}
	for _, q := range qs {
		progress := -1
		for _, p := range u.Progresses {
			if p.QuestionId == q.Id {
				progress = p.Progress
			}
		}
		qss = append(qss, QuestionSummary{
			Id:           q.Id,
			Title:        q.Title,
			Tags:         q.Tags,
			IsBookmarked: slices.Contains(u.QuestionIds, q.Id),
			Progress:     progress,
		})
	}
	return qss, nil
}

func (s *QuestionService) UpsertQuestions(ctx context.Context, questions []models.Question) error {
	totalQuestions := len(questions)
	if totalQuestions == 0 {
		return nil
	}

	// トランザクションの上限を超えないようにバッチに分けて挿入
	for i := 0; i < totalQuestions; i += infrastructures.MaxOperationsPerTransaction {
		end := i + infrastructures.MaxOperationsPerTransaction
		if end > totalQuestions {
			end = totalQuestions
		}
		batchQuestions := questions[i:end]

		err := s.tx.ExecQuestionTx(ctx, func(ctx context.Context, repo port.QuestionRepository) error {
			for _, question := range batchQuestions {
				err := repo.UpsertQuestionWithTx(ctx, question)
				if err != nil {
					return fmt.Errorf("failed to upsert question: %v", err)
				}
			}
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to execute transaction: %v", err)
		}
	}
	return nil
}
