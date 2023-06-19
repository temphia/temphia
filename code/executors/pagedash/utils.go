package pagedash

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

type Index struct {
	ActiveKey    string            `json:"active_key,omitempty"`
	BuildDate    string            `json:"build_date,omitempty"`
	PreviousKeys map[string]string `json:"previous_keys,omitempty"`
	Version      int64             `json:"-"`
}

func (pd *PageDash) getBuildData(key string) (map[string]any, error) {
	// pkv := pd.binder.PlugKVBindingsGet()

	// resp, err := pkv.Get(0, key)
	// if err != nil {
	// 	return nil, err
	// }

	// data := make(map[string]any)

	// err = json.Unmarshal([]byte(resp.Value), &data)
	// if err != nil {
	// 	return nil, err
	// }

	// return data, nil

	return nil, nil
}

func (pd *PageDash) setBuildData(key string, data map[string]any) error {
	// pkv := pd.binder.PlugKVBindingsGet()

	// out, err := json.Marshal(&data)
	// if err != nil {
	// 	return err
	// }

	// return pkv.Set(0, key, string(out), &store.SetOptions{})
	return nil
}

func (pd *PageDash) delBuldData(key string) error {
	// pkv := pd.binder.PlugKVBindingsGet()
	// return pkv.Del(0, key)

	return nil
}

func (pd *PageDash) setIndex(index *Index) error {

	out, err := json.Marshal(index)
	if err != nil {
		return err
	}

	return pd.setIndexRaw(string(out), index.Version)
}

func (pd *PageDash) getIndex() (*Index, error) {
	val, err := pd.getIndexRaw()
	if err != nil {
		return nil, err
	}
	index := &Index{}

	err = json.Unmarshal([]byte(val.Value), index)
	if err != nil {
		return nil, err
	}

	index.Version = val.Version
	if index.PreviousKeys == nil {
		index.PreviousKeys = make(map[string]string)
	}

	return index, nil
}

func (pd *PageDash) getIndexRaw() (*entities.PlugKV, error) {
	// pkv := pd.binder.PlugKVBindingsGet()

	// val, err := pkv.Get(0, "index")
	// if err != nil {
	// 	return nil, err
	// }

	// return val, nil

	return nil, nil

}

func (pd *PageDash) setIndexRaw(value string, version int64) error {
	// pkv := pd.binder.PlugKVBindingsGet()
	// err := pkv.Update(0, "index", value, &store.UpdateOptions{
	// 	WithVerison: true,
	// 	Version:     int(version),
	// })

	// if err == nil {
	// 	return nil
	// }

	// if errors.Is(err, db.ErrNoMoreRows) {
	// 	err = pkv.Set(0, "index", value, &store.SetOptions{})
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }

	// return err

	return nil

}

func convertModel(model *DashModel) {

	for sidx := range model.Sections {
		section := &model.Sections[sidx]

		if section.Options != nil {
			convertInner(section.Options)
		}

		for pidx := range section.Panels {

			panel := &section.Panels[pidx]
			if panel.DataOpts != nil {
				convertInner(panel.DataOpts)
			}

		}

	}

	convertInner(model.StaticData)
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

		case []interface{}:
			for innIdx := range v2 {
				innVal := v2[innIdx]
				switch innv2 := innVal.(type) {
				case map[interface{}]interface{}:
					m2 := make(map[string]interface{})
					for k, v := range innv2 {
						m2[fmt.Sprint(k)] = v
					}
					convertInner(m2)

					v2[innIdx] = m2
				}

			}

		default:
			continue
		}
	}
}
