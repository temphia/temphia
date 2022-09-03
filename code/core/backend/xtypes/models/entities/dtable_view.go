package entities

import (
	"database/sql/driver"
	"encoding/json"
)

type DataView struct {
	Id          int64       `json:"id,omitempty" db:"id,omitempty"`
	Name        string      `json:"name,omitempty" db:"name,omitempty"`
	Count       int64       `json:"count,omitempty" db:"count"`
	FilterConds FilterConds `json:"filter_conds,omitempty" db:"filter_conds"` // fixme => ?
	Selects     []string    `json:"selects,omitempty" db:"selects"`
	MainColumn  string      `json:"main_column,omitempty" db:"main_column"`
	SearchTerm  string      `json:"search_term,omitempty" db:"search_term"`
	TableID     string      `json:"table_id,omitempty" db:"table_id"`
	GroupID     string      `json:"group_id,omitempty" db:"group_id"`
	TenantID    string      `json:"tenant_id,omitempty" db:"tenant_id"`
	ExtraMeta   JsonStrMap  `json:"extra_meta,omitempty" db:"extra_meta,omitempty"`
}

type View struct {
	Name        string      `json:"name,omitempty"`
	Count       int64       `json:"count,omitempty"`
	FilterConds FilterConds `json:"filter_conds,omitempty"`
	Selects     []string    `json:"selects,omitempty"`
	MainColumn  string      `json:"main_column,omitempty"`
	SearchTerm  string      `json:"search_term,omitempty"`
}

type FilterConds []interface{}

func (j *FilterConds) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	out, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(out), nil
}

func (j *FilterConds) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch s := value.(type) {
	case string:
		if s == "" {
			return nil
		}
		return json.Unmarshal([]byte(s), &j)
	case []byte:
		return json.Unmarshal(s, j)
	}
	return nil
}
