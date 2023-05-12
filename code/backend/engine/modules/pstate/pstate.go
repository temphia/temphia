package pstate

import (
	"github.com/temphia/temphia/code/backend/engine/runtime/modipc"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type PlugStateMod struct {
	tenantId string
	plugId   string
	pkv      store.PlugStateKV

	modipc *modipc.ModIPC
}

func (ps *PlugStateMod) Handle(method string, args xtypes.LazyData) (xtypes.LazyData, error) {
	return ps.modipc.Handle(method, args)
}

func (ps *PlugStateMod) Close() error {
	ps.modipc = nil
	ps.pkv = nil

	return nil
}

type setOpts struct {
	Key   string            `json:"key,omitempty"`
	Value string            `json:"value,omitempty"`
	Opts  *store.SetOptions `json:"opts,omitempty"`
}

func (ps *PlugStateMod) Set(opts *setOpts) error {
	return ps.pkv.Set(0, ps.tenantId, ps.plugId, opts.Key, opts.Value, opts.Opts)
}

type updateOpts struct {
	Key   string               `json:"key,omitempty"`
	Value string               `json:"value,omitempty"`
	Opts  *store.UpdateOptions `json:"opts,omitempty"`
}

func (ps *PlugStateMod) Update(opts *updateOpts) error {
	return ps.pkv.Update(0, ps.tenantId, ps.plugId, opts.Key, opts.Value, opts.Opts)
}

func (ps *PlugStateMod) Get(key string) (*entities.PlugKV, error) {
	return ps.pkv.Get(0, ps.tenantId, ps.plugId, key)
}

func (ps *PlugStateMod) Query(query *store.PkvQuery) ([]*entities.PlugKV, error) {
	return ps.pkv.Query(0, ps.tenantId, ps.plugId, query)
}

func (ps *PlugStateMod) Del(key string) error {
	return ps.pkv.Del(0, ps.tenantId, ps.plugId, key)
}

func (ps *PlugStateMod) DelBatch(keys []string) error {
	return ps.pkv.DelBatch(0, ps.tenantId, ps.plugId, keys)
}
