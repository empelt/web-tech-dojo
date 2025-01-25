package main

import (
	"context"
	"os"

	"github.com/empelt/web-tech-dojo/handlers"
	"github.com/empelt/web-tech-dojo/infrastructures"
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

	genaiClient, err := infrastructures.New(ctx)
	if err != nil {
		e.Logger.Fatal(err)
	}

	service, err := services.New(genaiClient)
	if err != nil {
		e.Logger.Fatal(err)
	}

	handler, err := handlers.New(service)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.POST("/api/chat", handler.PostChatMessage)
	e.Logger.Fatal(e.Start(":" + port))
}
