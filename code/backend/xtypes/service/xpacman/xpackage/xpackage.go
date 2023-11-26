package xpackage

import (
	"encoding/json"
	"fmt"
)

const (
	TypeBundle     = "bundle"
	TypeDataGroup  = "data_group"
	TypeDataSheet  = "data_sheet"
	TypePlug       = "plug"
	TypeResource   = "resource"
	TypeTargetApp  = "target_app"
	TypeTargetHook = "target_hook"
)

type Manifest struct {
	Name        string            `toml:"name,omitempty"`
	Slug        string            `toml:"slug,omitempty"`
	Type        string            `toml:"type,omitempty"`
	Description string            `toml:"description,omitempty"`
	Icon        string            `toml:"icon,omitempty"`
	Screenshots []string          `toml:"screenshots,omitempty"`
	Version     string            `toml:"version,omitempty"`
	Tags        []string          `toml:"tags,omitempty"`
	Files       map[string]string `toml:"files,omitempty"`
	ExtraMeta   map[string]any    `toml:"extra_meta,omitempty"`
	EnvFile     string            `toml:"env_file,omitempty"`

	Objects map[string]AppObject `toml:"objects,omitempty"`
	Steps   []AppStep            `toml:"steps,omitempty"`
}

type AppSchema struct {
	Name    string               `toml:"name,omitempty"`
	Objects map[string]AppObject `toml:"objects,omitempty"`
	Steps   []AppStep            `toml:"steps,omitempty"`
}

type AppStep struct {
	Name     string `toml:"name,omitempty"`
	ObjectId string `toml:"object_id,omitempty"`
	Type     string `toml:"type,omitempty"`
	File     string `toml:"file,omitempty"`
	Data     any    `toml:"data,omitempty"`
}

type AppObject struct {
	Name    string         `toml:"name,omitempty"`
	Type    string         `toml:"type,omitempty"`
	Options map[string]any `toml:"options,omitempty"`
}

func (a *AppStep) DataAs(target any) error {

	out, err := json.Marshal(a.Data)
	if err != nil {
		return err
	}

	return json.Unmarshal(out, target)
}

func convert(m map[interface{}]interface{}) map[string]interface{} {
	res := map[string]interface{}{}
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
