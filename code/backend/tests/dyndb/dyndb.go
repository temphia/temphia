package dyndbtest

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	_ "embed"

	"github.com/rs/zerolog"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/data"
	"github.com/temphia/temphia/code/backend/stores/upper/vendors/sqlite"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

//go:embed test_example.json
var testex []byte

var testdb = "test.db"

func clean() {
	os.Remove(testdb)
	os.Remove("test.db-shm")
	os.Remove("test.db-wal")
}

func Run() {
	clean()

	//	defer clean()

	opts := step.MigrateOptions{}
	err := json.Unmarshal(testex, &opts)
	if err != nil {
		panic(err)
	}

	fullstep := opts.Steps
	opts.Steps = opts.Steps[:1]

	opts.BprintId = "testbp1"
	opts.BprintInstanceId = "inst1"
	opts.BprintItemId = "data1"

	store, err := sqlite.NewVendor(store.BuilderOptions{
		Config: &config.StoreSource{
			Name:     "",
			Vendor:   store.VendorSqlite,
			HostPath: testdb,
		},
		LogBuilder: func() zerolog.Logger {
			return zerolog.New(os.Stdout)
		},
	})
	if err != nil {
		panic(err)
	}

	dyndb := store.DynDB()

	sdriver := dyndb.GetDriver().(*sql.DB)

	out, err := data.DataDir.ReadFile("schema/sqlite.sql")
	if err != nil {
		panic(err)
	}

	_, err = sdriver.ExecContext(context.Background(), string(out))
	if err != nil {
		panic(err)
	}

	opts.Gslug = "test1"
	opts.New = true

	err = dyndb.MigrateSchema("default0", opts)
	if err != nil {
		panic(err)
	}

	opts.New = false
	opts.Steps = fullstep

	err = dyndb.MigrateSchema("default0", opts)
	if err != nil {
		panic(err)
	}
}
