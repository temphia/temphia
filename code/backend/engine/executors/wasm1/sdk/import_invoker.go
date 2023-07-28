package tasmsdk

import (
	"errors"

	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func InvokerName() string {
	var respOffset, respLen int32
	_invoker_name(intAddr(&respOffset), intAddr(&respLen))

	return kosher.Str(getBytes(respOffset))
}

func InvokerExec(method string, data []byte) ([]byte, error) {

	var respOffset, respLen int32

	mptr, mlen := stringToPtr(method)

	dptr, dlen := bytesToPtr(data)

	if _invoker_exec(0, mptr, mlen, dptr, dlen, respOffset, respLen) {
		return nil, getErr(respOffset)
	}

	return getBytes(respOffset), nil
}

func InvokerContextUser() (*claim.UserContext, error) {
	var respOffset, respLen int32

	if _invoker_ctx_user(0, respOffset, respLen) {
		ctx := &claim.UserContext{}
		err := getJSON(respOffset, ctx)
		if err != nil {
			return ctx, err
		}

	}

	return nil, nil
}

func InvokerContextInfo() (*entities.UserInfo, error) {
	var respOffset, respLen int32

	if _invoker_ctx_user_info(0, respOffset, respLen) {
		info := &entities.UserInfo{}
		err := getJSON(respOffset, info)
		if err != nil {
			return nil, err
		}

		return info, nil
	}

	return nil, getErr(respOffset)
}

func InvokerContextUserMessage(opts *bindx.UserMessage) error {

	optPtr, optLen := JsonPtr(opts)
	var respOffset, respLen int32

	ok := _invoker_ctx_message(0, int32(uintptr(optPtr)), optLen, intAddr(&respOffset), intAddr(&respLen))

	resp := getBytes(respOffset)
	if !ok {
		return errors.New(string(resp))
	}

	return nil
}

//go:wasm-module temphia1
//export invoker_name
func _invoker_name(respOffset, respLen int32) bool

//go:wasm-module temphia1
//export invoker_exec
func _invoker_exec(ctxid, mPtr, mLen, dptr, dlen int32, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export invoker_ctx_user
func _invoker_ctx_user(ctxid, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export invoker_ctx_user_info
func _invoker_ctx_user_info(ctxid, respOffset, respLen int32) bool

//go:wasm-module temphia1
//export invoker_ctx_message
func _invoker_ctx_message(ctxid, dptr, dlen, respOffset, respLen int32) bool
