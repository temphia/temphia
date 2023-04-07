package groupd

import (
	"sync"
	"time"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
	"github.com/temphia/temphia/code/backend/xtypes/xplane"
)

type instance struct {
	corehub store.CoreHub

	tenantId string
	group    string
	sLock    sync.Mutex
	selfMod  bool

	lastUpdate time.Time

	model     *entities.UserGroup
	userApps  []*entities.TargetApp
	userDatas []*entities.UserGroupData
	userAuths []*entities.UserGroupAuth
}

func (i *instance) handle(mtype string, msg xplane.Message) {
	i.sLock.Lock()
	if i.selfMod {
		i.sLock.Unlock()
		return
	}
	i.selfMod = true
	i.sLock.Unlock()

	i.fetchLatest()

	i.sLock.Lock()
	i.selfMod = false
	i.sLock.Unlock()
}

func (i *instance) fetchLatest() {
	gmodel, err := i.corehub.GetUserGroup(i.tenantId, i.group)
	if err == nil {
		i.model = gmodel
	}

	auths, err := i.corehub.ListUserGroupAuth(i.tenantId, i.group)
	if err == nil {
		i.userAuths = auths
	}

	apps, err := i.corehub.ListTargetAppByUgroup(i.tenantId, i.group)
	if err == nil {
		i.userApps = apps
	}

	gdatas, err := i.corehub.ListUserGroupData(i.tenantId, i.group)
	if err == nil {
		i.userDatas = gdatas

	}

	i.lastUpdate = time.Now()
}

func (i *instance) CheckScope() error {
	return nil
}
