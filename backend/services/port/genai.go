package port

import (
	"context"

	"github.com/empelt/web-tech-dojo/infrastructures"
)

type GenaiClient interface {
	CreateCachedContent(ctx context.Context, content string) (string, error)
	GetActiveCachedContentName(ctx context.Context) (string, error)
	GenerateContentFromText(ctx context.Context, message string, cachedContentName string) (*infrastructures.GenerateContentResponse, error)
}
