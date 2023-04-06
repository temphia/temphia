package pageform

import "github.com/thoas/go-funk"

type FormModel struct {
	Name     string               `json:"name,omitempty" yaml:"name,omitempty"`
	Stages   map[string]FormStage `json:"stages,omitempty" yaml:"stages,omitempty"`
	ExecHint []string             `json:"exec_hint,omitempty" yaml:"exec_hint,omitempty"`
	OnLoad   []Hook               `json:"on_load,omitempty" yaml:"on_load,omitempty"`
}

type FormStage struct {
	About    string         `json:"about,omitempty" yaml:"about,omitempty"`
	Items    []FormItem     `json:"items,omitempty" yaml:"items,omitempty"`
	OnSubmit []Hook         `json:"on_submit,omitempty" yaml:"on_submit,omitempty"` // validate data -> side_effect -> maybe_modify_data -> set_next_stage
	OnBuild  []Hook         `json:"on_build,omitempty" yaml:"on_build,omitempty"`   // load_fileds -> set_data
	Data     map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
}

type FormItem struct {
	Name     string            `json:"name,omitempty" yaml:"name,omitempty"`
	Info     string            `json:"info,omitempty" yaml:"info,omitempty"`
	Type     string            `json:"type,omitempty" yaml:"type,omitempty"`
	Options  []string          `json:"options,omitempty" yaml:"options,omitempty"`
	Pattern  string            `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	HtmlAttr map[string]string `json:"html_attr,omitempty" yaml:"html_attr,omitempty"`
	Disabled bool              `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}

func (fs *FormStage) GetItems(ignores []string) []FormItem {
	if len(ignores) == 0 {
		return fs.Items
	}

	items := make([]FormItem, 0)

	for _, fi := range fs.Items {
		if funk.ContainsString(ignores, fi.Name) {
			continue
		}
		items = append(items, fi)
	}

	return items

}

type Hook struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Target  string `json:"target,omitempty" yaml:"target,omitempty"`
	Type    string `json:"type,omitempty" yaml:"type,omitempty"`
	Options string `json:"options,omitempty" yaml:"options,omitempty"`
}

type LoadRequest struct {
	DataContextType string         `json:"data_context_type,omitempty" yaml:"data_context_type,omitempty"`
	Rows            []int64        `json:"rows,omitempty" yaml:"rows,omitempty"`
	Options         map[string]any `json:"options,omitempty" yaml:"options,omitempty"`
}

type SubmitRequest struct {
	Data  map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	Stage string         `json:"stage,omitempty" yaml:"stage,omitempty"`
}

type Response struct {
	Title   string         `json:"title,omitempty" yaml:"title,omitempty"`
	Ok      bool           `json:"ok,omitempty" yaml:"ok,omitempty"`
	Final   bool           `json:"final,omitempty" yaml:"final,omitempty"`
	Message string         `json:"message,omitempty" yaml:"message,omitempty"`
	Items   []FormItem     `json:"items,omitempty" yaml:"items,omitempty"`
	Data    map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	Stage   string         `json:"stage,omitempty" yaml:"stage,omitempty"`
}
