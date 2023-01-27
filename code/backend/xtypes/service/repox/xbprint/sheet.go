package xbprint

type NewSheetGroup struct {
	Name   string     `json:"name,omitempty"`
	Sheets []NewSheet `json:"sheets,omitempty"`
}

type NewSheet struct {
	Name     string           `json:"name,omitempty"`
	Columns  []NewSheetColumn `json:"columns,omitempty"`
	SeedData []map[string]any `json:"seed_data,omitempty"`
}

type NewSheetColumn struct {
	Name        string            `json:"name,omitempty"`
	Ctype       string            `json:"ctype,omitempty"`
	Color       string            `json:"color,omitempty"`
	ExtaOptions map[string]string `json:"extra_options,omitempty"`
}
