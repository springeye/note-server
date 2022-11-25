package main

import (
	"github.com/springeye/note-server/db"
	"golang.org/x/exp/slog"
	"net/http"
    "os"
    "time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

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

	return r
}
func RunWebServer(app *chi.Mux) error {
    loggerOpts := slog.HandlerOptions{
        AddSource: true,
        }
        slog.SetDefault(slog.New(loggerOpts.NewTextHandler(os.Stdout)))
    db.Setup()
	user := db.User{}
	db.Connection.Model(&user)
	slog.Info("http server run: 0.0.0.0:3000")
	return http.ListenAndServe(":3000", app)
}
