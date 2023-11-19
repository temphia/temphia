package local

type LocalCLI struct {
	Zip    struct{} `cmd:"" help:"Zip and package bashed on brpint.yaml."`
	Status struct{} `cmd:"" help:"Prints the current context for bprint actions."`
}
