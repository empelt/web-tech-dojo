package infrastructures

import (
	"context"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/vertexai/genai"
)

const (
	location  = "asia-northeast1"
	modelName = "gemini-1.5-flash-001"
)

func NewGenai(ctx context.Context) (*Genai, error) {
	gc, err := genai.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), location)
	if err != nil {
		return nil, err
	}
	return &Genai{
		Client: gc,
	}, nil
}

func (g *Genai) GenerateContentFromText(ctx context.Context, message string) (string, error) {
	gemini := g.Client.GenerativeModel(modelName)

	prompt := genai.Text(message)
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("error generating content: %w", err)
	}
	if resp.PromptFeedback != nil {
		return "", fmt.Errorf("generating content is blocked: %s", resp.PromptFeedback.BlockReasonMessage)
	}
	builder := strings.Builder{}
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			builder.WriteString(fmt.Sprintf("%s\n", part))
		}
	}
	return builder.String(), nil
}
