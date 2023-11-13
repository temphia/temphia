package ebrowser

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/temphia/temphia/code/distro/climux"
	webview "github.com/webview/webview_go"

	"github.com/k0kubun/pp"
)

type EbrowserApp struct {
	webview webview.WebView
	clictx  climux.Context
}

func New() *EbrowserApp {

	w := webview.New(true)
	w.SetSize(900, 700, webview.HintNone)

	return &EbrowserApp{
		webview: w,
	}
}

func (e *EbrowserApp) RunWithStartPage(opts TemplateOptions) {

	e.webview.Bind("__ebrowser_rpc__", e.__BindEbrowserRPC)

	shtml, err := RenderPage(opts)
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		pp.Println(e.__sendRPC("temphia_start", map[string]any{}))
	}()

	e.webview.SetTitle("Temphia Start")
	e.webview.SetHtml(shtml)
	e.webview.Run()

}

func (e *EbrowserApp) __BindEbrowserRPC(name string, opts map[string]string) {
	pp.Println("@ctx", name, opts)

	go func() {

		switch name {
		case "connect_local":

			if opts["init_instance"] == "true" {
				err := e.clictx.RunCLI("app", []string{"app", "init"})
				if err != nil {
					pp.Println("@cannot_init", err)
					return
				}
			}

			err := e.clictx.RunCLI("app", []string{"app", "start"})
			if err != nil {
				pp.Println("@cannot_start", err)
				return
			}

		case "connect_remote":

		}

	}()

}

func (e *EbrowserApp) __sendRPC(name string, data any) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	e.webview.Eval(fmt.Sprintf(`__handle_rpc__("%s", "%s" )`, name, out))
	return nil
}

func (e *EbrowserApp) Close() {
	e.webview.Destroy()
}
