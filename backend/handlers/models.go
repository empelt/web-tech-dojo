package handlers

import (
	"time"

	"github.com/empelt/web-tech-dojo/models"
)

type GetPreviousAnswerResponse struct {
	Messages []Message `json:"messages"`
}

func BuildGetPreviousAnswersReponse(a *models.Answer) GetPreviousAnswerResponse {
	mss := []Message{}
	for i := range a.Messages {
		mss = append(mss, buildMessage(a.Messages[i]))
	}
	return GetPreviousAnswerResponse{
		Messages: mss,
	}
}

type Message struct {
	Text       string        `json:"text"`
	SentByUser bool          `json:"sentByUser"`
	Params     MessageParams `json:"params"`
	CreatedAt  time.Time     `json:"createdAt"`
}

type MessageParams struct {
	Score              int `json:"score"`
	SugestedQuestionId int `json:"suggestedQuestion"`
}

func buildMessage(m models.Message) Message {
	return Message{
		Text:       m.Text,
		SentByUser: m.SentByUser,
		Params: MessageParams{
			Score:              m.Params.Score,
			SugestedQuestionId: m.Params.SugestedQuestionId,
		},
		CreatedAt: m.CreatedAt,
	}
}

type PostQuestionAnswerRequest struct {
	Answer string `json:"message" validate:"required"`
}

type PostQuestionAnswerResponse struct {
	Answer string `json:"message"`
}

type GetBookmarkResponse struct {
	QuestionIds []int `json:"questionIds"`
}
