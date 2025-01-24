package models

type PostChatMessageRequest struct {
	Message string `json:"message" validate:"required"`
}

type PostChatMessageResponse struct {
	Message string `json:"message"`
}
