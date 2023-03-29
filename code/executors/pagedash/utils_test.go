package pagedash

import (
	"testing"
)

func TestConvert(t *testing.T) {
	model := &DashModel{
		Name: "",
		Sections: []Section{
			{
				Name:   "section 1",
				Layout: "",
				Options: map[string]any{
					"test1": 1,
					"test2": map[any]any{
						"nested1": map[any]any{
							"nn1": true,
						},
					},
				},
			},
		},
	}

	if model.Sections[0].Options["test2"].(map[any]any)["nested1"].(map[any]any)["nn1"] != true {

		t.Fatal("failed to convert")

	}

}
