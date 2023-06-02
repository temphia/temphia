package wazero

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func BuildTemphiaModule(ctx context.Context, runtime wazero.Runtime) (api.Module, error) {

	mb := runtime.NewHostModuleBuilder("temphia")

	// core

	mb.NewFunctionBuilder().WithFunc(log).Export("log")
	mb.NewFunctionBuilder().WithFunc(sleep).Export("sleep")
	mb.NewFunctionBuilder().WithFunc(getSelfFile).Export("get_self_file")

	// self

	mb.NewFunctionBuilder().WithFunc(SelfListResources).Export("self_list_resources")
	mb.NewFunctionBuilder().WithFunc(SelfGetResource).Export("self_get_resource")
	mb.NewFunctionBuilder().WithFunc(SelfInLinks).Export("self_in_links")
	mb.NewFunctionBuilder().WithFunc(SelfOutLinks).Export("self_out_links")
	mb.NewFunctionBuilder().WithFunc(SelfLinkExec).Export("self_link_exec")
	mb.NewFunctionBuilder().WithFunc(SelfNewModule).Export("self_new_module")
	mb.NewFunctionBuilder().WithFunc(SelfModuleExec).Export("self_module_exec")
	mb.NewFunctionBuilder().WithFunc(SelfForkExec).Export("self_fork_exec")

	return mb.Instantiate(ctx)
}

// core binds

func log(ctx context.Context, offset, len int32) {
	e := getCtx(ctx)
	out, ok := e.mem.Read(uint32(offset), uint32(len))
	if !ok {
		panic(ErrOutofIndex)
	}

	e.bindings.Log(string(out))
}

func sleep(ctx context.Context, msec int32) {
	e := getCtx(ctx)
	e.bindings.Sleep(msec)
}

func getSelfFile(ctx context.Context, ctxid, filePtr, fileLen, respOffset, respLen, modPtr int32) int32 {
	e := getCtx(ctx)

	fout, mod, err := e.bindings.GetFileWithMeta(
		e.getString((filePtr), (fileLen)),
	)
	if err != nil {
		e.writeError(respOffset, respLen, err)
		return 0
	}

	e.writeBytesNPtr(fout, respOffset, respOffset)
	e.mem.WriteUint32Le(uint32(modPtr), uint32(mod))

	return 1
}
