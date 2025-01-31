package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

// 🔹 Firestore の Question モデル（main.go に直接定義）
type Question struct {
	ID        string    `firestore:"id,omitempty"`
	Title     string    `firestore:"title"`
	Content   string    `firestore:"content"`
	Tags      []string  `firestore:"tags"`
	CreatedAt time.Time `firestore:"createdAt"`
}

// 🔹 Firestore クライアント
type FirestoreQuestionRepository struct {
	Client *firestore.Client
}

func NewFirestoreClient(ctx context.Context, projectID string) (*FirestoreQuestionRepository, error) {
	emulatorHost := os.Getenv("FIRESTORE_EMULATOR_HOST")
	if emulatorHost == "" {
		return nil, fmt.Errorf("FIRESTORE_EMULATOR_HOST が設定されていません。エミュレータを起動してください")
	}

	client, err := firestore.NewClient(ctx, projectID, option.WithoutAuthentication())
	if err != nil {
		return nil, fmt.Errorf("Firestore クライアントの作成に失敗しました: %v", err)
	}

	fmt.Println("✅ Firestore クライアントの作成に成功")
	return &FirestoreQuestionRepository{Client: client}, nil
}

// 🔹 Firestore に `Question` を追加
func (r *FirestoreQuestionRepository) AddQuestion(ctx context.Context, question *Question) error {
    docRef := r.Client.Collection("questions").NewDoc() // ✅ Firestore の新しいドキュメントを作成（ID は自動生成）
    question.ID = docRef.ID // ✅ Go の構造体にも Firestore の自動生成 ID をセット
    question.CreatedAt = time.Now()

    _, err := docRef.Set(ctx, question) // ✅ Firestore にデータを保存
    if err != nil {
        return fmt.Errorf("データの保存に失敗しました: %v", err)
    }
    fmt.Println("✅ Firestore に問題を追加しました:", question.ID)
    return nil
}

// 🔹 Firestore から `Question` を取得
func (r *FirestoreQuestionRepository) GetQuestion(ctx context.Context, id string) (*Question, error) {
    if id == "" {
        return nil, fmt.Errorf("❌ 無効な ID: 空の ID が渡されました")
    }

    doc, err := r.Client.Collection("questions").Doc(id).Get(ctx)
    if err != nil {
        return nil, fmt.Errorf("データの取得に失敗しました: %v", err)
    }

    var question Question
    if err := doc.DataTo(&question); err != nil {
        return nil, fmt.Errorf("データの変換に失敗しました: %v", err)
    }

    fmt.Println("✅ Firestore から問題を取得しました:", question.ID)
    return &question, nil
}


func main() {
	ctx := context.Background()

	// Firestore クライアントの作成
	firestoreClient, err := NewFirestoreClient(ctx, "dummy-project-id")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer firestoreClient.Client.Close()

	// 🔹 Firestore に問題を追加
	question := &Question{
		Title:   "Goの構造体について",
		Content: "Goの構造体とは何か説明してください。",
		Tags:    []string{"Go", "Struct", "Programming"},
	}
	err = firestoreClient.AddQuestion(ctx, question)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 🔹 Firestore から問題を取得
	retrievedQuestion, err := firestoreClient.GetQuestion(ctx, question.ID)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 取得したデータを表示
	log.Printf("🎉 Firestore から取得した問題: %+v\n", retrievedQuestion)
}
