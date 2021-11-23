package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Hello, World!")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Infof("Start Server PORT=%s", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
