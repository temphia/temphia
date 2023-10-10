package ebrowser

import (
	_ "embed"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/distro/climux"
	"github.com/webview/webview"
)

//go:embed start.html
var StartHtml []byte

func Run() {
	w := webview.New(true)
	defer w.Destroy()

	w.Bind("__goto_page__", func(ctx map[string]string) error {
		pp.Println("@ctx", ctx)
		return nil
	})

	w.SetTitle("Temphia Start")
	w.SetSize(900, 700, webview.HintNone)
	w.SetHtml(string(StartHtml))

	w.Run()

}

func init() {

	climux.Register(&climux.CliAction{
		Name: "ebrowser",
		Help: "Run embed browser with state folder",
		Func: func(args []string) error {
			Run()
			return nil
		},
	})

}
