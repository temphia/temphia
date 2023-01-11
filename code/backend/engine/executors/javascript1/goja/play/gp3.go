package main

import (
	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

func checkStringToByteConvertFnArg() {

	const s2 = `
	funcme("hahhah")

	`

	vm := goja.New()

	vm.Set("funcme", func(arg []byte) {
		pp.Println(string(arg))
	})
	_, err := vm.RunString(s2)
	if err != nil {
		panic(err)
	}

	// result => yes it passes string to []byte
}
