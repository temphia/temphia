package xbprint

type Bundle struct {
	Type      string       `json:"type,omitempty"`
	Items     []BundleItem `json:"items,omitempty"`
	FormItems []FormItem   `json:"form_items,omitempty"`
}

type BundleItem struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	File      string `json:"file,omitempty"`
	Templated bool   `json:"templated,omitempty"`
}

type FormItem struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
