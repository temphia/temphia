package wazero

import (
	"context"

	"github.com/temphia/temphia/code/backend/libx/lazydata"
)

func SelfListResources(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfListResources()
	return e.writeJSONFinal(ctxid, respOffset, respLen, resp, err)
}

func SelfGetResource(ctx context.Context, ctxid, nPtr, nLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfGetResource(e.getString(nPtr, nLen))
	return e.writeJSONFinal(ctxid, respOffset, respLen, resp, err)
}

func SelfInLinks(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfInLinks()
	return e.writeJSONFinal(ctxid, respOffset, respLen, resp, err)
}

func SelfOutLinks(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfOutLinks()
	return e.writeJSONFinal(ctxid, respOffset, respLen, resp, err)
}

func SelfLinkExec(ctx context.Context, ctxid, nPtr, nLen, mPtr, mLen, dPtr, dLen, async, detached, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	out, err := e.bindSelf.SelfLinkExec(
		e.getString(nPtr, nLen),
		e.getString(mPtr, mLen),
		lazydata.NewJsonData(e.getBytes(dPtr, dLen)),
		async == 1,
		detached == 1,
	)
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	bytes, err := out.AsJsonBytes()
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(bytes, ctxid, respOffset, respLen)
	return 1
}

func SelfNewModule(ctx context.Context, ctxid, nPtr, nLen, dPtr, dLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	out, err := e.bindSelf.SelfNewModule(
		e.getString(nPtr, nLen),
		lazydata.NewJsonData(e.getBytes(dPtr, dLen)),
	)
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	return out
}

func SelfModuleExec(ctx context.Context, ctxid, mid int32, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	out, err := e.bindSelf.SelfModuleExec(
		mid,
		e.getString(mPtr, mLen),
		lazydata.NewJsonData(e.getBytes(dPtr, dLen)),
	)

	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	bytes, err := out.AsJsonBytes()
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(bytes, ctxid, respOffset, respLen)

	return 1
}

func SelfForkExec(ctx context.Context, ctxid, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfForkExec(
		e.getString(mPtr, mLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(ctxid, respOffset, respLen, err)
}
