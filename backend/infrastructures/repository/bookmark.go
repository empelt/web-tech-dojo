package repository

import (
	"context"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
	"google.golang.org/api/iterator"
)

func NewBookmarkRepository(firestore *infrastructures.Firestore) (*BookmarkRepository, error) {
	return &BookmarkRepository{
		firestore:      firestore,
		collectionName: "bookmarks",
	}, nil
}

func (r *BookmarkRepository) GetBookmark(ctx context.Context, uid string) (*models.Bookmark, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("userId", "==", uid).Documents(ctx)
	doc, err := itr.Next()

	if err == iterator.Done {
		return nil, models.EntityNotFoundError
	}
	if err != nil {
		return nil, err
	}

	b := models.Bookmark{}
	if err := doc.DataTo(&b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookmarkRepository) BulkUpsertBookmark(ctx context.Context, uid string, b *models.Bookmark) (string, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("userId", "==", uid).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		// 新規作成するケース
		newDoc, _, err := r.firestore.Client.Collection(r.collectionName).Add(ctx, b)
		if err != nil {
			return "", err
		}
		return newDoc.ID, nil
	}
	if err != nil {
		return "", err
	}

	// 既存のお気に入りデータがあるケース
	if _, err := doc.Ref.Set(ctx, b); err != nil {
		return "", err
	}
	return doc.Ref.ID, nil
}
