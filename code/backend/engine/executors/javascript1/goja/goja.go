package goja

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/event"

	"github.com/temphia/temphia/code/backend/xtypes/etypes"
)

var _ etypes.Executor = (*Goja)(nil)

type Goja struct {
	runtime *goja.Runtime
	binder  bindx.Bindings
}

func NewBuilder(app any) (etypes.ExecutorBuilder, error) {
	return etypes.ExecBuilderFunc(newExecutor), nil

}

func newExecutor(opts etypes.ExecutorOption) (etypes.Executor, error) {

	script, _, err := opts.Binder.GetFileWithMeta(opts.File)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	// fixme => snowflake

	rt.SetRandSource(func() float64 {
		fid := rand.Float64()
		return fid
	})

	_, err = rt.RunProgram(libesplug)
	if err != nil {
		panic(err)
	}

	_, err = rt.RunString(string(script))
	if err != nil {
		return nil, err
	}

	return New(opts.Binder, rt)
}

func New(b bindx.Bindings, rt *goja.Runtime) (*Goja, error) {

	g := &Goja{
		runtime: rt,
		binder:  b,
	}

	g.bind()

	_, err := rt.RunProgram(libesplug)
	if err != nil {
		return nil, err
	}

	return g, nil
}

type Response struct {
	Payload any `json:"payload,omitempty"`
}

type Request struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (g *Goja) Process(ev *event.Request) (*event.Response, error) {
	var entry func(ev *Request) (*Response, error)
	rawentry := g.runtime.Get(fmt.Sprintf("action_%s", ev.Name))
	if rawentry == nil {
		return nil, errors.New("method not found")
	}

	err := g.runtime.ExportTo(rawentry, &entry)
	if err != nil {
		return nil, err
	}

	resp, err := entry(&Request{
		Id:   ev.Id,
		Name: ev.Name,
		Data: ev.Data,
	})

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(&resp.Payload)
	if err != nil {
		return nil, err
	}

	return &event.Response{
		Payload: (out),
	}, nil

}
