package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
)

func cabAddFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, dataPtr, dataLen, respPtr, respLen int32) int32 {

	e := getCtx(ctx)

	contents, ok := e.instance.Memory().Read(e.context, uint32(dataPtr), uint32(dataLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	err := e.bindCab.AddFile(
		e.getString((bukPtr), (bukLen)),
		e.getString((filePtr), (fileLen)),
		contents,
	)

	if err != nil {
		e.writeError((respPtr), (respLen), err)
		return 0
	}

	return 1
}

func cabListFolder(ctx context.Context, bukPtr, bukLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	resp, err := e.bindCab.ListFolder(e.getString((bukPtr), (bukLen)))
	return e.writeJSONFinal((respPtr), (respLen), resp, err)
}

func cabGetFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	out, err := e.bindCab.GetFile(
		e.getString((bukPtr), (bukLen)),
		e.getString((filePtr), (fileLen)),
	)
	if err != nil {
		e.writeBytesNPtr(kosher.Byte(err.Error()), (respPtr), (respLen))
		return 0
	}

	e.writeBytesNPtr(out, (respPtr), (respLen))
	return 1
}

func cabDelFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindCab.DeleteFile(
		e.getString((bukPtr), (bukLen)),
		e.getString((filePtr), (fileLen)),
	)
	if err != nil {
		e.writeBytesNPtr(kosher.Byte(err.Error()), (respPtr), (respLen))
		return 0
	}

	return 1
}
