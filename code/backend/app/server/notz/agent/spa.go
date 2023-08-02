package agent

import (
	"bytes"
	"fmt"
	"strings"
)

type SpaBuilderOptions struct {
	Plug         string                 `json:"plug"`
	Agent        string                 `json:"agent"`
	APIBaseURL   string                 `json:"api_base_url"`
	EntryName    string                 `json:"entry"`
	ExecLoader   string                 `json:"exec_loader"`
	TenantID     string                 `json:"tenant_id"`
	JsPlugScript string                 `json:"js_plug_script"`
	StyleFile    string                 `json:"style_file"`
	ExtScripts   map[string]interface{} `json:"ext_scripts,omitempty"`
}

type SpaBuilder struct {
	opts        SpaBuilderOptions
	buf         bytes.Buffer
	linkScripts []linkType
	linkStyles  []string
}

type linkType struct {
	link string
	mod  bool
}

func (tb *SpaBuilder) Build() []byte {
	fmt.Println("@building_using (TemplatedBuilder)")

	tb.buildLoaderScripts()
	tb.buildLoaderStyles()
	tb.buildMainScripts()
	tb.buildMainStyles()
	tb.join()

	return (tb.buf.Bytes())
}

func (tb *SpaBuilder) buildLoaderScripts() {
	s := fmt.Sprintf("%s/engine/plug/%s/agent/%s/executor/%s/loader.js",
		tb.opts.APIBaseURL, tb.opts.Plug, tb.opts.Agent, tb.opts.ExecLoader)
	tb.linkScripts = append(tb.linkScripts, struct {
		link string
		mod  bool
	}{
		link: s,
		mod:  false,
	})
}

func (tb *SpaBuilder) buildLoaderStyles() {
	s := fmt.Sprintf("%s/engine/plug/%s/agent/%s/executor/%s/loader.css",
		tb.opts.APIBaseURL, tb.opts.Plug, tb.opts.Agent, tb.opts.ExecLoader)
	tb.linkStyles = append(tb.linkStyles, s)
}

func (tb *SpaBuilder) buildMainScripts() {
	scripts := strings.Split(tb.opts.JsPlugScript, ",")
	for _, script := range scripts {
		isMod := strings.HasSuffix(script, ".mjs")

		if strings.HasPrefix(script, "http://") || strings.HasPrefix(script, "https://") {
			tb.linkScripts = append(tb.linkScripts, linkType{
				link: script,
				mod:  isMod,
			})
		} else if strings.HasPrefix(script, "//lib") {
			link := fmt.Sprintf("/z/assets/lib%s", strings.ReplaceAll(script, "//lib", ""))
			tb.linkScripts = append(tb.linkScripts, linkType{
				link: link,
				mod:  isMod,
			})
		} else {
			s := fmt.Sprintf("%s/engine/plug/%s/agent/%s/serve/%s",
				tb.opts.APIBaseURL, tb.opts.Plug, tb.opts.Agent, script)
			tb.linkScripts = append(tb.linkScripts, linkType{
				link: s,
				mod:  isMod,
			})
		}
	}
}

func (tb *SpaBuilder) buildMainStyles() {
	styles := strings.Split(tb.opts.StyleFile, ",")
	for _, style := range styles {
		if strings.HasPrefix(style, "http://") || strings.HasPrefix(style, "https://") {
			tb.linkStyles = append(tb.linkStyles, style)
		} else if strings.HasPrefix(style, "//lib") {
			link := fmt.Sprintf("/z/assets/lib%s", strings.ReplaceAll(style, "//lib", ""))
			tb.linkStyles = append(tb.linkStyles, link)
		} else {
			s := fmt.Sprintf("%s/engine/plug/%s/agent/%s/serve/%s",
				tb.opts.APIBaseURL, tb.opts.Plug, tb.opts.Agent, style)
			tb.linkStyles = append(tb.linkStyles, s)
		}
	}
}

func (tb *SpaBuilder) join() {

	loaderOptions := tb.loaderOptions()

	tb.buf.Write([]byte(`
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
	`))

	tb.buf.WriteString(fmt.Sprintf("<title>%s [ %s ]</title>", tb.opts.Plug, tb.opts.Agent))
	tb.buf.WriteByte('\n')
	tb.buf.Write([]byte(`<script>window["__loader_options__"] = `))
	tb.buf.WriteString(loaderOptions)
	tb.buf.Write([]byte("</script> \n <script>"))

	// tb.buf.Write(``)

	tb.buf.Write([]byte("</script> \n"))

	for _, script := range tb.linkScripts {
		if script.mod {
			tb.buf.WriteString(fmt.Sprintf(`<script type="module" src="%s"></script>`, script.link))
			tb.buf.WriteByte('\n')
		} else {
			tb.buf.WriteString(fmt.Sprintf(`<script src="%s"></script>`, script.link))
			tb.buf.WriteByte('\n')
		}
	}

	for _, style := range tb.linkStyles {
		tb.buf.WriteString(fmt.Sprintf(`<link href="%s" rel="stylesheet"></link>`, style))
		tb.buf.WriteByte('\n')
	}

	tb.buf.Write([]byte(`</head>
	<body>
	<div id="plugroot" style="height:100vh;"></div>
	</body>
	</html>`))

}

func (tb *SpaBuilder) loaderOptions() string {
	return string(`{}`)
}
