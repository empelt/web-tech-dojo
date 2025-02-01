package main

import (
	"context"
	"os"

	"github.com/empelt/web-tech-dojo/handlers"
	"github.com/empelt/web-tech-dojo/infrastructures"
	"github.com/empelt/web-tech-dojo/infrastructures/repository"
	"github.com/empelt/web-tech-dojo/services"
	"github.com/empelt/web-tech-dojo/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))
	e.Validator = validator.NewValidator()

	ctx := context.Background()

	// Initialize Infrastructures
	firebaseApp, err := infrastructures.NewFirebaseApp(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	firestore, err := infrastructures.NewFirestore(ctx, firebaseApp)
	if err != nil {
		e.Logger.Fatal(err)
	}

	firebaseAuth, err := infrastructures.NewFirebaseAuth(ctx, firebaseApp)
	if err != nil {
		e.Logger.Fatal(err)
	}

	genai, err := infrastructures.NewGenai(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Repository
	questionRepository, err := repository.NewQuestionRepository(firestore)
	if err != nil {
		e.Logger.Fatal(err)
	}

	answerRepository, err := repository.NewAnswerRepository(firestore)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Services
	authService, err := services.NewAuthService(firebaseAuth)
	if err != nil {
		e.Logger.Fatal(err)
	}

	answerService, err := services.NewAnswerService(genai, questionRepository, answerRepository)
	if err != nil {
		e.Logger.Fatal(err)
	}

	questionService, err := services.NewQuestionService(questionRepository)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Handlers
	answerHandler, err := handlers.NewAnswerHandler(authService, answerService)
	if err != nil {
		e.Logger.Fatal(err)
	}

	questionHandler, err := handlers.NewQuestionHandler(questionService)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.GET("/api/question", questionHandler.SearchQuestions)
	e.GET("/api/question/:id", questionHandler.GetQuestion)
	e.GET("/api/question/:id/answer", answerHandler.GetPreviousAnswer)
	e.POST("/api/question/:id/answer", answerHandler.PostQuestionAnswer)
	e.Logger.Fatal(e.Start(":" + port))
}
