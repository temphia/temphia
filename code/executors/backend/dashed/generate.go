package dashed

import (
	"encoding/json"
	"time"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/event"
	"github.com/temphia/temphia/code/executors/backend/dashed/dashmodels"

	"github.com/ztrue/tracerr"
)

type GenerateResponse struct {
	Name        string                       `json:"name,omitempty"`
	Epoch       int64                        `json:"epoch,omitempty"`
	Sections    []*dashmodels.Section        `json:"sections,omitempty"`
	ErrorPanels map[string]map[string]string `json:"err_panels,omitempty"`
	Data        map[string]interface{}       `json:"data,omitempty"`
}

func (g *GenerateResponse) AddErr(section string, panel string, err string) {
	sec, ok := g.ErrorPanels[section]
	if !ok {
		sec = make(map[string]string)
		g.ErrorPanels[section] = sec
	}
	sec[panel] = err
}

func (g *GenerateResponse) AddData(source string, data interface{}) {
	_, ok := g.Data[source]
	if ok {
		return
	}
	g.Data[source] = data
}

const ErrSourceNotFound = "source not found"

func (s *SimpleDash) generate(ev *event.Request) (*event.Response, error) {

	go func() {
		time.Sleep(time.Second * 5)
		s.bindings.Log("Simpledash Exec!!1")
	}()

	resp := &GenerateResponse{
		Name:        s.model.Name,
		Epoch:       0,
		Sections:    s.model.Sections,
		ErrorPanels: make(map[string]map[string]string),
		Data:        make(map[string]interface{}),
	}

	for _, section := range s.model.Sections {
		for _, panel := range section.Panels {
			source, ok := s.model.Sources[panel.Source]
			if !ok {
				resp.AddErr(section.Name, panel.Name, ErrSourceNotFound)
				continue
			}
			resp.AddData(panel.Source, source.StaticData)
		}
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &event.Response{
		Payload: (out),
	}, nil
}
