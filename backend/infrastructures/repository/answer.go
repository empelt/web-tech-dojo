package repository

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

func NewAnswerRepository(firestore *infrastructures.Firestore) (*AnswerRepository, error) {
	return &AnswerRepository{
		firestore:         firestore,
		collectionName:    "answers",
		subCollectionName: "messages",
	}, nil
}

func (r *AnswerRepository) FindAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error) {
	// 1. Answerを取得
	itr := r.firestore.Client.Collection(r.collectionName).
		Where("userId", "==", uid).
		Where("questionId", "==", qid).
		Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, models.EntityNotFoundError
	}
	if err != nil {
		return nil, err
	}

	a := models.Answer{}
	if err = doc.DataTo(&a); err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AnswerRepository) UpsertAnswer(ctx context.Context, a *models.Answer) (string, error) {
	// 1. 既存データ存在確認
	itr := r.firestore.Client.Collection(r.collectionName).
		Where("userId", "==", a.UserId).
		Where("questionId", "==", a.QuestionId).
		Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		// 2.1 新規作成するケース
		newDoc, _, err := r.firestore.Client.Collection(r.collectionName).Add(ctx, a)
		if err != nil {
			return "", err
		}
		return newDoc.ID, nil
	}
	if err != nil {
		return "", err
	}

	// 2.2 既存データがあるケース
	if _, err := doc.Ref.Set(ctx, a); err != nil {
		return "", nil
	}
	return doc.Ref.ID, nil
}
