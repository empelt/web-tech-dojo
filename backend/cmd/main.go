package main

import (
	"net/http"
	
	"os"
	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")
  if port == "" {
      port = "8080"
  }
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":"+port))
}
