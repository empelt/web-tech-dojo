package services

import (
	"context"
	"slices"

	"github.com/empelt/web-tech-dojo/models"
)

func NewUserService(bookmarkRepository UserRepository) (*UserService, error) {
	return &UserService{
		userRepository: bookmarkRepository,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, uid string) (*models.User, error) {
	// 1.既存のお気に入りデータを取得
	b, err := s.userRepository.GetUser(ctx, uid)
	if err == models.EntityNotFoundError {
		// 既存のお気に入りデータがない場合は空データを作成
		return &models.User{
			UserId:      uid,
			QuestionIds: []int{},
			Progresses:  []models.Progress{},
		}, nil
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *UserService) AddBookmark(ctx context.Context, uid string, qid int) error {
	// 1. データを取得
	u, err := s.GetUser(ctx, uid)
	if err != nil {
		return err
	}

	// 2. データを更新
	if slices.Contains(u.QuestionIds, qid) {
		return nil
	}
	u = &models.User{
		UserId:      u.UserId,
		QuestionIds: append(u.QuestionIds, qid),
		Progresses:  u.Progresses,
	}
	if _, err := s.userRepository.BulkUpsertUser(ctx, uid, u); err != nil {
		return err
	}
	return nil
}

func (s *UserService) RemoveBookmark(ctx context.Context, uid string, qid int) error {
	// 1. データを取得
	u, err := s.GetUser(ctx, uid)
	if err != nil {
		return err
	}

	// 2. データを更新
	if !slices.Contains(u.QuestionIds, qid) {
		return models.EntityNotFoundError
	}
	qids := []int{}
	for _, id := range u.QuestionIds {
		if id != qid {
			qids = append(qids, id)
		}
	}
	u = &models.User{
		UserId:      u.UserId,
		QuestionIds: qids,
		Progresses:  u.Progresses,
	}
	if _, err := s.userRepository.BulkUpsertUser(ctx, uid, u); err != nil {
		return err
	}
	return nil
}
