package infrastructures

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/vertexai/genai"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type FirebaseApp struct {
	firebaseApp *firebase.App
}

type FirestoreClient struct {
	firestoreClient *firestore.Client
}

type FirebaseAuthClient struct {
	firebaseAuthClient *auth.Client
}

type GenaiClient struct {
	genaiClient *genai.Client
}
