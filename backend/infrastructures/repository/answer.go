package repository

import (
	"context"
	"time"

	"google.golang.org/api/iterator"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

type AnswerDocument struct {
	UserId     string    `firestore:"userId,omitempty"`
	QuestionId int       `firestore:"questionId,omitempty"`
	UpdatedAt  time.Time `firestore:"updatedAt,omitempty"`
}

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

	var a AnswerDocument
	if err = doc.DataTo(&a); err != nil {
		return nil, err
	}
	// 2. Messagesを取得
	mssItr := doc.Ref.Collection(r.subCollectionName).Documents(ctx)
	mSnapshots, err := mssItr.GetAll()
	if err != nil {
		return nil, err
	}
	mss := []models.Message{}
	for i := 0; i < len(mSnapshots); i++ {
		var m models.Message
		mSnapshots[i].DataTo(&m)
		mss = append(mss, m)
	}

	return &models.Answer{
		UserId:     a.UserId,
		QuestionId: a.QuestionId,
		Messages:   mss,
		UpdatedAt:  a.UpdatedAt,
	}, nil
}

func (r *AnswerRepository) UpsertAnswer(ctx context.Context, a *models.Answer, mss []models.Message) (string, error) {
	aDoc := AnswerDocument{
		UserId:     a.UserId,
		QuestionId: a.QuestionId,
		UpdatedAt:  a.UpdatedAt,
	}

	// 1. 既存データ存在確認
	itr := r.firestore.Client.Collection(r.collectionName).
		Where("userId", "==", a.UserId).
		Where("questionId", "==", a.QuestionId).
		Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		// 2.1 新規作成するケース
		newDoc, _, err := r.firestore.Client.Collection(r.collectionName).Add(ctx, aDoc)
		if err != nil {
			return "", err
		}
		for i := 0; i < len(mss); i++ {
			_, _, err := newDoc.Collection(r.subCollectionName).Add(ctx, mss[i])
			if err != nil {
				return "", err
			}
		}
		return newDoc.ID, nil
	}
	if err != nil {
		return "", err
	}

	// 2.2 既存データがあるケース
	if _, err := doc.Ref.Set(ctx, aDoc); err != nil {
		return "", nil
	}
	for i := 0; i < len(mss); i++ {
		_, _, err := doc.Ref.Collection(r.subCollectionName).Add(ctx, mss[i])
		if err != nil {
			return "", err
		}
	}
	return doc.Ref.ID, nil
}
