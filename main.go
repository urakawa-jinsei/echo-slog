package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

func main() {
	// Create a slog logger, which:
	//   - Logs to stdout.
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				switch a.Key {
				case slog.MessageKey:
					a.Key = "message"
				case slog.LevelKey:
					a.Key = "severity"
				}
				return a
			},
		}),
	)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(slogecho.New(logger))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/info", func(c echo.Context) error {
		e.Logger.Info("info")
		return c.String(http.StatusOK, "info")
	})
	e.GET("/error", func(c echo.Context) error {
		e.Logger.Error("error")
		return c.String(http.StatusOK, "error")
	})

	// Start server
	e.Logger.Fatal(e.Start(":4242"))
}
