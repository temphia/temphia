package cached

import "github.com/temphia/temphia/code/backend/xtypes/models/entities"

type UserGroupCache interface {
	GetDatas() []*entities.UserGroupData
	SetData(*entities.UserGroupData)

	GetAuths() []*entities.UserGroupAuth
	SetAuth(*entities.UserGroupAuth)

	GetApps() []*entities.TargetApp
	SetApp(*entities.TargetApp)

	GetHooks() []*entities.TargetHook
	SetHook(*entities.TargetHook)
}

type UserGroupService struct {
	Cache UserGroupCache
}
