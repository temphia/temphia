package main

import (
	"os"
	"testing"

	"github.com/dop251/goja"
)

func TestBuild(t *testing.T) {
	t.Logf("smoke test build js files....\n")

	rt := goja.New()

	rt.RunString(`var __dirname = ""; var module = {};`)

	fout, err := os.ReadFile("dist/libesplug.js")
	if err != nil {
		return
	}

	val, err := rt.RunString(string(fout))
	if err != nil {
		t.Fatal("FAILED", val, err)
	} else {
		t.Log("OK ", val)
	}
}
