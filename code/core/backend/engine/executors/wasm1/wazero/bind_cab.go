package wazero

import (
	"context"
)

func cabAddFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, dataPtr, dataLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	contents := e.getBytes(dataPtr, dataLen)
	err := e.bindCab.AddFile(
		e.getString(bukPtr, bukLen),
		e.getString(filePtr, fileLen),
		contents,
	)

	return e.writeFinal(respOffset, respLen, err)
}

func cabListFolder(ctx context.Context, bukPtr, bukLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	resp, err := e.bindCab.ListFolder(e.getString(bukPtr, bukLen))
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func cabGetFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	out, err := e.bindCab.GetFile(
		e.getString(bukPtr, bukLen),
		e.getString(filePtr, fileLen),
	)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(out, (respOffset), (respLen))
	return 1
}

func cabDelFile(ctx context.Context, bukPtr, bukLen, filePtr, fileLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindCab.DeleteFile(
		e.getString(bukPtr, bukLen),
		e.getString(filePtr, fileLen),
	)

	return e.writeFinal(respOffset, respLen, err)
}
