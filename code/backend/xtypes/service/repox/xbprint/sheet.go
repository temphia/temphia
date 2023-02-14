package xbprint

// syncme => data/sheets.ts

const (
	SheetColTypeText          = "shorttext"
	SheetColTypeLongText      = "longtext"
	SheetColTypeNumber        = "number"
	SheetColTypeDate          = "datetime"
	SheetColTypeBoolean       = "bool"
	SheetColTypeRatings       = "ratings"
	SheetColTypeLocation      = "location"
	SheetColTypeFile          = "file"
	SheetColTypeReferenceNum  = "ref_text"
	SheetColTypeReferenceText = "ref_number"
	SheetColTypeRemoteText    = "remote_text"
	SheetColTypeRemoteNum     = "remote_number"
)

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
	Options     string            `json:"opts,omitempty"`
	ExtaOptions map[string]string `json:"extra_options,omitempty"`
}
