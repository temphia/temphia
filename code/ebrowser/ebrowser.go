package ebrowser

import (
	_ "embed"

	webview "github.com/webview/webview_go"

	"github.com/k0kubun/pp"
)

type EbrowserApp struct {
	webview webview.WebView
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

	e.webview.SetTitle("Temphia Start")
	e.webview.SetHtml(shtml)
	e.webview.Run()

}

func (e *EbrowserApp) __BindEbrowserRPC(name string, opts map[string]string) error {
	pp.Println("@ctx", name, opts)
	return nil
}

func (e *EbrowserApp) Close() {
	e.webview.Destroy()
}
