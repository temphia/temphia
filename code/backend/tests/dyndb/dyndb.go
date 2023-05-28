package dyndbtest

import (
	"encoding/json"
	"os"

	_ "embed"

	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/config"
	"github.com/temphia/temphia/code/backend/stores/upper/vendors/sqlite"
	"github.com/temphia/temphia/code/backend/xtypes/service/repox/step"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

//go:embed test_example.json
var testex []byte

func Run() {
	testdb := "test.db"

	defer func() {
		os.Remove(testdb)
		os.Remove("test.db-shm")
		os.Remove("test.db-wal")
	}()

	opts := step.MigrateOptions{}
	err := json.Unmarshal(testex, &opts)
	if err != nil {
		panic(err)
	}

	store, err := sqlite.NewVendor(&config.StoreSource{
		Name:     "",
		Vendor:   store.VendorSqlite,
		HostPath: testdb,
	})
	if err != nil {
		panic(err)
	}

	dyndb := store.DynDB()

	opts.Gslug = "test1"

	pp.Println(
		dyndb.MigrateSchema("default0", opts),
	)

}
