package xpackage

// syncme => data/sheets.ts

const (
	SheetColTypeText        = "shorttext"
	SheetColTypeLongText    = "longtext"
	SheetColTypeSelect      = "select"
	SheetColTypeMultiSelect = "multi_select"
	SheetColTypeNumber      = "number"
	SheetColTypeDate        = "datetime"
	SheetColTypeBoolean     = "bool"
	SheetColTypeRatings     = "ratings"
	SheetColTypeLocation    = "location"
	SheetColTypeFile        = "file"
	SheetColTypeUser        = "user"
	SheetColTypeReference   = "reference"
	SheetColTypeRemote      = "remote"
)

type NewSheetGroup struct {
	Name   string     `json:"name,omitempty"`
	Info   string     `json:"info,omitempty"`
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
	Options     string            `json:"opts,omitempty"`
	ExtaOptions map[string]string `json:"extra_options,omitempty"`
}
