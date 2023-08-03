package schema

import (
	_ "embed"
)

//go:embed sqlite.sql
var SQLiteSchema string
