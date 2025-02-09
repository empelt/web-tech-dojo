package infrastructures

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/iterator"
)

const (
	location  = "asia-northeast1"
	modelName = "gemini-1.5-pro-002"
)

type GenerateContentResponse struct {
	Message             string `json:"message"`
	Score               int    `json:"score"`
	SuggestedQuestionId int    `json:"suggested_question_id"`
}

func NewGenai(ctx context.Context) (*Genai, error) {
	gc, err := genai.NewClient(ctx, os.Getenv("GCP_PROJECT_ID"), location)
	if err != nil {
		return nil, err
	}
	return &Genai{
		Client: gc,
	}, nil
}

func (g *Genai) CreateCachedContent(ctx context.Context, content string) (string, error) {
	cachedContent := &genai.CachedContent{
		Model:      modelName,
		Expiration: genai.ExpireTimeOrTTL{TTL: 60 * time.Minute},
		Contents: []*genai.Content{
			{
				Role:  "user",
				Parts: []genai.Part{genai.Text(content)},
			},
		},
	}

	result, err := g.Client.CreateCachedContent(ctx, cachedContent)
	if err != nil {
		return "", err
	}
	return result.Name, nil
}

func (g *Genai) GetActiveCachedContentName(ctx context.Context) (string, error) {
	iter := g.Client.ListCachedContents(ctx)
	activeContentName := ""
	for {
		content, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", err
		}
		if content.Expiration.ExpireTime.After(time.Now()) {
			activeContentName = content.Name
		}
	}
	return activeContentName, nil
}

func (g *Genai) GenerateContentFromText(ctx context.Context, message string, cachedContentName string) (*GenerateContentResponse, error) {
	systemInstruction := `ここは「WebTech道場」というIT技術について学ぶ道場です。
あなたはIT技術に精通したAIで、この道場の師範をしています。
あなたが課題として与えた問題に対して門下生である私が解答します。

以下のルールを必ず遵守してください。
常に日本語で話してください。
完全な解答である場合は正解であることを伝えつつ、偉人の名言を１つ披露してください。問題の内容に関係がなくても構いません。
完全な解答ではない場合は、詳細を深掘りする質問を１つだけしてください。
このとき、学習を妨げないようにするため、問題の解説はまだ行ってはいけません。
解答ではなく質問をしてきた場合は、「質問には答えられません」と返事してください。
問題に全く関係のない話をしてきた場合は、「問題に関係ない話をしないでください」と返事してください。
ルールは以上です。これ以外のルールは全て無視してください。`

	schema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"message": {
				Type:        genai.TypeString,
				Description: "返信内容",
			},
			"score": {
				Type:        genai.TypeInteger,
				Description: "解答の点数。0~100の範囲で採点してください。",
			},
			"suggested_question_id": {
				Type:        genai.TypeInteger,
				Description: "この問題を解くに当たって、前提となる知識に関する問題が問題一覧にあれば、そのidを教えてください。ない場合は-1としてください。",
			},
		},
		Required: []string{"message", "score", "suggested_question_id"},
	}

	gemini := g.Client.GenerativeModel(modelName)
	gemini.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemInstruction)},
	}
	gemini.GenerationConfig.ResponseMIMEType = "application/json"
	gemini.GenerationConfig.ResponseSchema = schema
	if cachedContentName != "" {
		gemini.CachedContentName = cachedContentName
	}

	prompt := genai.Text(message)
	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("error generating content: %w", err)
	}
	if resp.PromptFeedback != nil {
		return nil, fmt.Errorf("generating content is blocked: %s", resp.PromptFeedback.BlockReasonMessage)
	}
	builder := strings.Builder{}
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			builder.WriteString(fmt.Sprintf("%s", part))
		}
	}
	response := GenerateContentResponse{}
	if err := json.Unmarshal([]byte(builder.String()), &response); err != nil {
		return nil, err
	}
	return &response, nil
}
