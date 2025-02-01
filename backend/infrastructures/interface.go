package infrastructures

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/vertexai/genai"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type Firebase struct {
	app *firebase.App
}

type Firestore struct {
	Client *firestore.Client
}

type FirebaseAuth struct {
	Client *auth.Client
}

type Genai struct {
	Client *genai.Client
}
