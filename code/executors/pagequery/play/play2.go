package main

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

type MyValue2 struct {
	Test string `json:"test,omitempty"`
}

func main() {

	rt := goja.New()
	rt.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

	/*
		`

		const resp = {}

		const mv = {
			"test": "mnop"
		};
		mv;


		`
	*/

	const hook = `




	function hook() {

		throw new Error('This is how its done!');


		const mv = {
			"test": "mnop"
		};

		return mv;
	}

	hook();
	
	`

	val, err := rt.RunString(hook)
	if err != nil {
		pp.Println("ERR", err.Error())
		return
	}

	mv := &MyValue2{}

	mval := map[string]any{}

	fmt.Println("@export_val", val)

	pp.Println(rt.ExportTo(val, mv))
	pp.Println(rt.ExportTo(val, &mval))

	pp.Println("@export_to", mv)

	pp.Println("@export_to", mval)

}

// OUTPUT
// "ERR" "Error: This is how its done! at hook (<eval>:8:9(3))"
