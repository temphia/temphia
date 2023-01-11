package wazero

import (
	"context"

	"github.com/temphia/temphia/code/backend/libx/lazydata"
)

func SelfListResources(ctx context.Context, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfListResources()
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func SelfGetResource(ctx context.Context, nPtr, nLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfGetResource(e.getString(nPtr, nLen))
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func SelfInLinks(ctx context.Context, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfInLinks()
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func SelfOutLinks(ctx context.Context, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfOutLinks()
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func SelfLinkExec(ctx context.Context, nPtr, nLen, mPtr, mLen, dPtr, dLen, async, detached, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	out, err := e.bindSelf.SelfLinkExec(
		e.getString(nPtr, nLen),
		e.getString(mPtr, mLen),
		lazydata.NewJsonData(e.getBytes(dPtr, dLen)),
		async == 1,
		detached == 1,
	)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	bytes, err := out.AsJsonBytes()
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(bytes, respOffset, respLen)
	return 1
}

func SelfModuleExec(ctx context.Context, nPtr, nLen, mPtr, mLen, pPtr, pLen, dPtr, dLen, respOffset, respLen int32) int32 {

	e := getCtx(ctx)

	out, err := e.bindSelf.SelfModuleExec(
		e.getString(nPtr, nLen),
		e.getString(mPtr, mLen),
		e.getString(pPtr, pLen),
		lazydata.NewJsonData(e.getBytes(dPtr, dLen)),
	)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	bytes, err := out.AsJsonBytes()
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(bytes, respOffset, respLen)
	return 1
}

func SelfForkExec(ctx context.Context, mPtr, mLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfForkExec(
		e.getString(mPtr, mLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(respOffset, respLen, err)
}

func SelfAddFile(ctx context.Context, fPtr, fLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfAddFile(
		e.getString(fPtr, fLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(respOffset, respLen, err)
}

func SelfUpdateFile(ctx context.Context, fPtr, fLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfUpdateFile(
		e.getString(fPtr, fLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(respOffset, respLen, err)

}

func SelfAddDataFile(ctx context.Context, fPtr, fLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfAddDataFile(
		e.getString(fPtr, fLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(respOffset, respLen, err)

}

func SelfUpdateDataFile(ctx context.Context, fPtr, fLen, dPtr, dLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)

	err := e.bindSelf.SelfUpdateDataFile(
		e.getString(fPtr, fLen),
		e.getBytes(dPtr, dLen),
	)

	return e.writeFinal(respOffset, respLen, err)
}

func SelfGetDataFile(ctx context.Context, fptr, flen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	fout, err := e.bindSelf.SelfGetDataFile(e.getString(fptr, flen))
	return e.writeBytesFinal(respOffset, respLen, fout, err)
}

func SelfListDataFiles(ctx context.Context, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	resp, err := e.bindSelf.SelfListDataFiles()
	return e.writeJSONFinal(respOffset, respLen, resp, err)
}

func SelfDelDataFile(ctx context.Context, fptr, flen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	err := e.bindSelf.SelfDeleteDataFile(e.getString(fptr, flen))
	return e.writeFinal(respOffset, respLen, err)
}
