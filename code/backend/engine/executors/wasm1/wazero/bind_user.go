package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/xtypes/etypes/bindx"
)

func UserList(ctx context.Context, gPtr, gLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindUser.ListUser(e.getString(gPtr, gLen))
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func UserMessage(ctx context.Context, gPtr, gLen, uPtr, uLen, optPtr, optLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	opts := bindx.UserMessage{}

	err := e.getJSON(optPtr, optLen, &opts)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	err = e.bindUser.MessageUser(e.getString(gPtr, gLen), e.getString(uPtr, uLen), &opts)
	return e.writeFinal(respOffset, respLen, err)
}

func UserGet(ctx context.Context, gPtr, gLen, uPtr, uLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	usr, err := e.bindUser.GetUser(e.getString(gPtr, gLen), e.getString(uPtr, uLen))
	return e.writeJSONFinal(respOffset, respLen, usr, err)
}

func UserCurrentMessage(ctx context.Context, optPtr, optLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	opts := bindx.UserMessage{}

	err := e.getJSON(optPtr, optLen, &opts)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	err = e.bindUser.MessageCurrentUser(&opts)
	return e.writeFinal(respOffset, respLen, err)
}

func UserCurrent(ctx context.Context, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	usrs, err := e.bindUser.CurrentUser()
	return e.writeJSONFinal(respOffset, respLen, usrs, err)
}
