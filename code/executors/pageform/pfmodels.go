package pageform

type FormModel struct {
	Items          []FormItem     `json:"items,omitempty"`
	Data           map[string]any `json:"data,omitempty"`
	Message        string         `json:"message,omitempty"`
	ServerOnLoad   string         `json:"server_onload,omitempty"`
	ServerOnSubmit string         `json:"server_onsubmit,omitempty"`
	ClientOnLoad   string         `json:"client_onload,omitempty"`
	ClientOnSubmit string         `json:"client_onsubmit,omitempty"`
}

type FormItem struct {
	Name     string            `json:"name,omitempty"`
	Info     string            `json:"info,omitempty"`
	Type     string            `json:"type,omitempty"`
	Pattern  string            `json:"pattern,omitempty"`
	HtmlAttr map[string]string `json:"html_attr,omitempty"`
}

type LoadRequest struct {
	DataContextType string         `json:"data_context_type,omitempty"`
	Rows            int64          `json:"rows,omitempty"`
	Options         map[string]any `json:"options,omitempty"`
}

type LoadResponse struct {
	Ok       bool           `json:"ok,omitempty"`
	Message  string         `json:"message,omitempty"`
	Items    []FormItem     `json:"items,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
	OnLoad   string         `json:"onload,omitempty"`
	OnSubmit string         `json:"onsubmit,omitempty"`
}

type Submit struct {
	Data map[string]any `json:"data,omitempty"`
}

type Result struct {
	Ok      bool         `json:"ok,omitempty"`
	Message string       `json:"message,omitempty"`
	Items   []ResultItem `json:"items,omitempty"`
}

type ResultItem struct {
	Type     string            `json:"type,omitempty"`
	Data     any               `json:"data,omitempty"`
	HtmlAttr map[string]string `json:"html_attr,omitempty"`
}
