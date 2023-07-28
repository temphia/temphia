package admin

import (
	"fmt"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (c *Controller) AdapterSelfUpdate(aclaim *claim.AdapterEditor, data map[string]any) error {
	return c.coredb.UpdateDomain(aclaim.TenantId, aclaim.AdapterId, data)
}

// app

func (c *Controller) AdapterListApps(aclaim *claim.AdapterEditor) ([]*entities.TargetApp, error) {
	return c.coredb.ListTargetAppByType(aclaim.TenantId, entities.TargetAppTypeDomainWidget, fmt.Sprintf("%d", aclaim.AdapterId))
}

func (c *Controller) AdapterNewApp(aclaim *claim.AdapterEditor, data *entities.TargetApp) error {

	_, err := c.coredb.AddTargetApp(&entities.TargetApp{
		Id:          0,
		Name:        data.Name,
		Icon:        data.Icon,
		Policy:      data.Policy,
		TargetType:  entities.TargetAppTypeDomainWidget,
		Target:      fmt.Sprintf("%d", aclaim.AdapterId),
		ContextType: data.ContextType,
		PlugId:      data.PlugId,
		AgentId:     data.AgentId,
		ExecDomain:  data.ExecDomain,
		ExecMeta:    data.ExecMeta,
		ExtraMeta:   data.ExtraMeta,
		TenantId:    aclaim.TenantId,
	})

	return err

}

func (c *Controller) AdapterGetApp(aclaim *claim.AdapterEditor, id int64) (*entities.TargetApp, error) {
	data, err := c.coredb.GetTargetApp(aclaim.TenantId, entities.TargetAppTypeDomainWidget, id)
	if err != nil {
		return nil, err
	}

	if data.Target != fmt.Sprintf("%d", aclaim.AdapterId) {
		return nil, easyerr.NotAuthorized()
	}
	return data, nil
}

func (c *Controller) AdapterUpdateApp(aclaim *claim.AdapterEditor, id int64, data map[string]any) error {
	if _, ok := data["target"]; ok {
		return easyerr.NotAuthorized()
	}

	return c.coredb.UpdateTargetApp(aclaim.TenantId, entities.TargetAppTypeDomainWidget, id, data)
}

func (c *Controller) AdapterDeleteApp(aclaim *claim.AdapterEditor, id int64) error {
	_, err := c.AdapterGetApp(aclaim, id)
	if err != nil {
		return err
	}

	return c.coredb.RemoveTargetApp(aclaim.TenantId, entities.TargetAppTypeDomainWidget, id)
}

// hook

func (c *Controller) AdapterListHooks(aclaim *claim.AdapterEditor) ([]*entities.TargetHook, error) {
	return c.coredb.ListTargetHookByType(aclaim.TenantId, entities.TargetHookTypeDomainHook, fmt.Sprintf("%d", aclaim.AdapterId))
}

func (c *Controller) AdapterNewHook(aclaim *claim.AdapterEditor, data *entities.TargetHook) error {
	_, err := c.coredb.AddTargetHook(&entities.TargetHook{
		Id:         0,
		Name:       data.Name,
		Policy:     data.Policy,
		TargetType: entities.TargetHookTypeDomainHook,
		Target:     fmt.Sprintf("%d", aclaim.AdapterId),
		EventType:  data.EventType,
		PlugId:     data.PlugId,
		AgentId:    data.AgentId,
		ExecMeta:   data.ExecMeta,
		ExtraMeta:  data.ExtraMeta,
		TenantId:   aclaim.TenantId,
		Handler:    data.Handler,
	})

	return err
}

func (c *Controller) AdapterGetHook(aclaim *claim.AdapterEditor, id int64) (*entities.TargetHook, error) {
	data, err := c.coredb.GetTargetHook(aclaim.TenantId, entities.TargetHookTypeDomainHook, id)
	if err != nil {
		return nil, err
	}

	if data.Target != fmt.Sprintf("%d", aclaim.AdapterId) {
		return nil, easyerr.NotAuthorized()
	}
	return data, nil

}

func (c *Controller) AdapterUpdateHook(aclaim *claim.AdapterEditor, id int64, data map[string]any) error {
	if _, ok := data["target"]; ok {
		return easyerr.NotAuthorized()
	}

	return c.coredb.UpdateTargetHook(aclaim.TenantId, entities.TargetHookTypeDomainHook, id, data)
}

func (c *Controller) AdapterDeleteHook(aclaim *claim.AdapterEditor, id int64) error {
	_, err := c.AdapterGetHook(aclaim, id)
	if err != nil {
		return err
	}

	return c.coredb.RemoveTargetHook(aclaim.TenantId, entities.TargetAppTypeDomainWidget, id)
}

type DomainAdapterEditorIssueResp struct {
	AdapterType string `json:"adapter_type,omitempty"`
	Token       string `json:"token,omitempty"`
	DomainName  string `json:"domain_name,omitempty"`
}

func (c *Controller) DomainAdapterEditorIssue(uclaim *claim.Session, did int64) (*DomainAdapterEditorIssueResp, error) {

	tdomain, err := c.coredb.GetDomain(uclaim.TenantId, did)
	if err != nil {
		return nil, err
	}

	tok, err := c.signer.SignAdapterEditor(uclaim.TenantId, &claim.AdapterEditor{
		TenantId:   uclaim.TenantId,
		AdapterId:  did,
		DomainName: tdomain.Name,
		Type:       "",
		UserID:     uclaim.UserID,
		UserGroup:  uclaim.UserGroup,
		SessionID:  uclaim.SessionID, // fixme => should i share session id cross different claims
		DeviceId:   uclaim.DeviceId,
	})
	if err != nil {
		return nil, err
	}

	return &DomainAdapterEditorIssueResp{
		AdapterType: tdomain.AdapterType,
		Token:       tok,
		DomainName:  tdomain.Name,
	}, nil

}
