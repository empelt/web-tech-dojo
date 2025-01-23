package main

import (
	"os"

	"github.com/empelt/web-tech-dojo/handler"
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

	api := e.Group("/api")
	h := handler.NewHandler()
	h.Register(api)
	e.Logger.Fatal(e.Start(":" + port))
}
