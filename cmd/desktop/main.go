package main

import (
	_ "embed"

	"github.com/k0kubun/pp"
	"github.com/webview/webview"
)

//go:embed start.html
var StartHtml []byte

func main() {
	w := webview.New(false)
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
