package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func main() {
	path := build()

	out, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	buildSubOrigin(string(out))

}

func buildSubOrigin(artifact string) {
	final := fmt.Sprintf("var __dirname = ''; var module = {}; module['exports']={};%s", string(artifact))
	cdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	newPath := path.Join(cdir, "../../../../backend/data/assets/engine_launcher.js")

	err = os.WriteFile(newPath, []byte(final), 0777)
	if err != nil {
		panic(err)
	}

}

func build() string {
	build := path.Join(os.TempDir(), "entry_compiled")
	cmd := exec.Command("ncc", "build", "../altentry/execiframe/index.ts", "--out", build)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	filePath := path.Join(build, "index.js")

	return filePath
}

/*

✦ ❯ make generate_launcher_templates
cd code/frontend/lib/launcher && go run buildentry.go
ncc: Version 0.33.1
ncc: Compiling file index.js into CJS
ncc: Using typescript@3.9.10 (local user-provided)
24kB  ../../../../../../../../../../tmp/entry_compiled/index.js
24kB  [2541ms] - ncc 0.33.1


*/
