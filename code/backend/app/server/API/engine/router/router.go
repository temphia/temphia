package router

const (
	RouteItemModeRPX   = "rpx"
	RouteItemModeRaw   = "raw_http"
	RouteItemModeServe = "serve"
)

type RouteConfig struct {
	Type  string      `json:"type,omitempty"`
	Items []RouteItem `json:"items,omitempty"`
}

type RouteItem struct {
	Path        string `json:"path,omitempty"`
	Method      string `json:"method,omitempty"`
	Mode        string `json:"mode,omitempty"`
	Target      string `json:"target,omitempty"`
	File        string `json:"file,omitempty"`
	ApppendHTML bool   `json:"append_html,omitempty"` // append .html to no file  /xyz => /xyz.html
	TrimSlash   bool   `json:"trim_slash,omitempty"`
	Wildcard    bool   `json:"wildcard,omitempty"`
}

type RouteResponse struct {
	Mode   string `json:"mode,omitempty"`
	Target string `json:"target,omitempty"`
	File   string `json:"file,omitempty"`
}

func (r *RouteItem) Copy() *RouteResponse {
	ri := &RouteResponse{
		Mode:   r.Mode,
		Target: r.Target,
		File:   r.File,
	}

	return ri
}
