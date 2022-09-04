package main

import (
	"fmt"

	"github.com/temphia/temphia/code/core/backend/app"
	"github.com/temphia/temphia/code/core/backend/app/logservice"
	"github.com/temphia/temphia/code/core/backend/app/registry"
)

func main() {
	fmt.Println("lets begin the madness")

	builder := app.NewBuilder()
	builder.SetConfig(nil)
	builder.SetLogger(logservice.New(logservice.LogOptions{}))
	builder.SetRegistry(registry.G)

}
