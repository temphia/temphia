package wazero

import "context"

func log(ctx context.Context, offset, len int32) {
	e := getCtx(ctx)
	out, ok := e.getMem().Read(ctx, uint32(offset), uint32(len))
	if !ok {
		panic(ErrOutofIndex)
	}

	e.bindings.Log(string(out))
}

func sleep(ctx context.Context, msec int32) {
	e := getCtx(ctx)
	e.bindings.Sleep(msec)
}

func getFileWithMeta(ctx context.Context, filePtr, fileLen, respPtr, respLen, mod int32) int32 {
	e := getCtx(ctx)
	out, ok := e.getMem().Read(ctx, uint32(filePtr), uint32(fileLen))
	if ok {
		panic(ErrOutofIndex)
	}

	fout, _, err := e.bindings.GetFileWithMeta(string(out))
	if err != nil {
		e.write([]byte(err.Error()))
		return 0
	}

	ok = e.write2(fout, uint32(respPtr), uint32(respPtr))
	if !ok {
		panic(ErrOutofMemory)
	}

	return 1
}
