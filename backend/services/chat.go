package services

import (
	"context"
)

func New(genaiClient GenaiClient) (*ChatService, error) {
	return &ChatService{
		genaiClient: genaiClient,
	}, nil
}

func (s *ChatService) PostChatMessage(ctx context.Context, message string) (string, error) {
	return s.genaiClient.GenerateContentFromText(ctx, message)
}
