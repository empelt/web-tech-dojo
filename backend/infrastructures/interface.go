package infrastructures

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

import (
	firebase "firebase.google.com/go"
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/vertexai/genai"
)

type FirebaseApp struct {
	firebaseApp *firebase.App
}

type FirestoreClient struct {
	firestoreClient *firestore.Client
}

type GenaiClient struct {
	genaiClient *genai.Client
}
