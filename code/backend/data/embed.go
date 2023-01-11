package data

import (
	"embed"
)

//go:embed schema
//go:embed schema/postgres
//go:embed schema/postgres.sql
//go:embed templates
//go:embed repo
//go:embed interfaces
//go:embed assets/build
//go:embed assets/lib
//go:embed assets/static
var DataDir embed.FS
