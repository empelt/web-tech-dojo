package services

import (
	"context"
	"slices"

	"github.com/empelt/web-tech-dojo/models"
)

func NewBookmarkService(bookmarkRepository BookmarkRepository) (*BookmarkService, error) {
	return &BookmarkService{
		bookmarkRepository: bookmarkRepository,
	}, nil
}

func (s *BookmarkService) GetBookmark(ctx context.Context, uid string) (*models.Bookmark, error) {
	// 1.既存のお気に入りデータを取得
	b, err := s.bookmarkRepository.GetBookmark(ctx, uid)
	if err == models.EntityNotFoundError {
		// 既存のお気に入りデータがない場合は空データを作成
		return &models.Bookmark{
			UserId:      uid,
			QuestionIds: []int{},
		}, nil
	}
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (s *BookmarkService) AddBookmark(ctx context.Context, uid string, qid int) error {
	// 1. データを取得
	b, err := s.GetBookmark(ctx, uid)
	if err != nil {
		return err
	}

	// 2. データを更新
	if slices.Contains(b.QuestionIds, qid) {
		return nil
	}
	b = &models.Bookmark{
		UserId:      b.UserId,
		QuestionIds: append(b.QuestionIds, qid),
	}
	if _, err := s.bookmarkRepository.BulkUpsertBookmark(ctx, uid, b); err != nil {
		return err
	}
	return nil
}

func (s *BookmarkService) RemoveBookmark(ctx context.Context, uid string, qid int) error {
	// 1. データを取得
	b, err := s.GetBookmark(ctx, uid)
	if err != nil {
		return err
	}

	// 2. データを更新
	if !slices.Contains(b.QuestionIds, qid) {
		return models.EntityNotFoundError
	}
	qids := []int{}
	for _, id := range b.QuestionIds {
		if id != qid {
			qids = append(qids, id)
		}
	}
	b = &models.Bookmark{
		UserId:      b.UserId,
		QuestionIds: qids,
	}
	if _, err := s.bookmarkRepository.BulkUpsertBookmark(ctx, uid, b); err != nil {
		return err
	}
	return nil
}
