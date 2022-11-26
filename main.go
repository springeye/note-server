package main

import (
	"errors"
	"github.com/go-chi/docgen"
	"github.com/springeye/note-server/cmd"
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
			}, {
				Name:  "user",
				Usage: "User add/update/delete/disable/enable",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Usage:     "Show user list",
						HideHelp:  true,
						UsageText: "note-server user list",
						Action: func(context *cli.Context) error {
							return cmd.ListUser(context)
						},
					},
					{
						Name:      "add",
						Usage:     "Add a user",
						HideHelp:  true,
						UsageText: "note-server user add <username> <password>",
						Action: func(context *cli.Context) error {
							return cmd.AddUser(context)
						},
					},
					{
						Name:      "delete",
						Usage:     "Delete a user",
						UsageText: "note-server user delete <username>",
						HideHelp:  true,
						Action: func(context *cli.Context) error {
							return cmd.DeleteUser(context)
						},
					},
					{
						Name:      "password",
						Usage:     "Set new password for user",
						UsageText: "note-server user password <username> <password>",
						HideHelp:  true,
						Action: func(context *cli.Context) error {
							return cmd.SetPassword(context)
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		println(err.Error())
		os.Exit(1)
	}

}
