package pageform

type FormModel struct {
	Name           string                `json:"name,omitempty"`
	Items          map[string][]FormItem `json:"items,omitempty"`
	Data           map[string]any        `json:"data,omitempty"`
	Message        string                `json:"message,omitempty"`
	ServerOnLoad   string                `json:"server_onload,omitempty"`   // load_fileds -> set_data
	ServerOnSubmit string                `json:"server_onsubmit,omitempty"` // validate data -> side_effect -> maybe_modify_data -> set_next_stage
	ClientOnLoad   string                `json:"client_onload,omitempty"`
	ClientOnSubmit string                `json:"client_onsubmit,omitempty"`
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
	Ok       bool           `json:"ok,omitempty"`
	Message  string         `json:"message,omitempty"`
	Items    []FormItem     `json:"items,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	OnLoad   string         `json:"onload,omitempty"`
	OnSubmit string         `json:"onsubmit,omitempty"`
	Stage    string         `json:"stage,omitempty"`
}
