package engineapi

type RouteItemMode int8

const (
	RouteItemModeRPX RouteItemMode = iota
	RouteItemModeRaw
	RouteItemModeServe
)

type RouteConfig struct {
	Type        string
	ApppendHTML bool // append .html to no file  /xyz => /xyz.html
	Items       []RouteItem
}

type RouteItem struct {
	Path        string
	Method      string
	Mode        RouteItemMode
	ApppendHTML bool
	TrimSlash   bool
}
