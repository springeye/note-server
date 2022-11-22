package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

// example see https://github.com/go-chi/chi/blob/master/_examples/rest/main.go
func main() {
	flag.Parse()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, "\"welcome\"")
	})
	if *routes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/springeye/note-server",
			Intro:       "Welcome to the note-server generated docs.",
		}))
		return
	}

	http.ListenAndServe(":3000", r)
}
