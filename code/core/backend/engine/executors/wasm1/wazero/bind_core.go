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
		ExportFunction("get_self_file", getSelfFile).
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

func getSelfFile(ctx context.Context, filePtr, fileLen, respPtr, respLen, modPtr int32) int32 {
	e := getCtx(ctx)

	fout, mod, err := e.bindings.GetFileWithMeta(
		e.getString(uint32(filePtr), uint32(fileLen)),
	)
	if err != nil {
		e.writeError(uint32(respPtr), uint32(respLen), err)
		return 0
	}

	e.writeWithOffsetPtr(fout, uint32(respPtr), uint32(respPtr))
	e.mem.WriteUint32Le(e.context, uint32(modPtr), uint32(mod))

	return 1
}
