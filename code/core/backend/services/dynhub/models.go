package dynhub

type RowMod struct {
	Table   string      `json:"table,omitempty"`
	Rows    []int64     `json:"rows,omitempty"`
	ModType string      `json:"mod_type,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
