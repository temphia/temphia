package cfolder

import "github.com/temphia/temphia/code/climux/bdev/core"

type CFolderCLI struct {
	List         struct{} `cmd:"" help:"list files inside this bprint/app."`
	UploadFile   struct{} `cmd:"" help:"upload file into bprint."`
	DownloadFile struct{} `cmd:"" help:"download file into bprint."`
	DeleteFile   struct{} `cmd:"" help:"delete file into bprint."`
}

func (c *CFolderCLI) Run(ctx core.BdevContext) error { return nil }
