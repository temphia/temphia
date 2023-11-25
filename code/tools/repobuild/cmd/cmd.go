package cmd

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/temphia/temphia/code/tools/repobuild/builder"
	"github.com/temphia/temphia/code/tools/repobuild/models"
)

func Run(repofile string) error {

	cbytes, err := os.ReadFile(repofile)
	if err != nil {
		return err
	}

	conf := &models.BuildConfig{}
	err = toml.Unmarshal(cbytes, conf)
	if err != nil {
		return err
	}

	builder := builder.New(conf)

	err = builder.Build()
	if err != nil {
		return err
	}

	builder.PrintResult()

	return nil

}
