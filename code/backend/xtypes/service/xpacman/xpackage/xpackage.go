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
	Name        string            `yaml:"name,omitempty"`
	Slug        string            `yaml:"slug,omitempty"`
	Type        string            `yaml:"type,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Icon        string            `yaml:"icon,omitempty"`
	Screenshots []string          `yaml:"screenshots,omitempty"`
	Version     string            `yaml:"version,omitempty"`
	Tags        []string          `yaml:"tags,omitempty"`
	Files       map[string]string `yaml:"files,omitempty"`
	ExtraMeta   map[string]any    `yaml:"extra_meta,omitempty"`
	EnvFile     string            `yaml:"env_file,omitempty"`
}

type AppSchema struct {
	Name    string               `yaml:"name,omitempty"`
	Objects map[string]AppObject `yaml:"objects,omitempty"`
	Steps   []AppStep            `yaml:"steps,omitempty"`
}

type AppStep struct {
	Name     string `yaml:"name,omitempty"`
	ObjectId string `yaml:"object_id,omitempty"`
	Type     string `yaml:"type,omitempty"`
	File     string `yaml:"file,omitempty"`
	Data     any    `yaml:"data,omitempty"`
}

type AppObject struct {
	Name string `yaml:"name,omitempty"`
	Type string `yaml:"type,omitempty"`
}

func (a *AppStep) DataAs(target any) error {

	a.Data = convert(a.Data.(map[interface{}]interface{}))

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
