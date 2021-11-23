package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go-echo-bbs/sqlc"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	pgx, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	q := sqlc.New(pgx)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = renderer

	/**
	ルーティング
	 */
	e.GET("/", func(context echo.Context) error {
		messages, err := q.GetThreadMessages(ctx, 100)
		if err != nil {
			e.Logger.Fatal(err)
		}
		e.Logger.Infof("messages.len=%d", len(messages))

		return context.Render(http.StatusOK, "index.html", map[string]interface{}{
			"Messages": messages,
		})
	})
	e.POST("/messages", func(context echo.Context) error {
		messageText := context.FormValue("message")
		if len(messageText) == 0 {
			e.Logger.Warn("message value is empty")
			return context.Redirect(http.StatusMovedPermanently, "/")
		}

		params := sqlc.CreateThreadMessageParams{
			Message: messageText,
			CreatedAt: time.Now(),
		}
		message, err := q.CreateThreadMessage(ctx, params)
		if err != nil {
			return err
		}
		e.Logger.Info("Creatad message=%v", message)
		return context.Redirect(http.StatusMovedPermanently, "/")
	})

	/**
	起動
	 */
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
