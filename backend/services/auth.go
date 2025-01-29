package services

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/empelt/web-tech-dojo/models"
)

func NewAuthService(firebaseAuthClient auth.Client) (*AuthService, error) {
	return &AuthService{
		firebaseAuthClient: firebaseAuthClient,
	}, nil
}

func (s *AuthService) AuthorizeAsUser(ctx context.Context, idToken string) (string, error) {
	// 1. UIDの取得
	token, err := s.firebaseAuthClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", models.EntityNotFoundError
	}
	// 2. BAN対象のユーザを弾く[TODO]

	return token.UID, nil
}
