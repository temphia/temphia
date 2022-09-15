package wazero

import (
	"context"
	"strings"

	"github.com/temphia/temphia/code/core/backend/xtypes/store"
)

func PlugKVSet(ctx context.Context, txid, keyPtr, keyLen, valPtr, valLen, optPtr, optLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	var opts *store.SetOptions
	if optLen != 0 {
		opts = &store.SetOptions{}
		err := e.getJSON((optPtr), (optLen), opts)
		if err != nil {
			e.writeError((respOffset), (respLen), err)
			return 0
		}
	}

	err := e.bindPluKV.Set(
		uint32(txid),
		e.getString((keyPtr), (keyLen)),
		e.getString((valPtr), (valLen)),
		opts,
	)

	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1

}

func PlugKVGet(ctx context.Context, txid, keyPtr, keyLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindPluKV.Get(uint32(txid), e.getString((keyPtr), (keyLen)))
	return e.writeJSONFinal((respOffset), (respLen), resp, err)
}

func PlugKVUpdate(ctx context.Context, txid, keyPtr, keyLen, valPtr, valLen, optPtr, optLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	var opts *store.UpdateOptions
	if optLen != 0 {
		opts = &store.UpdateOptions{}
		err := e.getJSON((optPtr), (optLen), opts)
		if err != nil {
			e.writeError((respOffset), (respLen), err)
			return 0
		}
	}

	err := e.bindPluKV.Update(
		uint32(txid),
		e.getString((keyPtr), (keyLen)),
		e.getString((valPtr), (valLen)),
		opts,
	)

	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1

}

func PlugKVDel(ctx context.Context, txid, keyPtr, keyLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	err := e.bindPluKV.Del(uint32(txid), e.getString((keyPtr), (keyLen)))
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1
}

func PlugKVDelBatch(ctx context.Context, txid, keyPtr, keyLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindPluKV.DelBatch(uint32(txid), strings.Split(e.getString((keyPtr), (keyLen)), ","))
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1
}

func PlugKVQuery(ctx context.Context, txid, optPtr, optLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	opts := &store.PkvQuery{}
	err := e.getJSON((optPtr), (optLen), opts)
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	resp, err := e.bindPluKV.Query(uint32(txid), opts)
	return e.writeJSONFinal((respOffset), (respLen), resp, err)
}

func PlugKVNewTxn(ctx context.Context, txidPtr, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	txid, err := e.bindPluKV.NewTxn()
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	ok := e.mem.WriteUint32Le(e.context, uint32(txidPtr), txid)
	if !ok {
		panic(ErrOutofIndex)
	}

	return 1
}

func PlugKVRollBack(ctx context.Context, txid, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	err := e.bindPluKV.RollBack(uint32(txid))
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1
}

func PlugKVCommit(ctx context.Context, txid, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	err := e.bindPluKV.Commit(uint32(txid))
	if err != nil {
		e.writeError((respOffset), (respLen), err)
		return 0
	}

	return 1
}
