package services

import (
	"context"
	"log"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/models"
)

func NewAuthService(firebaseAuth *infrastructures.FirebaseAuth) (*AuthService, error) {
	return &AuthService{
		firebaseAuth: firebaseAuth,
	}, nil
}

func (s *AuthService) AuthorizeAsUser(ctx context.Context, idToken string) (string, error) {
	// 1. UIDの取得
	token, err := s.firebaseAuth.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Println(err.Error())
		return "", models.EntityNotFoundError
	}
	// 2. BAN対象のユーザを弾く[TODO]

	return token.UID, nil
}
