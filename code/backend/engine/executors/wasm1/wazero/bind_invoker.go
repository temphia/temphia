package wazero

import (
	"context"

	"github.com/temphia/temphia/code/backend/libx/lazydata"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

func InvokerName(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	e.writeBytesNPtr(kosher.Byte(e.invoker.Name()), ctxid, respOffset, respLen)
	return 1
}

func InvokerExec(ctx context.Context, ctxid, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	method := e.getString(mPtr, mLen)
	data := lazydata.NewJsonData(e.getBytes(dPtr, dLen))

	rdata, err := e.invoker.ExecMethod(method, data)
	if err != nil {
		return 0
	}

	bytes, err := rdata.AsJsonBytes()
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(bytes, ctxid, respOffset, respLen)

	return 1

}

func InvokerContextUser(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	user := e.invoker.UserContext()
	if user == nil {
		return 0
	}

	return e.writeJSONFinal(ctxid, respOffset, respLen, user, nil)
}

func InvokerContextInfo(ctx context.Context, ctxid, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	user, err := e.invoker.UserInfo()
	return e.writeJSONFinal(ctxid, respOffset, respLen, user, err)
}

func InvokerContextUserMessage(ctx context.Context, ctxid, dPtr, dLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	msg := &bindx.UserMessage{}

	err := e.getJSON(dPtr, dLen, msg)
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	err = e.invoker.UserMessage(msg)
	if err != nil {
		e.writeError(ctxid, respOffset, respLen, err)
		return 0
	}

	return 1
}
