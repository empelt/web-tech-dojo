package repository

import (
	"cloud.google.com/go/firestore"
	"github.com/empelt/web-tech-dojo/infrastructures"
)

type QuestionRepository struct {
	firestore      *infrastructures.Firestore
	collectionName string
	tx             *firestore.Transaction
}

type AnswerRepository struct {
	firestore         *infrastructures.Firestore
	collectionName    string
	subCollectionName string
}

type UserRepository struct {
	firestore      *infrastructures.Firestore
	collectionName string
}
