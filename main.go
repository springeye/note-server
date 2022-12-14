package main

import (
	"errors"
	"os"
)
import cli "github.com/urfave/cli/v2"

// @title           Note Server API
// @version         1.0
// @description     Note Server API
// @termsOfService  https://github.com/springeye

// @contact.name   API Support
// @contact.url    https://github.com/springeye/note-server
// @contact.email  henjue@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// // @host petstore.swagger.io
// @BasePath      /

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Description for what is this security definition being used

func main() {
	application := InitApplication()
	// see https://cli.urfave.org/v2/getting-started/
	app := &cli.App{
		EnableBashCompletion: true,
		Suggest:              true,
		Before: func(context *cli.Context) error {
			return application.init()
		},
		Action: func(context *cli.Context) error {
			if context.NArg() > 0 {
				return errors.New("args error")
			}
			return application.start()
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Value:       "config.json",
				Aliases:     []string{"c"},
				Usage:       "Load app config from `FILE`",
				EnvVars:     []string{"OPLIN_CONFIG"},
				DefaultText: "config.json",
				Action: func(context *cli.Context, s string) error {
					return application.conf.ReInit(s)
				},
			},
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
					return application.server.GenDoc()
				},
			}, {
				Name:  "user",
				Usage: "User add/update/delete/disable/enable",

				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Usage:     "Show user list",
						HideHelp:  true,
						UsageText: "oplin user list",
						Action: func(context *cli.Context) error {
							return application.cmd.ListUser(context)
						},
					},
					{
						Name:      "add",
						Usage:     "Add a user",
						HideHelp:  true,
						UsageText: "oplin user add <username> <password>",
						Action: func(context *cli.Context) error {
							return application.cmd.AddUser(context)
						},
					},
					{
						Name:      "delete",
						Usage:     "Delete a user",
						UsageText: "oplin user delete <username>",
						HideHelp:  true,
						Action: func(context *cli.Context) error {
							return application.cmd.DeleteUser(context)
						},
					},
					{
						Name:      "password",
						Usage:     "Set new password for user",
						UsageText: "oplin user password <username> <password>",
						HideHelp:  true,
						Action: func(context *cli.Context) error {
							return application.cmd.SetPassword(context)
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
