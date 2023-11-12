package ebrowser

import (
	_ "embed"
	"html/template"
	"strings"
)

//go:embed start.tpl.html
var StartTemplate string

type TemplateOptions struct {
	LocalExists  bool
	LocalRunning bool
	LocalFile    string
}

func RenderPage(opts TemplateOptions) (string, error) {

	tpl, err := template.New("start_page").Parse(StartTemplate)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	err = tpl.Execute(&b, &opts)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
