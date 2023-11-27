package ebrowser

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/climux"
	"github.com/temphia/temphia/code/distro"
	webview "github.com/webview/webview_go"

	"github.com/k0kubun/pp"
)

// p2p-eproxy

type EbrowserApp struct {
	webview webview.WebView
	clictx  climux.Context
}

func New(clictx climux.Context) *EbrowserApp {

	w := webview.New(true)
	w.SetSize(900, 700, webview.HintNone)

	return &EbrowserApp{
		webview: w,
		clictx:  clictx,
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

			if opts["start_instance"] == "true" {
				go func() {
					err := e.clictx.RunCLI("app", []string{"app", "start"})
					if err != nil {
						pp.Println("@cannot_start", err)
						return
					}
				}()

				time.Sleep(time.Second * 5)
			}

			pp.Println(e.NavigateLocal(distro.TemphiaConfigFile))

		case "connect_remote":

		}

	}()

}

func (e *EbrowserApp) NavigateLocal(file string) error {
	pp.Println("@nav", file)

	out, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	conf := config.Config{}
	err = toml.Unmarshal(out, &conf)
	if err != nil {
		return err
	}

	e.webview.Navigate(fmt.Sprintf("http://localhost%s/z/pages", conf.ServerPort))

	return nil

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
