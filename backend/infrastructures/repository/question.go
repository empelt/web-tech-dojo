package repository

import (
	"context"

	"google.golang.org/api/iterator"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

func NewQuestionRepository(firestore *infrastructures.Firestore) (*QuestionRepository, error) {
	return &QuestionRepository{
		firestore:      firestore,
		collectionName: "questions",
	}, nil
}

func (r *QuestionRepository) FindQuestion(ctx context.Context, id int) (*models.Question, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("id", "==", id).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		return nil, models.EntityNotFoundError
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

func (r *QuestionRepository) GetAllQuestions(ctx context.Context) ([]models.Question, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Documents(ctx)
	questions := []models.Question{}
	for {
		doc, err := itr.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var q models.Question
		doc.DataTo(&q)
		questions = append(questions, q)
	}
	return questions, nil
}
