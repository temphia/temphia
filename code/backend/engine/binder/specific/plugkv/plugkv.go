package plugkv

import (
	"github.com/temphia/temphia/code/backend/engine/binder/handle"
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/etypes"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx/ticket"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/service"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type Binding struct {
	stateKv   store.PlugStateKV
	signer    service.Signer
	namespace string
	plugId    string
	agentid   string
	txns      []uint32
	handle    *handle.Handle
}

func New(handle *handle.Handle) Binding {
	return Binding{
		signer:    handle.Deps.Signer,
		stateKv:   handle.Deps.PlugKV,
		namespace: handle.Namespace,
		plugId:    handle.PlugId,
		agentid:   handle.AgentId,
		txns:      make([]uint32, 0, 1),
		handle:    handle,
	}
}

func (pkv *Binding) checkTxn(tx uint32) error {
	if tx == 0 {
		return nil
	}
	if !funk.ContainsUInt32(pkv.txns, tx) {
		return easyerr.NotFound()
	}
	return nil
}

func (pkv *Binding) NewTxn() (uint32, error) {
	tx, err := pkv.stateKv.NewTxn()
	if err != nil {
		return 0, err
	}
	pkv.txns = append(pkv.txns, tx)
	return tx, nil
}

func (pkv *Binding) RollBack(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.RollBack(txid)
}

func (pkv *Binding) Commit(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.Commit(txid)
}

func (pkv *Binding) Set(txid uint32, key, value string, opts *store.SetOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Set(txid, pkv.namespace, pkv.plugId, key, value, opts)
}

func (pkv *Binding) Update(txid uint32, key, value string, opts *store.UpdateOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Update(txid, pkv.namespace, pkv.plugId, key, value, opts)
}

func (pkv *Binding) Get(txid uint32, key string) (*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}
	return pkv.stateKv.Get(txid, pkv.namespace, pkv.plugId, key)
}

func (pkv *Binding) Del(txid uint32, key string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Del(txid, pkv.namespace, pkv.plugId, key)
}

func (pkv *Binding) DelBatch(txid uint32, keys []string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.DelBatch(txid, pkv.namespace, pkv.plugId, keys)
}

func (pkv *Binding) Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}

	return pkv.stateKv.Query(txid, pkv.namespace, pkv.plugId, query)
}

func (pkv *Binding) Ticket(opts *ticket.PlugState) (string, error) {

	uctx := pkv.handle.Job.Invoker.UserContext()
	if uctx == nil {
		return "", easyerr.Error(etypes.EmptyUserContext)
	}

	return pkv.signer.SignPlugState(pkv.namespace, &claim.PlugState{
		TenantId:  pkv.namespace,
		Type:      "",
		UserId:    uctx.Id,
		DeviceId:  uctx.DeviceId,
		SessionId: uctx.SessionId,
		ExecId:    0,
		PlugId:    pkv.plugId,
		AgentId:   pkv.agentid,
		KeyPrefix: opts.KeyPrefix,
	})

}