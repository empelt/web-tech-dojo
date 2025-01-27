package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/empelt/web-tech-dojo/models"
)

func NewQuestion(firestoreClient *firestore.Client) (*QuestionRepository, error) {
	return &QuestionRepository{
		firestoreClient: firestoreClient,
		collectionName:  "questions",
	}, nil
}

func (r *QuestionRepository) FindQuestion(ctx context.Context, id int) (*models.Question, error) {
	itr := r.firestoreClient.Collection(r.collectionName).Where("id", "==", id).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var q models.Question
	if err = doc.DataTo(&q); err != nil {
		return nil, err
	}
	return &q, nil
}
