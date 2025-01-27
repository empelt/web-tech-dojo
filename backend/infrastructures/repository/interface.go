package repository

import (
	"cloud.google.com/go/firestore"
)

type QuestionRepository struct {
	firestoreClient *firestore.Client
	collectionName  string
}
