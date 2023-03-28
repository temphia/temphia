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

func convertInner(m map[string]interface{}) {
	for idx, data := range m {
		switch v2 := data.(type) {
		case map[interface{}]interface{}:
			m2 := make(map[string]interface{})
			for k, v := range v2 {
				m2[fmt.Sprint(k)] = v
			}
			convertInner(m2)
			m[idx] = m2
		default:
			continue
		}
	}
}

/*

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


*/
