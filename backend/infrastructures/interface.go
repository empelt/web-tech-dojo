package infrastructures

//go:generate mockgen -source=interface.go -destination=mock/interface.go -package=mock

import (
	"context"

	"cloud.google.com/go/vertexai/genai"
)

type GenaiClient struct {
	ctx         context.Context
	genaiClient *genai.Client
}
