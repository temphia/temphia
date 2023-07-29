package goja

import (
	_ "embed"

	"github.com/dop251/goja"
)

var libesplug *goja.Program

//go:embed libesplug.js
var libesPlug string

func init() {
	program, err := goja.Compile("libesplug", libesPlug, true)
	if err != nil {
		panic(err)
	}
	libesplug = program
}
