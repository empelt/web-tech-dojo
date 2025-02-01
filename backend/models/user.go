package models

// お気に入り
//
// userId:     ユーザId(Firebase auth Id)
// questionId: 問題Id
type User struct {
	UserId      string     `firestore:"userId"`
	QuestionIds []int      `firestore:"questionIds"`
	Progresses  []Progress `firestore:"progresses"`
}

type Progress struct {
	QuestionId int `firestore:"questionId"`
	Progress   int `firestore:"progress"`
}
