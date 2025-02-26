package infrastructures

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
)

// firestore has a limit of 500 operations per transaction
// see https://cloud.google.com/firestore/quotas?hl=ja&utm_source=chatgpt.com#writes_and_transactions
const MaxOperationsPerTransaction = 500

func NewFirebaseApp(ctx context.Context) (*Firebase, error) {
	conf := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	return &Firebase{
		app: app,
	}, nil
}

func NewFirestore(ctx context.Context, f *Firebase) (*Firestore, error) {
	client, err := f.app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &Firestore{
		Client: client,
	}, nil
}

func NewFirebaseAuth(ctx context.Context, f *Firebase) (*FirebaseAuth, error) {
	client, err := f.app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &FirebaseAuth{
		Client: client,
	}, nil
}
