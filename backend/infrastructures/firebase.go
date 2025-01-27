package infrastructures

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func NewFirebaseApp(ctx context.Context) (*FirebaseApp, error) {
	conf := &firebase.Config{ProjectID: os.Getenv("GCP_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}
	return &FirebaseApp{
		firebaseApp: app,
	}, nil
}

func (f *FirebaseApp) NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	client, err := f.firebaseApp.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}
