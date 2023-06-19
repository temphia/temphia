package binder

import (
	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/thoas/go-funk"
)

type PkvBindings struct {
	stateKv   store.PlugStateKV
	namespace string
	plugId    string
	agentid   string
	txns      []uint32
}

func NewPKV(pkv store.PlugStateKV, tenantid, plugId, agentid string) PkvBindings {
	return PkvBindings{
		stateKv:   pkv,
		namespace: tenantid,
		plugId:    plugId,
		agentid:   agentid,
		txns:      make([]uint32, 0, 1),
	}
}

func (pkv *PkvBindings) checkTxn(tx uint32) error {
	if tx == 0 {
		return nil
	}
	if !funk.ContainsUInt32(pkv.txns, tx) {
		return easyerr.NotFound("txn not found")
	}
	return nil
}

func (pkv *PkvBindings) NewTxn() (uint32, error) {
	tx, err := pkv.stateKv.NewTxn()
	if err != nil {
		return 0, err
	}
	pkv.txns = append(pkv.txns, tx)
	return tx, nil
}

func (pkv *PkvBindings) RollBack(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.RollBack(txid)
}

func (pkv *PkvBindings) Commit(txid uint32) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}

	return pkv.stateKv.Commit(txid)
}

func (pkv *PkvBindings) Set(txid uint32, key, value string, opts *store.SetOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Set(txid, pkv.namespace, pkv.plugId, key, value, opts)
}

func (pkv *PkvBindings) Update(txid uint32, key, value string, opts *store.UpdateOptions) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Update(txid, pkv.namespace, pkv.plugId, key, value, opts)
}

func (pkv *PkvBindings) Get(txid uint32, key string) (*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}
	return pkv.stateKv.Get(txid, pkv.namespace, pkv.plugId, key)
}

func (pkv *PkvBindings) Del(txid uint32, key string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.Del(txid, pkv.namespace, pkv.plugId, key)
}

func (pkv *PkvBindings) DelBatch(txid uint32, keys []string) error {
	err := pkv.checkTxn(txid)
	if err != nil {
		return err
	}
	return pkv.stateKv.DelBatch(txid, pkv.namespace, pkv.plugId, keys)
}

func (pkv *PkvBindings) Query(txid uint32, query *store.PkvQuery) ([]*entities.PlugKV, error) {
	err := pkv.checkTxn(txid)
	if err != nil {
		return nil, err
	}

	return pkv.stateKv.Query(txid, pkv.namespace, pkv.plugId, query)
}
