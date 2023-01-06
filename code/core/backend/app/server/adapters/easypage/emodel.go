package easypage

type Page struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

type NewPage struct {
	Slug string `json:"slug,omitempty"`
	Data string `json:"data,omitempty"`
}
