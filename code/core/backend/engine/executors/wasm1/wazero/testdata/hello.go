package main

import tasmsdk "github.com/temphia/temphia/core/backend/server/engine/executors/wasm1/sdk"

func main() {}

// _greet is a WebAssembly export that accepts a string pointer (linear memory
// offset) and calls greet.
//
//export action_main
func actionMain(ptr, size uint32) {
	tasmsdk.Log(string(tasmsdk.GetBytes(int32(ptr))))
}

// tinygo build -o hello.wasm -scheduler=none -target=wasi hello.go
