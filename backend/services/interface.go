package services

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type GenaiClient interface {
	GenerateContentFromText(ctx context.Context, message string) (string, error)
}

type ChatService struct {
	genaiClient GenaiClient
}
