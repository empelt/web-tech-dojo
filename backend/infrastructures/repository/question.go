package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

func NewQuestionRepository(firestore *infrastructures.Firestore, tx *firestore.Transaction) (*QuestionRepository, error) {
	return &QuestionRepository{
		firestore:      firestore,
		collectionName: "questions",
		tx:             tx,
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

func (r *QuestionRepository) UpsertQuestion(ctx context.Context, question models.Question) (string, error) {
	itr := r.firestore.Client.Collection(r.collectionName).Where("id", "==", question.Id).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		newDoc, _, err := r.firestore.Client.Collection(r.collectionName).Add(ctx, question)
		if err != nil {
			return "", err
		}
		return newDoc.ID, nil
	}
	if err != nil {
		return "", err
	}

	if _, err := doc.Ref.Set(ctx, question); err != nil {
		return "", nil
	}
	return doc.Ref.ID, nil
}

// トランザクションの中で呼ぶこと
func (r *QuestionRepository) UpsertQuestionWithTx(ctx context.Context, question models.Question) error {
	itr := r.firestore.Client.Collection(r.collectionName).Where("id", "==", question.Id).Documents(ctx)
	doc, err := itr.Next()
	if err == iterator.Done {
		newDocRef := r.firestore.Client.Collection(r.collectionName).NewDoc()
		err = r.tx.Create(newDocRef, question)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	if err := r.tx.Set(doc.Ref, question); err != nil {
		return err
	}

	return nil
}
