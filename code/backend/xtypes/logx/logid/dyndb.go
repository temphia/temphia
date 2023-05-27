package logid

const (
	DyndbMigrationErr     = "dyndb_migration_err"
	DyndbMigrationOk      = "dyndb_migration_ok"
	DyndbSetMigHeadErr    = "dyndb_set_mighead_err"
	DyndbUpdateMigHeadErr = "dyndb_update_mighead_err"

	DyndbNewGroupRollBack = "new_group_rollback"
	DyndbGlobalLockErr    = "global_lock_err"

	DyndbNewGroupMetadataCreated   = "new_group_metadata_created"
	DyndbNewGroupMetadataCreateErr = "new_group_metadata_created_err"
	DyndbNewGroupSchemaExecErr     = "new_group_schema_exec_err"

	DyndbColumnsCleanupErr = "columns_cleanup_err"
	DyndbTablesCleanupErr  = "tables_cleanup_err"
	DyndbTableCleanupErr   = "table_cleanup_err"
	DyndbGroupCleanupErr   = "group_cleanup_err"
)
