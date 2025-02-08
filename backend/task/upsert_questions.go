package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/infrastructures/repository"
	"github.com/empelt/web-tech-dojo/models"
	"github.com/empelt/web-tech-dojo/services"
)

const questionsFileName = "task/questions.json"

type questionsJson struct {
	Questions []models.Question `json:"questions"`
}

func main() {
	ctx := context.Background()
	firebaseApp, err := infrastructures.NewFirebaseApp(ctx)
	if err != nil {
		log.Fatalf("Failed to create firebase app: %v", err)
	}
	firestore, err := infrastructures.NewFirestore(ctx, firebaseApp)
	if err != nil {
		log.Fatalf("Failed to create firestore: %v", err)
	}

	questionRepository, err := repository.NewQuestionRepository(firestore, nil)
	if err != nil {
		log.Fatalf("Failed to create question repository: %v", err)
	}
	transaction, err := repository.NewTxExecutor(firestore)
	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err)
	}
	userRepository, err := repository.NewUserRepository(firestore)
	if err != nil {
		log.Fatalf("Failed to create user repository: %v", err)
	}

	userService, err := services.NewUserService(userRepository)
	if err != nil {
		log.Fatalf("Failed to create user service: %v", err)
	}
	questionService, err := services.NewQuestionService(questionRepository, userService, transaction)
	if err != nil {
		log.Fatalf("Failed to create question service: %v", err)
	}

	file, err := os.Open(questionsFileName)
	if err != nil {
		log.Fatalf("Failed to open questions file: %v", err)
	}
	defer file.Close()

	var qj questionsJson
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&qj); err != nil {
		log.Fatalf("Failed to decode questions: %v", err)
	}

	questions := []models.Question{}
	for _, q := range qj.Questions {
		if q.Id == 0 || q.Title == "" || q.Content == "" {
			log.Fatalf("Invalid question: %v", q)
		}
		questions = append(questions, models.Question{
			Id:        q.Id,
			Title:     q.Title,
			Content:   q.Content,
			Tags:      q.Tags,
			CreatedAt: time.Now(),
		})
	}
	err = questionService.UpsertQuestions(ctx, questions)
	if err != nil {
		log.Fatalf("Failed to insert data: %v", err)
	}

	log.Println("Data upserted successfully")
}
