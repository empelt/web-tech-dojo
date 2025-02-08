package port

import (
	"context"

	"github.com/empelt/web-tech-dojo/infrastructures"
)

type GenaiClient interface {
	GenerateContentFromText(ctx context.Context, message string) (*infrastructures.GenerateContentResponse, error)
}
