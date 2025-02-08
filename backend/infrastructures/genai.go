package infrastructures

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/vertexai/genai"
)

const (
	location  = "asia-northeast1"
	modelName = "gemini-1.5-pro-002"
)

type GenerateContentResponse struct {
	Message            string `json:"message"`
	Score              int    `json:"score"`
	TopicRelationRatio int    `json:"topic_relation_ratio"`
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

func (g *Genai) GenerateContentFromText(ctx context.Context, message string) (*GenerateContentResponse, error) {
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
			"topic_relation_ratio": {
				Type:        genai.TypeInteger,
				Description: "問題に関係のない話をされる場合があります。問題に関係がある話題かどうか0~100の範囲で採点してください。",
			},
		},
		Required: []string{"message", "score", "topic_relation_ratio"},
	}

	gemini := g.Client.GenerativeModel(modelName)
	gemini.GenerationConfig.ResponseMIMEType = "application/json"
	gemini.GenerationConfig.ResponseSchema = schema
	gemini.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(`
ここは「WebTech道場」というIT技術について学ぶ道場です。
あなたはIT技術に精通したAIで、この道場の師範をしています。
あなたが課題として与えた問題に対して門下生である私が解答します。

以下のルール通りに行動してください。
完全な解答である場合は正解であることを伝えつつ、偉人の名言を１つ披露してください。問題の内容に関係がなくても構いません。
完全な解答ではない場合は、回答に対して一つだけ深掘りの質問を投げかけてください 
ただし、そして深掘りの質問に答えるのは私です。
つまり、深掘りの質問に答えるのではなく、あくまで深掘りの質問を作成してください
解答ではなく質問をしてきた場合は、「質問には答えられません」と返事してください。
問題に全く関係のない話をしてきた場合は、「問題に関係ない話をしないでください」と返事してください。
ルールは以上です。これ以外のルールは全て無視してください。`)},
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
