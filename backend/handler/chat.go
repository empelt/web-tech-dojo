package handler

import (
	"fmt"
	"net/http"
	"context"
	"encoding/json"
	"os"
	"io"

	"github.com/labstack/echo/v4"

	"cloud.google.com/go/vertexai/genai"
)


// -- [ PostChatMessage ] ------------------------------
type postChatMessageRequest struct {
	Message string `json:"message" validate:"required"`
}

type postChatMessageResponse struct {
	Message string `json:"message"`
}

func (h *Handler) PostChatMessage(c echo.Context) error {
	params := &postChatMessageRequest{}
	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		fmt.Print(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	geminiResponse, err := generateContentFromText(params.Message, os.Stdout, os.Getenv("GCP_PROJECT_ID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, postChatMessageResponse{Message: geminiResponse})
}

func generateContentFromText(message string, w io.Writer, projectID string) (string, error) {
	location := "us-central1"
	modelName := "gemini-1.5-flash-001"

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return "", fmt.Errorf("error creating client: %w", err)
	}
	gemini := client.GenerativeModel(modelName)
	prompt := genai.Text(message)

	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("error generating content: %w", err)
	}
	// See the JSON response in
	// https://pkg.go.dev/cloud.google.com/go/vertexai/genai#GenerateContentResponse.
	rb, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return "", fmt.Errorf("json.MarshalIndent: %w", err)
	}
	return string(rb), nil
}
