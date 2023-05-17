package easypage

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/temphia/temphia/code/backend/engine/modules/pstate"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type Page struct {
	Name     string `json:"name,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Type     string `json:"type,omitempty"`
	Data     string `json:"data,omitempty"`
	Template string `json:"template,omitempty"`
}

type EasyPageBuilder struct {
	parent       *EasyPage
	buildCtxData map[string]any
	Pages        map[string]*Page

	buildInProgress map[string]bool

	root        *template.Template
	buildFolder string
}

func (e *EasyPage) build() (any, error) {
	buildFolder, err := os.MkdirTemp("", "adapter*")
	if err != nil {
		return nil, easyerr.Wrap("cannot create build folder", err)
	}

	// rc := rcagent.NewRCAgent(e.tenantId, e.editorHook.PlugId, e.editorHook.AgentId, e.corehub)

	psmod := pstate.New(e.tenantId, e.editorHook.PlugId, e.pkv)
	psmod.Query(&store.PkvQuery{})

	builder := EasyPageBuilder{
		parent:          e,
		buildCtxData:    make(map[string]any),
		root:            template.New("easypage").Funcs(sprig.FuncMap()),
		buildFolder:     buildFolder,
		buildInProgress: make(map[string]bool),
		Pages:           make(map[string]*Page), // fixme => load this
	}

	err = builder.Build()
	if err != nil {
		return nil, err
	}

	return map[string]any{"ok": true}, nil
}

func (eb EasyPageBuilder) Build() error {

	for _, page := range eb.Pages {

		if page.Type == "template" {
			continue
		}

		err := eb.build(page)

		if err != nil {
			return err
		}
	}

	return nil
}

func (eb *EasyPageBuilder) buildInner(page *Page) (string, error) {

	if page.Template == "" || eb.Pages[page.Template] == nil {
		return page.Data, nil
	}

	if eb.buildInProgress[page.Slug] {
		return "", easyerr.Error("Nested template yield detected")
	}
	eb.buildInProgress[page.Slug] = true
	defer func() {
		delete(eb.buildInProgress, page.Slug)
	}()

	var buf strings.Builder

	tplstr := eb.Pages[page.Template].Data

	tpl, err := eb.root.Clone()
	if err != nil {
		return "", err
	}

	tpl.Funcs(template.FuncMap{
		"yield_content": func() string {
			return page.Data
		},

		"yield_block": func(bname string) (string, error) {
			nextPage := eb.Pages[bname]
			if nextPage == nil {
				return "", easyerr.NotFound("yield block")
			}

			return eb.buildInner(nextPage)
		},
	})

	tpl, err = tpl.Parse(tplstr)
	if err != nil {
		return "", easyerr.Wrap("err parsing template", err)
	}

	bctx := BuildContext{
		Data:    eb.buildCtxData,
		Current: page.Name,
		Type:    page.Type,
	}

	err = tpl.Execute(&buf, &bctx)
	if err != nil {
		return "", easyerr.Wrap("err building template", err)
	}

	return buf.String(), nil

}

func (eb *EasyPageBuilder) build(page *Page) error {

	file, err := os.Create(path.Join(eb.buildFolder, fmt.Sprintf("%s_%s", page.Type, page.Slug)))
	if err != nil {
		return easyerr.Wrap("err creating file for page", err)
	}

	defer file.Close()

	oustr, err := eb.buildInner(page)
	if err != nil {
		return err
	}

	_, err = file.WriteString(oustr)
	return err

}

type BuildContext struct {
	Data    map[string]any
	Current string
	Type    string
}
