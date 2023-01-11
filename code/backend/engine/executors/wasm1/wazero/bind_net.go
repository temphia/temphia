package wazero

import (
	"context"
	"encoding/json"

	"github.com/temphia/temphia/code/backend/xtypes/etypes/bindx"
)

func HttpRaw(ctx context.Context, mPtr, mLen, pPtr, pLen, hPtr, hLen, bPtr, bLen,
	rStatus, rHeadPtr, rHeadLen, risJson, rbodyPtr, rbodyLen int32) int32 {

	e := getCtx(ctx)

	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(rbodyPtr, rbodyLen, err)
		return 0
	}

	req := bindx.HttpRequest{
		Method:  e.getString(mPtr, mLen),
		Path:    e.getString(pPtr, pLen),
		Headers: headers,
		Body:    e.getBytes(bPtr, bLen),
	}

	resp := e.bindNet.HttpRaw(&req)

	ok := e.mem.WriteUint32Le(e.context, uint32(rStatus), uint32(resp.SatusCode))
	if !ok {
		panic(ErrOutofIndex)
	}

	headBytes, err := json.Marshal(resp.Headers)
	if err != nil {
		e.writeError(rbodyPtr, rbodyLen, err)
		return 0
	}

	e.writeBytesNPtr(headBytes, rHeadPtr, rHeadLen)
	e.writeBytesNPtr(req.Body, rbodyPtr, rbodyLen)

	return 1

}

func HttpRawBatch(ctx context.Context, reqsPtr, reqsLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	reqs := make([]*bindx.HttpRequest, 0)
	err := e.getJSON(reqsPtr, reqsLen, &reqs)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	resp := e.bindNet.HttpRawBatch(reqs)
	return e.writeJSON(respOffset, respLen, resp)
}

func HttpQuickGet(ctx context.Context, uPtr, uLen, hPtr, hLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	url := e.getString(uPtr, uLen)
	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	out, err := e.bindNet.HttpQuickGet(url, headers)
	return e.writeBytesFinal(respOffset, respLen, out, err)
}

func HttpQuickPost(ctx context.Context, uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	url := e.getString(uPtr, uLen)
	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	bodyBytes := e.getBytes(dPtr, dLen)
	out, err := e.bindNet.HttpQuickPost(url, headers, bodyBytes)
	return e.writeBytesFinal(respOffset, respLen, out, err)
}

func HttpFormPost(ctx context.Context, uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	url := e.getString(uPtr, uLen)
	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	bodyBytes := e.getBytes(dPtr, dLen)
	out, err := e.bindNet.HttpFormPost(url, headers, bodyBytes)
	return e.writeBytesFinal(respOffset, respLen, out, err)
}

func HttpJsonGet(ctx context.Context, uPtr, uLen, hPtr, hLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	url := e.getString(uPtr, uLen)
	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	out, err := e.bindNet.HttpJsonGet(url, headers)
	return e.writeBytesFinal(respOffset, respLen, out, err)

}

func HttpJsonPost(ctx context.Context, uPtr, uLen, hPtr, hLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	url := e.getString(uPtr, uLen)
	headers, err := e.getStrMap(hPtr, hLen)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	bodyBytes := e.getBytes(dPtr, dLen)
	out, err := e.bindNet.HttpJsonPost(url, headers, bodyBytes)
	return e.writeBytesFinal(respOffset, respLen, out, err)
}
