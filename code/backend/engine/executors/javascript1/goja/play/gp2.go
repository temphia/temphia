package main

import (
	"os"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

func main3() {

	out, err := os.ReadFile("with_ts.js")
	if err != nil {
		panic(err)
	}

	script := string(out)

	vm := goja.New()

	vm.Set("_log", func(msg any) {
		pp.Println(msg)
	})

	vm.Set("_multi_return_err", func(message string) (any, any) {
		/*
			This did not work

			var resp2 = _multi_return_not_err("no err example");
			_log(typeof resp2[0] + "__no_err___" + typeof resp2[1]);

		*/

		return nil, "This is a error"
	})

	vm.Set("_multi_return_not_err", func(message string) (any, any) {
		return 42, nil
	})

	vm.Set("_byte_return", func(message string) (any, any) {
		return []byte(`AAAA`), nil
	})

	vm.Set("_byte_return2", func(message string) (any, any) {
		return vm.NewArrayBuffer([]byte(`AAAA`)), nil
	})

	_, err = vm.RunString(script)
	if err != nil {
		panic(err)
	}

}

/*

✦ ❯ go run *.go
"Example Err"
"[null,{}]"
"[42,null]"

*/
