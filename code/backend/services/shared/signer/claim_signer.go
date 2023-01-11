package signer

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/hako/branca"
	"github.com/rs/xid"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"

	"golang.org/x/crypto/pbkdf2"
)

// fixme => currently claims are not really expired, OOPS
// fixme => one type of claim could be coherced into aother type of claim

var _ service.ClaimSigner = (*ClaimSigner)(nil)

type ClaimSigner struct {
	signer Signer
}

func New(key []byte, salt string) *ClaimSigner {
	masterKey := pbkdf2.Key(key, []byte(salt), 2048, 32, sha256.New)

	core := Signer{
		masterKey:    masterKey,
		signers:      make(map[string]*branca.Branca),
		globalSigner: branca.NewBranca(string(pbkdf2.Key(key, []byte(salt+"global_key"), 1024, 32, sha256.New))),
	}

	return &ClaimSigner{
		signer: core,
	}
}

func (cs *ClaimSigner) SignOperator(data *claim.Operator) (string, error) {
	data.XID = xid.New().String()
	data.Type = "operator"

	out, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return cs.signer.GlobalSignRaw(kosher.Str(out))
}

func (cs *ClaimSigner) ParseOperator(payload string) (*claim.Operator, error) {

	out, err := cs.signer.GlobalParseRaw(payload)
	if err != nil {
		return nil, err
	}

	c := &claim.Operator{}
	err = json.Unmarshal(kosher.Byte(out), c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (cs *ClaimSigner) SignSite(data *claim.Site) (string, error) {
	out, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return cs.signer.GlobalSignRaw(kosher.Str(out))
}

func (cs *ClaimSigner) ParseSite(payload string) (*claim.Site, error) {

	raw, err := cs.signer.GlobalParseRaw(payload)
	if err != nil {
		return nil, err
	}

	data := &claim.Site{}

	err = json.Unmarshal(kosher.Byte(raw), data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (cs *ClaimSigner) SignUser(tenantId string, data *claim.User) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseUser(tenantId, payload string) (*claim.User, error) {
	data := &claim.User{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignSession(tenantId string, data *claim.Session) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseSession(tenantId, payload string) (*claim.Session, error) {
	data := &claim.Session{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignExecutor(tenantId string, data *claim.Executor) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseExecutor(tenantId, payload string) (*claim.Executor, error) {
	data := &claim.Executor{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}

	data.TenentId = tenantId

	return data, nil
}

func (cs *ClaimSigner) SignAdapterEditor(tenantId string, data *claim.AdapterEditor) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseAdapterEditor(tenantId, payload string) (*claim.AdapterEditor, error) {
	data := &claim.AdapterEditor{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}

	data.TenentId = tenantId

	return data, nil
}

//  auth related

func (cs *ClaimSigner) SignOauthState(tenantId string, data *claim.OauthState) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseOauthState(tenantId, payload string) (*claim.OauthState, error) {
	data := &claim.OauthState{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignPreAuthed(tenantId string, data *claim.PreAuthed) (string, error) {
	return cs.signer.Sign(tenantId, data)
}
func (cs *ClaimSigner) ParsePreAuthed(tenantId, payload string) (*claim.PreAuthed, error) {
	data := &claim.PreAuthed{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignAutheFirst(tenantId string, data *claim.AuthFirst) (string, error) {
	return cs.signer.Sign(tenantId, data)
}
func (cs *ClaimSigner) ParseAutheFirst(tenantId, payload string) (*claim.AuthFirst, error) {
	data := &claim.AuthFirst{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignAutheNext(tenantId string, data *claim.AuthNext) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseAutheNext(tenantId, payload string) (*claim.AuthNext, error) {
	data := &claim.AuthNext{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignData(tenantId string, data *claim.Data) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseData(tenantId, payload string) (*claim.Data, error) {
	data := &claim.Data{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}

	data.TenentId = tenantId

	return data, nil
}

func (cs *ClaimSigner) SignFolder(tenantId string, data *claim.Folder) (string, error) {
	return cs.signer.Sign(tenantId, data)
}
func (cs *ClaimSigner) ParseFolder(tenantId, payload string) (*claim.Folder, error) {
	data := &claim.Folder{}

	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}

	data.TenentId = tenantId

	return data, nil
}

// ticket related

func (cs *ClaimSigner) SignAdviseryTkt(tenantId string, data *claim.AdviseryTkt) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseAdviseryTkt(tenantId, payload string) (*claim.AdviseryTkt, error) {
	data := &claim.AdviseryTkt{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignSockdTkt(tenantId string, data *claim.SockdTkt) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseSockdTkt(tenantId, payload string) (*claim.SockdTkt, error) {
	data := &claim.SockdTkt{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignUserMgmtTkt(tenantId string, data *claim.UserMgmtTkt) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseUserMgmtTkt(tenantId, payload string) (*claim.UserMgmtTkt, error) {
	data := &claim.UserMgmtTkt{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (cs *ClaimSigner) SignRoomTagTkt(tenantId string, data *claim.RoomTagTkt) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParseRoomTagTkt(tenantId, payload string) (*claim.RoomTagTkt, error) {
	data := &claim.RoomTagTkt{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (cs *ClaimSigner) SignPlugDevTkt(tenantId string, data *claim.PlugDevTkt) (string, error) {
	return cs.signer.Sign(tenantId, data)
}

func (cs *ClaimSigner) ParsePlugDevTkt(tenantId, payload string) (*claim.PlugDevTkt, error) {
	data := &claim.PlugDevTkt{}
	err := cs.signer.Parse(tenantId, payload, data)
	if err != nil {
		return nil, err
	}

	data.TenantId = tenantId
	return data, nil
}
