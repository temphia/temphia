package sloader

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/executors/backend/wizard/wmodels"
)

const (
	LOADER_SOURCE_STATIC      = "static"
	LOADER_SOURCE_JSSCRIPT    = "js_script"
	LOADER_DTABLE_SIMPLEQUERY = "dtable_simple_query"
	LOADER_DGROUP_HSQL_QUERY  = "dgroup_hsql_query"
	LOADER_DGROUP_RAW_QUERY   = "dgroup_raw_query"
)

type SLoader struct {
	Binding     bindx.Bindings
	Model       *wmodels.Wizard
	SubData     *wmodels.Submission
	Stage       *wmodels.Stage
	Group       *wmodels.StageGroup
	DataSources map[string]interface{}
}

func (s *SLoader) Process() error {

	for _, field := range s.Stage.Fields {
		if field.Source == "" {
			continue
		}

		_, ok := s.DataSources[field.Source]
		if ok {
			continue
		}

		source := s.Model.Sources[field.Source]
		if source == nil {
			return easyerr.NotFound()
		}

		value, err := s.process(field, source)
		if err != nil {
			return err
		}
		s.DataSources[field.Source] = value
	}

	return nil
}

func (s *SLoader) process(field *wmodels.Field, source *wmodels.Source) (interface{}, error) {
	switch source.Type {
	case LOADER_SOURCE_STATIC:
		return source.Data, nil
	case LOADER_SOURCE_JSSCRIPT:
		return s.jsScript(field, source)
	default:
		return nil, easyerr.Error((fmt.Sprint("Skipping field, source not implemented", field)))
	}
}

func (s *SLoader) jsScript(field *wmodels.Field, source *wmodels.Source) (interface{}, error) {
	pp.Println("Executing =>", source.Target, s.bindings())
	return nil, nil
}

func (s *SLoader) bindings() map[string]interface{} {
	return map[string]interface{}{
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},
	}
}
