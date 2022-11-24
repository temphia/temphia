package apiadmin

type DataSeed struct {
	Source  string
	Group   string
	Tables  []string
	Columns map[string][]string
}
