package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/dev"
)

func main() {

	err := dev.Run()
	if err != nil {
		pp.Println(err)
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
