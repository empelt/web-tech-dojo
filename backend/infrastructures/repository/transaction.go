package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/services/port"
)

type TxExecutor struct {
	firestore *infrastructures.Firestore
}

func NewTxExecutor(firestore *infrastructures.Firestore) (*TxExecutor, error) {
	return &TxExecutor{
		firestore: firestore,
	}, nil
}

func (t *TxExecutor) ExecQuestionTx(ctx context.Context, f func(ctx context.Context, repo port.QuestionRepository) error) error {
	return t.firestore.Client.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		repo, err := NewQuestionRepository(t.firestore, tx)
		if err != nil {
			return err
		}
		return f(ctx, repo)
	})
}
