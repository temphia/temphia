package main

import (
	"io/ioutil"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

func main3() {

	out, err := ioutil.ReadFile("with_ts.js")
	if err != nil {
		panic(err)
	}

	script := string(out)

	vm := goja.New()

	vm.Set("_log", func(msg interface{}) {
		pp.Println(msg)
	})

	vm.Set("_multi_return_err", func(message string) (interface{}, interface{}) {
		/*
			This did not work

			var resp2 = _multi_return_not_err("no err example");
			_log(typeof resp2[0] + "__no_err___" + typeof resp2[1]);

		*/

		return nil, "This is a error"
	})

	vm.Set("_multi_return_not_err", func(message string) (interface{}, interface{}) {
		return 42, nil
	})

	vm.Set("_byte_return", func(message string) (interface{}, interface{}) {
		return []byte(`AAAA`), nil
	})

	vm.Set("_byte_return2", func(message string) (interface{}, interface{}) {
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
