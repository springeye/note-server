package server

import (
	"fmt"
	"github.com/springeye/oplin/db"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)
import "github.com/springeye/oplin/config"

type H map[string]interface{}

// example see https://github.com/go-chi/chi/blob/master/_examples/rest/main.go
func MainRouter() *chi.Mux {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))

	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, H{
			"code": 0,
			"name": "springeye",
		})
	})
	r.Mount("/user", userRouter())

	return r
}

func RunWebServer(app *chi.Mux) error {
	loggerOpts := slog.HandlerOptions{
		AddSource: true,
	}
	if config.Default.Debug {
		loggerOpts.Level = slog.DebugLevel
	} else {
		loggerOpts.Level = slog.ErrorLevel
	}
	slog.SetDefault(slog.New(loggerOpts.NewTextHandler(os.Stdout)))
	db.Setup()

	port := config.Default.Port
	println("http server run: 0.0.0.0:", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), app)
}