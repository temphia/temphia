package admin

import (
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (c *Controller) AdapterSelfUpdate(aclaim *claim.AdapterEditor, data map[string]any) {

}

// app

func (c *Controller) AdapterListApps(aclaim *claim.AdapterEditor) ([]*entities.TargetApp, error) {
	return nil, nil
}

func (c *Controller) AdapterNewApp(aclaim *claim.AdapterEditor, data *entities.TargetApp) error {
	return nil
}

func (c *Controller) AdapterGetApp(aclaim *claim.AdapterEditor, id int64) (*entities.TargetApp, error) {
	return nil, nil
}

func (c *Controller) AdapterUpdateApp(aclaim *claim.AdapterEditor, id int64, data map[string]any) error {

	return nil
}

func (c *Controller) AdapterDeleteApp(aclaim *claim.AdapterEditor, id int64) error {

	return nil
}

// hook

func (c *Controller) AdapterListHooks(aclaim *claim.AdapterEditor) ([]*entities.TargetHook, error) {
	return nil, nil
}

func (c *Controller) AdapterNewHook(aclaim *claim.AdapterEditor, data *entities.TargetHook) error {
	return nil
}

func (c *Controller) AdapterGetHook(aclaim *claim.AdapterEditor, id int64) (*entities.TargetApp, error) {
	return nil, nil
}

func (c *Controller) AdapterUpdateHook(aclaim *claim.AdapterEditor, id int64, data map[string]any) error {
	return nil
}

func (c *Controller) AdapterDeleteHook(aclaim *claim.AdapterEditor, id int64) error {
	return nil
}

type DomainAdapterEditorIssueResp struct {
	AdapterType string `json:"adapter_type,omitempty"`
	Token       string `json:"token,omitempty"`
}

func (c *Controller) DomainAdapterEditorIssue(uclaim *claim.Session, did int64) (*DomainAdapterEditorIssueResp, error) {

	tdomain, err := c.coredb.GetDomain(uclaim.TenentId, did)
	if err != nil {
		return nil, err
	}

	tok, err := c.signer.SignAdapterEditor(uclaim.TenentId, &claim.AdapterEditor{
		TenentId:   uclaim.TenentId,
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
	}, nil

}
