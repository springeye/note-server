package server

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "golang.org/x/exp/slog"
	"net/http"
	_ "os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)
import "github.com/springeye/oplin/config"
import _ "github.com/springeye/oplin/docs"

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

	r.Get("/swagger/*", httpSwagger.Handler(
	//httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, H{
			"code": 0,
			"name": "springeye",
		})
	})
	r.Mount("/user", userRouter())

	return r
}

func RunWebServer(root chi.Router) error {

	port := config.Default.Port
	println("http server run: 0.0.0.0:", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), root)
}
