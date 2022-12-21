package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/springeye/oplin/db"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

type Server struct {
	router chi.Router
	store  *db.Store
}

func NewServer(store *db.Store) *Server {
	return &Server{router: mainRouter(store), store: store}
}
func (s *Server) Start(port int) error {
	slog.Info(fmt.Sprintf("http server run: 0.0.0.0:%d", port))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.router)
}
func (s *Server) GenDoc() error {
	markdownRoutesDoc := docgen.MarkdownRoutesDoc(s.router, docgen.MarkdownOpts{
		ProjectPath: "github.com/springeye/oplin",
		Intro:       "Welcome to the oplin generated docs.",
	})
	println(markdownRoutesDoc)
	err := os.WriteFile("api.md", []byte(markdownRoutesDoc), 0777)
	return err
}
