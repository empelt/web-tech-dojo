package port

import (
	"context"

	"github.com/empelt/web-tech-dojo/models"
)

type UserRepository interface {
	GetUser(ctx context.Context, uid string) (*models.User, error)
	UpsertUser(ctx context.Context, uid string, u *models.User) (string, error)
}
