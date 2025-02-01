package models

// お気に入り
//
// userId: ユーザId(Firebase auth Id)
type Bookmark struct {
	UserId      string `firestore:"userId"`
	QuestionIds []int  `firestore:"questionIds"`
}
