package models

import (
	"time"
)

// 問題
//
// id:        問題番号(!= Document.id), incremental
// title:     問題タイトル
// content:   問題文
// tags:      問題カテゴリ, 難易度など
// createdAt: 問題作成日時
type Question struct {
	Id        int       `json:"id" firestore:"id,omitempty"`
	Title     string    `json:"title" firestore:"title,omitempty"`
	Content   string    `json:"content" firestore:"content,omitempty"`
	Tags      []string  `json:"tags" firestore:"tags,omitempty"`
	CreatedAt time.Time `json:"createdAt" firestore:"createAt,omitempty"`
}
