package xutils

import (
	"bytes"
	"html/template"
)

type inlineTemplateOpts struct {
	Tpl  string `json:"tpl,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (x *Xutils) InlineTemplate(opts *inlineTemplateOpts) (string, error) {

	tpl, err := template.New("inline").Parse(opts.Tpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, opts.Data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
