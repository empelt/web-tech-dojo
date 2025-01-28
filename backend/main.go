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

	firestoreClient, err := firebaseApp.NewFirestoreClient(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	genaiClient, err := infrastructures.NewGenaiClient(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Repository
	questionRepository, err := repository.NewQuestion(firestoreClient)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Services
	service, err := services.New(genaiClient)
	if err != nil {
		e.Logger.Fatal(err)
	}

	questionService, err := services.NewQuestionService(questionRepository)
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize Handlers
	handler, err := handlers.New(service)
	if err != nil {
		e.Logger.Fatal(err)
	}

	questionHandler, err := handlers.NewQuestionHandler(questionService)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/api/chat", handler.PostQuestionAnswer)
	e.GET("/api/question/:id", questionHandler.GetQuestion)
	e.Logger.Fatal(e.Start(":" + port))
}
