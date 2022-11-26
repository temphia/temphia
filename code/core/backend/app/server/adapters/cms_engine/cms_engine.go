package cmsengine

import "github.com/temphia/temphia/code/core/backend/xtypes/models/entities"

type CMSEngine struct {
	Hooks   []*entities.TargetHook
	Apps    []*entities.TargetApp
	PageMap map[string]PageRoute
}
