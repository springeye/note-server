package gui

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/springeye/note-server/db"
	"os"
)

func ShowUserList(users []db.User) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"ID", "Username", "IsAdmin", "IsDelete"})
	var rows []table.Row
	for _, user := range users {
		rows = append(rows, table.Row{
			user.ID,
			user.Username,
			false,
			false,
		})
	}
	t.AppendRows(rows)
	t.AppendSeparator()
	t.SetStyle(table.StyleColoredBright)
	t.AppendFooter(table.Row{"", "Total", len(users)})
	t.Style().Format.Header = text.FormatDefault
	t.Render()
}
