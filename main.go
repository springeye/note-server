package main

import (
	"errors"
	"github.com/go-chi/docgen"
	"log"
	"os"
)
import cli "github.com/urfave/cli/v2"

func main() {
	r := MainRouter()
	// see https://cli.urfave.org/v2/getting-started/
	app := &cli.App{
		Action: func(context *cli.Context) error {
			return RunWebServer(r)
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"a"},
				Usage:   "Run server application in background",
				Action: func(cCtx *cli.Context) error {
					return errors.New("feature not implemented")
				},
			},
			{
				Name:  "doc",
				Usage: "Generate api documents",
				Action: func(cCtx *cli.Context) error {
					markdownRoutesDoc := docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
						ProjectPath: "github.com/springeye/note-server",
						Intro:       "Welcome to the note-server generated docs.",
					})
					println(markdownRoutesDoc)
					err := os.WriteFile("api.md", []byte(markdownRoutesDoc), 0777)
					if err != nil {
						panic(err)
					}
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
