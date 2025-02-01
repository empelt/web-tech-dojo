package repository

import (
	"github.com/empelt/web-tech-dojo/infrastructures"
)

type QuestionRepository struct {
	firestore      *infrastructures.Firestore
	collectionName string
}

type AnswerRepository struct {
	firestore         *infrastructures.Firestore
	collectionName    string
	subCollectionName string
}
