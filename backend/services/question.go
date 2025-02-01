package services

import (
	"context"
	"slices"
	"time"
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

func NewQuestionService(questionRepository QuestionRepository, userService *UserService) (*QuestionService, error) {
	return &QuestionService{
		questionRepository: questionRepository,
		userService:        userService,
	}, nil
}

func (s *QuestionService) GetQuestion(ctx context.Context, uid string, qid int) (*GetQuestionResponse, error) {
	q, err := s.questionRepository.FindQuestion(ctx, qid)
	if err != nil {
		return nil, err
	}
	u, err := s.userService.userRepository.GetUser(ctx, uid)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	qss := []QuestionSummary{}
	for _, q := range qs {
		progress := 0
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
