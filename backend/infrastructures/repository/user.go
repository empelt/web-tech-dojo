package repository

import (
	"context"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
	"google.golang.org/api/iterator"
)

func NewUserRepository(firestore *infrastructures.Firestore) (*UserRepository, error) {
	return &UserRepository{
		firestore:      firestore,
		collectionName: "users",
	}, nil
}

func (r *UserRepository) GetUser(ctx context.Context, uid string) (*models.User, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("userId", "==", uid).Documents(ctx)
	doc, err := itr.Next()

	if err == iterator.Done {
		return nil, models.EntityNotFoundError
	}
	if err != nil {
		return nil, err
	}

	b := models.User{}
	if err := doc.DataTo(&b); err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *UserRepository) BulkUpsertUser(ctx context.Context, uid string, u *models.User) (string, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("userId", "==", uid).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		// 新規作成するケース
		newDoc, _, err := r.firestore.Client.Collection(r.collectionName).Add(ctx, u)
		if err != nil {
			return "", err
		}
		return newDoc.ID, nil
	}
	if err != nil {
		return "", err
	}

	// 既存のお気に入りデータがあるケース
	if _, err := doc.Ref.Set(ctx, u); err != nil {
		return "", err
	}
	return doc.Ref.ID, nil
}
