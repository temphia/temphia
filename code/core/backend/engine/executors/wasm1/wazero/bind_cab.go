package wazero

import (
	"context"

	"github.com/temphia/temphia/code/core/backend/libx/xutils/kosher"
)

func CabinetAddFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, dataPtr, dataLen, respPtr, respLen int32) int32 {

	e := getCtx(ctx)

	contents, ok := e.instance.Memory().Read(e.context, uint32(dataPtr), uint32(dataLen))
	if !ok {
		panic(ErrOutofIndex)
	}

	err := e.bindCab.AddFile(
		e.getString(uint32(bukPtr), uint32(bukLen)),
		e.getString(uint32(filePtr), uint32(fileLen)),
		contents,
	)

	if err != nil {
		e.write2(kosher.Byte(err.Error()), uint32(respPtr), uint32(respLen))
		return 0
	}

	return 1
}

func ListFolder(ctx context.Context, bukPtr, bukLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	resp, err := e.bindCab.ListFolder(e.getString(uint32(bukPtr), uint32(bukLen)))
	return e.writeJSON(uint32(respPtr), uint32(respLen), resp, err)
}

func GetFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	out, err := e.bindCab.GetFile(
		e.getString(uint32(bukPtr), uint32(bukLen)),
		e.getString(uint32(filePtr), uint32(fileLen)),
	)
	if err != nil {
		e.write2(kosher.Byte(err.Error()), uint32(respPtr), uint32(respLen))
		return 0
	}

	e.write2(out, uint32(respPtr), uint32(respLen))
	return 1
}

func DelFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respPtr, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindCab.DeleteFile(
		e.getString(uint32(bukPtr), uint32(bukLen)),
		e.getString(uint32(filePtr), uint32(fileLen)),
	)
	if err != nil {
		e.write2(kosher.Byte(err.Error()), uint32(respPtr), uint32(respLen))
		return 0
	}

	return 1
}
