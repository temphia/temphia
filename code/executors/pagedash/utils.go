package pagedash

import (
	"fmt"
)

func convertModel(model *DashModel) {

	for sidx := range model.Sections {
		section := &model.Sections[sidx]

		if section.Options != nil {
			convertInner(section.Options)
		}

		for pidx := range section.Panels {

			panel := &section.Panels[pidx]
			if panel.Options != nil {
				convertInner(panel.Options)
			}

		}

	}

}

func convertInner(m map[string]any) {

	for idx, data := range m {

		switch v2 := data.(type) {
		case map[any]any:
			m[idx] = convert(v2)
		default:
			continue
		}
	}

}

func convert(m map[any]any) map[string]any {
	res := map[string]any{}
	for k, v := range m {
		switch v2 := v.(type) {
		case map[interface{}]interface{}:
			res[fmt.Sprint(k)] = convert(v2)
		default:
			res[fmt.Sprint(k)] = v
		}
	}
	return res
}
