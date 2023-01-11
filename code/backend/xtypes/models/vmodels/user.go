package vmodels

type ProfileUpdate struct {
	Firstname  string `json:"firstname,omitempty" db:"firstname"`
	Middlename string `json:"middlename,omitempty" db:"middlename"`
	Lastname   string `json:"lastname,omitempty" db:"lastname"`
}
