package infrastructures

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

const (
	location  = "asia-northeast1"
	modelName = "gemini-1.5-flash-001"
)

func NewGenaiClient(ctx context.Context) (*GenaiClient, error) {
	gc, err := genai.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), location)
	if err != nil {
		return nil, err
	}
	return &GenaiClient{
		genaiClient: gc,
	}, nil
}


func (g *GenaiClient) GenerateContentFromText(ctx context.Context, message string) (string, error) {
	gemini := g.genaiClient.GenerativeModel(modelName)
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
