package autotarget

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
	"github.com/temphia/temphia/code/backend/xtypes/store"
)

type AutoTarget struct {
	corehub  store.CoreHub
	domainId int64
	tenantId string

	// loaded data

	editorHook *entities.TargetHook
	bprintid   string
}

func New(tenantId string, domainId int64, corehub store.CoreHub) *AutoTarget {

	return &AutoTarget{
		corehub:  corehub,
		domainId: domainId,
		tenantId: tenantId,
	}
}

func (a *AutoTarget) Reset() error {

	return nil
}

func (a *AutoTarget) IsInit() (bool, error) {

	hooks, err := a.corehub.ListTargetHookByType(a.tenantId, entities.TargetHookDomainEditor, fmt.Sprint(a.domainId))
	if err != nil {
		return false, err
	}
	if len(hooks) == 0 {
		return false, nil
	}

	hook := hooks[0]

	plug, err := a.corehub.PlugGet(a.tenantId, hook.PlugId)
	if err != nil {
		return false, err
	}

	a.bprintid = plug.BprintId
	a.editorHook = hook

	return false, nil
}

func (a *AutoTarget) InitWithZip() error {

	return nil
}

func (a *AutoTarget) AutoInit() error {
	key := a.key()

	bp, _ := a.corehub.BprintGet(a.tenantId, key)
	if bp == nil {
		err := a.corehub.BprintNew(a.tenantId, &entities.BPrint{
			ID:          key,
			Name:        "Domain Adpter Container",
			Slug:        "domain-adapter",
			Type:        "container",
			Description: fmt.Sprintf("This is a bprint controlled by adapter %d", a.domainId),
			TenantID:    a.tenantId,
			Tags:        entities.JsonArray{},
			Files:       entities.JsonArray{},
			ExtraMeta:   entities.JsonMap{},
		})
		if err != nil {
			return err
		}
	}

	plug, _ := a.corehub.PlugGet(a.tenantId, key)
	if plug == nil {
		err := a.corehub.PlugNew(a.tenantId, &entities.Plug{
			Id:        key,
			Name:      fmt.Sprintf("Domain Adpter plug %d", a.domainId),
			Live:      true,
			BprintId:  key,
			ExtraMeta: entities.JsonStrMap{},
			TenantId:  a.tenantId,
		})
		if err != nil {
			return err
		}

		err = a.corehub.AgentNew(a.tenantId, &entities.Agent{
			Id:        "default",
			Executor:  "adapter_editor",
			ExtraMeta: entities.JsonStrMap{},
			TenantId:  a.tenantId,
			PlugId:    key,
			WebFiles:  entities.JsonStrMap{},
			Type:      "web",
		})

		if err != nil {
			return err
		}
	}

	a.corehub.AddTargetApp(&entities.TargetApp{
		Name:        "Adapter Editor",
		ContextType: "editor.main",
		TargetType:  entities.TargetAppTypeDomainEditor,
		Target:      key,
		PlugId:      key,
		ExecMeta:    entities.JsonStrMap{},
		ExtraMeta:   entities.JsonStrMap{},
		AgentId:     "default",
		TenantId:    a.tenantId,
	})

	hid, err := a.corehub.AddTargetHook(&entities.TargetHook{
		Name:       "Adapter Editor",
		TargetType: entities.TargetAppTypeDomainEditor,
		EventType:  "editor.main",
		Handler:    "",
		Target:     key,
		PlugId:     key,
		ExecMeta:   entities.JsonStrMap{},
		ExtraMeta:  entities.JsonStrMap{},
		AgentId:    "default",
		TenantId:   a.tenantId,
	})

	if err != nil {
		return err
	}

	hook, err := a.corehub.GetTargetHook(a.tenantId, entities.TargetAppTypeDomainEditor, hid)
	if err != nil {
		return err
	}

	a.editorHook = hook
	a.bprintid = key

	return nil
}

func (a *AutoTarget) BprintId() string {
	return a.bprintid
}

func (a *AutoTarget) EditorHooks() *entities.TargetHook {
	return a.editorHook
}

// private

func (a *AutoTarget) key() string {
	return fmt.Sprintf("adapter-%d", a.domainId)
}
