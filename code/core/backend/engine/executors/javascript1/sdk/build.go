package main

// build helper file

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	generateD()
	generateJS()

}

func generateD() {
	fmt.Println("generating decleration")

	var files = []string{"cabinet.d.ts", "core.d.ts", "http.d.ts", "plugkv.d.ts", "sockd.d.ts", "utils.d.ts"}

	var buf bytes.Buffer

	for _, f := range files {
		file := fmt.Sprintf("dist/lib/%s", f)

		out, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}
		buf.Write(out)
	}

	err := os.WriteFile("dist/libesplug.d.ts", buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}

}

func generateJS() {
	out, err := os.ReadFile("dist/lib/index.js")
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	var buf2 bytes.Buffer

	buf2.Write([]byte(`var __dirname = ""; var module = {};`))
	buf2.WriteString("\n")
	buf2.Write(out)
	buf2.WriteString(`Object.assign(globalThis, module.exports);`)

	err = os.WriteFile("dist/libesplug.js", buf2.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
