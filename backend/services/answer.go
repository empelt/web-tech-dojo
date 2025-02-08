package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/empelt/web-tech-dojo/models"
	"github.com/empelt/web-tech-dojo/services/port"
)

type PostQuestionAnswerResponse struct {
	Message string
}

type AnswerService struct {
	genaiClient        port.GenaiClient
	userRepository     port.UserRepository
	questionRepository port.QuestionRepository
	answerRepository   port.AnswerRepository
}

func NewAnswerService(genaiClient port.GenaiClient, userRepository port.UserRepository, questionRepository port.QuestionRepository, answerRepository port.AnswerRepository) (*AnswerService, error) {
	return &AnswerService{
		genaiClient:        genaiClient,
		userRepository:     userRepository,
		questionRepository: questionRepository,
		answerRepository:   answerRepository,
	}, nil
}

func (s *AnswerService) GetPreviousAnswers(ctx context.Context, uid string, qid int) (*models.Answer, error) {
	// 1. 既存の解答データを取得
	a, err := s.answerRepository.FindAnswer(ctx, uid, qid)
	if err != nil {
		if err == models.EntityNotFoundError {
			// 既存の解答データがない場合は空データを作成
			a = &models.Answer{
				UserId:     uid,
				QuestionId: qid,
				Messages:   []models.Message{},
				UpdatedAt:  time.Now(),
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
	a, err := s.GetPreviousAnswers(ctx, uid, qid)
	if err != nil {
		return nil, err
	}

	// 3. AIへ送るプロンプトを作成
	prompt := buildPromptMessage(q, a, message)

	// 4. 解答に対するAIの返信を生成
	response, err := s.genaiClient.GenerateContentFromText(ctx, prompt)
	if err != nil {
		return nil, err
	}

	// 5. データを保存
	// 5.1 解答と返信を保存
	m := a.Messages
	m = append(m, models.CreateMessage(message, true, models.MessageParams{
		Score:              response.Score,
		SugestedQuestionId: -1,
	}))
	m = append(m, models.CreateMessage(response.Message, false, models.MessageParams{
		Score:              0,
		SugestedQuestionId: -1,
	}))
	if _, err := s.answerRepository.UpsertAnswer(ctx, &models.Answer{
		UserId:     a.UserId,
		QuestionId: a.QuestionId,
		Messages:   m,
		UpdatedAt:  time.Now(),
	}); err != nil {
		return nil, err
	}

	// 5.2 進行状況を保存
	u, err := s.userRepository.GetUser(ctx, uid)
	if err != nil {
		if err == models.EntityNotFoundError {
			u = &models.User{
				UserId:      uid,
				QuestionIds: []int{},
				Progresses:  []models.Progress{},
			}
		} else {
			return nil, err
		}
	}
	hasProgress := false
	needUpsert := true
	for i := range u.Progresses {
		if u.Progresses[i].QuestionId == qid {
			if u.Progresses[i].Progress < response.Score {
				u.Progresses[i].Progress = response.Score
			} else {
				needUpsert = false
			}
			hasProgress = true
			break
		}
	}
	if !hasProgress {
		u.Progresses = append(u.Progresses, models.Progress{
			QuestionId: qid,
			Progress:   response.Score,
		})
	}
	if needUpsert {
		if _, err := s.userRepository.UpsertUser(ctx, uid, u); err != nil {
			return nil, err
		}
	}

	return &PostQuestionAnswerResponse{
		Message: response.Message,
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
		`
今回の問題は、
%s
という問題でした。
%s, 
%s,
以下が今回の解答です。
%s
`, q.Content, prevPrompt, builder.String(), m)
}
