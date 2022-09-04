package main

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/app"
	"github.com/temphia/temphia/code/core/backend/app/log"
	"github.com/temphia/temphia/code/core/backend/app/registry"
)

func main() {
	fmt.Println("lets begin the madness")

	builder := app.NewBuilder()
	builder.SetConfig(nil)
	builder.SetLogger(log.New(log.LogOptions{}))
	builder.SetRegistry(registry.G)

}
