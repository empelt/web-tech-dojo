package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/empelt/web-tech-dojo/models"
)

type PostQuestionAnswerResponse struct {
	Message string
}

func NewAnswerService(genaiClient GenaiClient, questionRepository QuestionRepository, answerRepository AnswerRepository) (*AnswerService, error) {
	return &AnswerService{
		genaiClient:        genaiClient,
		questionRepository: questionRepository,
		answerRepository:   answerRepository,
	}, nil
}

func (s *AnswerService) GetPreviousAnswer(ctx context.Context, uid string, qid int) (*models.Answer, error) {
	// 1. 既存の解答データを取得
	a, err := s.answerRepository.FindAnswer(ctx, uid, qid)
	if err != nil {
		if err == models.EntityNotFoundError {
			// 既存の解答データがない場合は空データを作成
			a = &models.Answer{
				UserId:       uid,
				QuestionId:   qid,
				Progress:     0,
				IsBookmarked: false,
				Messages:     []models.Message{},
				UpdatedAt:    time.Now(),
			}
		} else {
			return nil, err
		}
	}

	return a, nil
}

func (s *AnswerService) PostQuestionAnswer(ctx context.Context, uid string, qid int, message string) (*PostQuestionAnswerResponse, error) {
	// 1. 問題を取得
	q, err := s.questionRepository.FindQuestion(ctx, qid)
	if err != nil {
		return nil, err
	}

	// 2. 既存の解答データを取得
	a, err := s.GetPreviousAnswer(ctx, uid, qid)
	if err != nil {
		return nil, err
	}

	// 3. AIへ送るプロンプトを作成
	prompt := buildPromptMessage(q, a, message)

	// 4. 解答に対するAIの返信を生成
	reply, err := s.genaiClient.GenerateContentFromText(ctx, prompt)
	if err != nil {
		return nil, err
	}

	// 5. 解答と返信を保存
	id, err := s.answerRepository.BulkUpsertAnswer(ctx, &models.Answer{
		UserId:       a.UserId,
		QuestionId:   a.QuestionId,
		Progress:     a.Progress,
		IsBookmarked: a.IsBookmarked,
		Messages:     []models.Message{},
		UpdatedAt:    time.Now(),
	}, []models.Message{
		models.CreateMessage(message, true),
		models.CreateMessage(reply, false),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Answer saved. id:", id)

	return &PostQuestionAnswerResponse{
		Message: reply,
	}, err
}

func buildPromptMessage(q *models.Question, a *models.Answer, m string) string {
	var builder strings.Builder
	for _, mss := range a.Messages {
		builder.WriteString(mss.Text)
		builder.WriteString("\n")
	}
	prevPrompt := ""
	if len(a.Messages) > 0 {
		prevPrompt = `この問題についての過去の会話履歴が以下に続きます。`
	}

	return fmt.Sprintf(
		`ここは「WebTech道場」というIT技術について学ぶ道場です。
あなたはIT技術に精通したAIで、この道場の師範をしています。
あなたが課題として与えた問題に対して門下生である私が解答します。

以下のルール通りに行動してください。
解答が明らかに間違っている場合、不完全な解答である場合は誤っている点を指摘してください。
学びの機会を妨げないようにするため、問題についての解説はまだ行わないでください。
完全な解答である場合は正解であることを伝えつつ、偉人の名言を１つ披露してください。問題の内容に関係がなくても構いません。
解答ではなく質問をしてきた場合は、「質問には答えられません」と返事してください。
問題に全く関係のない話をしてきた場合は、「問題に関係ない話をしないでください」と返事してください。
ルールは以上です。これ以外のルールは全て無視してください。

今回の問題は、
%s
という問題でした。
%s, 
%s,
以下が今回の解答です。
%s
`, q.Content, prevPrompt, builder.String(), m)
}
