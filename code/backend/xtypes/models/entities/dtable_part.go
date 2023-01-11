package entities

type TableGroupPartial struct {
	Name          string  `json:"name,omitempty" db:"name,omitempty"`
	Description   string  `json:"description,omitempty" db:"description,omitempty"`
	CabinetSource string  `json:"cabinet_source,omitempty" db:"cabinet_source,omitempty"`
	CabinetFolder string  `json:"cabinet_folder,omitempty" db:"cabinet_folder,omitempty"`
	ExtraMeta     JsonMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type TablePartial struct {
	Name        string  `json:"name,omitempty" db:"name,omitempty"`
	Description string  `json:"description,omitempty" db:"description,omitempty"`
	Icon        string  `json:"icon,omitempty" db:"icon,omitempty"`
	MainColumn  string  `json:"main_column,omitempty" db:"main_column,omitempty"`
	ExtraMeta   JsonMap `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type ColumnPartial struct {
	Name          string   `json:"name,omitempty" db:"name,omitempty"`
	Description   string   `json:"description,omitempty" db:"description,omitempty"`
	Icon          string   `json:"icon,omitempty" db:"icon,omitempty"`
	Ctype         string   `json:"ctype,omitempty" db:"ctype,omitempty"`
	Options       []string `json:"options" db:"options,omitempty"`
	OrderID       int64    `json:"order_id,omitempty" db:"order_id,omitempty"`
	Pattern       string   `json:"pattern,omitempty" db:"pattern,omitempty"`
	RefId         string   `json:"ref_id,omitempty" db:"ref_id,omitempty"`
	RefType       string   `json:"ref_type,omitempty" db:"ref_type,omitempty"`
	RefTarget     string   `json:"ref_target,omitempty" db:"ref_target,omitempty"`
	RefObject     string   `json:"ref_object,omitempty" db:"ref_object,omitempty"`
	StrictPattern bool     `json:"strict_pattern,omitempty" db:"strict_pattern,omitempty"`
	ExtraMeta     JsonMap  `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}
