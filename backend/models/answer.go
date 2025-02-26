package models

import (
	"time"
)

// 解答
//
// userId:       ユーザId(Firebase auth Id)
// questionId:   問題番号(Question.Id)
// progress:     この問題の現在の最高得点
// Messages:     メッセージ
// updatedAt:    最終更新日時
type Answer struct {
	UserId     string    `firestore:"userId,omitempty"`
	QuestionId int       `firestore:"questionId,omitempty"`
	Messages   []Message `firestore:"messages,omitempty"`
	UpdatedAt  time.Time `firestore:"updatedAt,omitempty"`
}

// メッセージ
//
// text:      メッセージ内容
// sentByAI:  送信者がユーザか否か
// params:    詳細データ
// createdAt: 送信日時
type Message struct {
	Text       string        `firestore:"text,omitempty"`
	SentByUser bool          `firestore:"sentByUser,omitempty"`
	Params     MessageParams `firestore:"params,omitempty"`
	CreatedAt  time.Time     `firestore:"createdAt,omitempty"`
}

// メッセージ詳細データ
//
// progress:            解答の点数
// suggestedQuestionId: この問題を解く前に解くべきと提案されている問題のId
type MessageParams struct {
	Score              int `firestore:"score,omitempty"`
	SugestedQuestionId int `firestore:"suggestedQuestion,omitempty"`
}

func CreateMessage(m string, sentByUser bool, param MessageParams) Message {
	return Message{
		Text:       m,
		SentByUser: sentByUser,
		Params:     param,
		CreatedAt:  time.Now(),
	}
}
