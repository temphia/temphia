package store

type StoreType string

const (
	TypeCoreDB    StoreType = "core_db"
	TypeDynDB     StoreType = "dyn_db"
	TypeStateDB   StoreType = "state_db"
	TypeSyncDB    StoreType = "sync_db"
	TypeBlobStore StoreType = "blob_store"
)

const (
	VendorSqlite   = "sqlite"
	VendorQL       = "ql"
	VendorPostgres = "postgres"
)

type Store interface {
	Supports(StoreType) bool
	CoreDB() CoreDB
	DynDB() DynDB
	StateDB() PlugStateKV
	CabinetSource() CabinetSource
	Name() string

	/*
		// ref

		CheckMigrate() error
		Migrate(comp string) error
		EnsureTenant() error
		CheckTenant() error
	*/
}
