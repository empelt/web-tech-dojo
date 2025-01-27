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
	Id        int
	Title     string
	Content   string
	Tags      []string
	CreatedAt time.Time
}
