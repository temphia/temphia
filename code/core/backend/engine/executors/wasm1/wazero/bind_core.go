package wazero

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func BuildTemphiaModule(runtime wazero.Runtime) (api.Module, error) {
	return runtime.NewModuleBuilder("temphia").
		ExportFunction("log", log).
		ExportFunction("sleep", sleep).
		ExportFunction("get_file_with_meta", getFileWithMeta).
		Instantiate(context.TODO(), runtime)
}

// core binds

func log(ctx context.Context, offset, len int32) {
	e := getCtx(ctx)
	out, ok := e.mem.Read(ctx, uint32(offset), uint32(len))
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
	out, ok := e.mem.Read(ctx, uint32(filePtr), uint32(fileLen))
	if ok {
		panic(ErrOutofIndex)
	}

	fout, _, err := e.bindings.GetFileWithMeta(string(out))
	if err != nil {
		e.writeError(uint32(respPtr), uint32(respLen), err)
		return 0
	}

	ok = e.writeWithOffsetPtr(fout, uint32(respPtr), uint32(respPtr))
	if !ok {
		panic(ErrOutofMemory)
	}

	return 1
}
