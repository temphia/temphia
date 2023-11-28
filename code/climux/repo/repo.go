package repo

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/tools/repobuild/cmd"
)

func init() {

	climux.Register(&climux.Action{
		Name: "repo",
		Help: "repositary related actions",
		Func: func(cctx climux.Context) error {

			PrintRadomTable()

			if cctx.Args[0] == "build" {
				return cmd.Run(".repo.toml")
			}

			return nil
		},
	})

}

func PrintRadomTable() {

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRows([]table.Row{
		{1, "Arya", "Stark", 3000},
		{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"},
	})
	t.AppendSeparator()
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	t.AppendFooter(table.Row{"", "", "Total", 10000})
	t.Render()
}
