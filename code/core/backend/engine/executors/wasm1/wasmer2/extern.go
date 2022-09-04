package wasmer2

func (w *wasmer2) buildBindings() {
	w.buildCore()
	w.buildPlugKV()
	//	w.buildSockd()
	w.buildCabinet()
}
