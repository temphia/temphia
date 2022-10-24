package cmse

type Func struct {
	Name     string
	PlugId   string
	AgentId  string
	CacheTag uint8
}

type Asset struct {
	Name     string
	URL      string
	Type     string
	Path     string
	AssetTag uint8
}

type Action struct {
	Name     string
	Method   string
	CacheTag uint8
}

type Filter struct {
	Name   string
	Method string
}

type Engine interface {
	RunActions(name string, data any) (any, error)
	RunFilters(name string, data any) (any, error)
}

type CmseInvoker interface {
	RegisterActions([]Action)
	RegisterFilters([]Filter)
	RegisterAssets([]Asset)
}
