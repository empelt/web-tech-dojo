package handlers

import "context"

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

type ChatService interface {
	PostChatMessage(ctx context.Context, message string) (string, error)
}

type ChatHandler struct {
	chatService ChatService
}
