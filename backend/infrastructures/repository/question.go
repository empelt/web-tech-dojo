package repository

import (
	"context"

	"cloud.google.com/go/firestore"
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

func (r *QuestionRepository) GetQuestions(ctx context.Context, qids []int) ([]models.Question, error) {
	itr := r.firestore.Client.Collection(r.collectionName).
		Where("questionId", "in", qids).
		Documents(ctx)
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

func (r *QuestionRepository) FilterQuestionsByTags(ctx context.Context, tags []string) ([]models.Question, error) {
	query := r.firestore.Client.Collection(r.collectionName).OrderBy("id", firestore.Asc)
	for _, tag := range tags {
		query = query.Where("tags", "array-contains", tag)
	}
	itr := query.Documents(ctx)
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
