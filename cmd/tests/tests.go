package main

import (
	"github.com/k0kubun/pp"
	dyndbtest "github.com/temphia/temphia/code/backend/tests/dyndb"
)

func main() {
	pp.Println("Starting Test")

	dyndbtest.Run()

	pp.Println("End Test")

}
