package pageform

import "github.com/thoas/go-funk"

type FormModel struct {
	Name     string               `json:"name,omitempty"`
	Stages   map[string]FormStage `json:"stages,omitempty"`
	ExecHint []string             `json:"exec_hint,omitempty"`
	OnLoad   string               `json:"on_load,omitempty"`
}

type FormStage struct {
	About      string         `json:"about,omitempty"`
	Items      []FormItem     `json:"items,omitempty"`
	OnSubmit   string         `json:"on_submit,omitempty"`   // validate data -> side_effect -> maybe_modify_data -> set_next_stage
	OnGenerate string         `json:"on_generate,omitempty"` // load_fileds -> set_data
	Data       map[string]any `json:"data,omitempty"`
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

type FormItem struct {
	Name     string            `json:"name,omitempty"`
	Info     string            `json:"info,omitempty"`
	Type     string            `json:"type,omitempty"`
	Options  []string          `json:"options,omitempty"`
	Pattern  string            `json:"pattern,omitempty"`
	HtmlAttr map[string]string `json:"html_attr,omitempty"`
	Disabled bool              `json:"disabled,omitempty"`
}

type LoadRequest struct {
	DataContextType string         `json:"data_context_type,omitempty"`
	Rows            []int64        `json:"rows,omitempty"`
	Options         map[string]any `json:"options,omitempty"`
}

type SubmitRequest struct {
	Data  map[string]any `json:"data,omitempty"`
	Stage string         `json:"stage,omitempty"`
}

type Response struct {
	Ok      bool           `json:"ok,omitempty"`
	Message string         `json:"message,omitempty"`
	Items   []FormItem     `json:"items,omitempty"`
	Data    map[string]any `json:"data,omitempty"`
	Stage   string         `json:"stage,omitempty"`
}
