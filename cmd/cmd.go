package cmd

import "github.com/springeye/oplin/db"

type Command struct {
	store *db.Store
}

func NewCommand(store *db.Store) *Command {
	return &Command{store: store}
}
