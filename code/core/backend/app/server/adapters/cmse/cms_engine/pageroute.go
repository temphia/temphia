package cmsengine

const (
	PageRouteDynamic = "dynamic"
	PageRouteStatic  = "static"
	PageRoutePolicy  = "policy"
)

type PageRoute struct {
	Type    string
	Methods []string
}
