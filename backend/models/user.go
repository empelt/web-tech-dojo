package models

// お気に入り
//
// userId:     ユーザId(Firebase auth Id)
// questionId: お気に入りした問題のIdリスト [memo]テーブル設計として微妙だがfirestoreの制約回避のためにここに置いた
// Progress:   各問題の進行度リスト [mempo]テーブル設計として微妙だがfirestoreの制約回避のためにここに置いた
type User struct {
	UserId      string     `firestore:"userId"`
	QuestionIds []int      `firestore:"questionIds"`
	Progresses  []Progress `firestore:"progresses"`
}

type Progress struct {
	QuestionId int `firestore:"questionId"`
	Progress   int `firestore:"progress"`
}
