package dyndb

type DyndbCLI struct {
	List struct{} `cmd:"" help:"list dydbs inside this bprint/app."`

	Status struct {
		Id string `arg:"" help:"Id to check status."`
	} `cmd:"" help:"status of a dydb inside this bprint/app."`

	Migrate  struct{} `cmd:"" help:"migrate dydb."`
	Rollback struct{} `cmd:"" help:"rollback dydb."`
	Autoseed struct{} `cmd:"" help:"autoseed dydb."`
	Truncate struct{} `cmd:"" help:"truncate dydb."`
	Peek     struct{} `cmd:"" help:"get random records from dyndb."`
	Query    struct{} `cmd:"" help:"query records from dyndb."`
}
