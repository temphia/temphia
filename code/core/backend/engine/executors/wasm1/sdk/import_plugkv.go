package tasmsdk

import (
	"encoding/json"
	"errors"
)

func PlugKvSet(txid int32, key, value string, opts map[string]any) error {

	keyPtr, keyLen := stringToPtr(key)
	valuePtr, valueLen := stringToPtr(value)
	optPtr, optLen := JsonPtr(opts)

	var respOffset, respLen int32

	ok := _plugkv_set(txid, keyPtr, keyLen, valuePtr, valueLen, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func PlugKvUpdate(txid int32, key, value string, opts map[string]any) error {

	keyPtr, keyLen := stringToPtr(key)
	valuePtr, valueLen := stringToPtr(value)
	optPtr, optLen := JsonPtr(opts)

	var respOffset, respLen int32

	ok := _plugkv_update(txid, keyPtr, keyLen, valuePtr, valueLen, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func PlugKvGet(txid int32, key string) (map[string]any, error) {
	keyPtr, keyLen := stringToPtr(key)
	var respOffset, respLen int32

	ok := _plugkv_get(txid, keyPtr, keyLen, intAddr(&respOffset), intAddr(&respLen))
	resp := getBytes(respOffset)
	if !ok {
		return nil, errors.New(string(resp))
	}

	data := make(map[string]any)
	err := json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PlugKvDelete(txid int32, key string) error {
	keyPtr, keyLen := stringToPtr(key)
	var respOffset, respLen int32

	ok := _plugkv_delete(txid, keyPtr, keyLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)
}

func PlugKvDeleteBatch(txid int32, keys []string) error {

	keysJson, err := json.Marshal(&keys)
	if err != nil {
		return err
	}

	keyPtr, keyLen := bytesToPtr(keysJson)
	var respOffset, respLen int32

	ok := _plugkv_delete_batch(txid, keyPtr, keyLen, intAddr(&respOffset), intAddr(&respLen))
	if ok {
		return nil
	}

	return getErr(respOffset)

}

func PlugKvQuery(txid int32, opt map[string]any) ([]map[string]any, error) {

	optPtr, optLen := JsonPtr(opt)
	var respOffset, respLen int32

	ok := _plugkv_query(txid, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))

	resp := getBytes(respOffset)

	if !ok {
		return nil, errors.New(string(resp))
	}

	data := make([]map[string]any, 0)

	err := json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func PlugKvTxnNew() (int32, error) {
	var txid, respOffset, respLen int32

	ok := _plugkv_txn_new(intAddr(&txid), intAddr(&respOffset), intAddr(&respLen))
	if !ok {
		resp := getBytes(respOffset)
		return 0, errors.New(string(resp))
	}

	return txid, nil
}

func PlugKvTxnRollback(txid int32) error {
	var respOffset, respLen int32

	ok := _plugkv_txn_rollback(txid, intAddr(&respOffset), intAddr(&respLen))
	if !ok {
		resp := getBytes(respOffset)
		return errors.New(string(resp))
	}

	return nil
}

func PlugKvTxnCommit(txid int32) error {
	var respOffset, respLen int32

	ok := _plugkv_txn_commit(txid, intAddr(&respOffset), intAddr(&respLen))
	if !ok {
		resp := getBytes(respOffset)
		return errors.New(string(resp))
	}

	return nil
}

// private

//go:wasm-module temphia1
//export plugkv_set
func _plugkv_set(txid, keyPtr, keyLen, valuePtr, valueLen, optsPtr, optsLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_update
func _plugkv_update(txid, keyPtr, keyLen, valuePtr, valueLen, optsPtr, optsLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_get
func _plugkv_get(txid, keyPtr, keyLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_del
func _plugkv_delete(txid, keyPtr, keyLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_del_batch
func _plugkv_delete_batch(txid, keyPtr, keyLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_query
func _plugkv_query(txid, optsPtr, optsLen, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_txn_new
func _plugkv_txn_new(txidPtr, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_txn_rollback
func _plugkv_txn_rollback(txid, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export plugkv_txn_commit
func _plugkv_txn_commit(txid, respOffset, respLen int32) bool
