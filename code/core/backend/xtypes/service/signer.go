package service

import "github.com/temphia/temphia/code/core/backend/xtypes/models/claim"

type SignerCore interface {
	GlobalSignRaw(payload string) (string, error)
	GlobalParseRaw(token string) (string, error)
	Sign(ns string, payload any) (string, error)
	Parse(ns string, token string, target any) error
	SignRaw(ns string, payload string) (string, error)
	ParseRaw(ns string, token string) (string, error)
}

type Signer interface {
	//SignerCore
	ClaimSigner
}

type ClaimSigner interface {
	SignOperator(data *claim.Operator) (string, error)
	ParseOperator(payload string) (*claim.Operator, error)
	SignSite(data *claim.Site) (string, error)
	ParseSite(payload string) (*claim.Site, error)
	SignUser(tenantId string, data *claim.User) (string, error)
	ParseUser(tenantId, payload string) (*claim.User, error)
	SignSession(tenantId string, data *claim.Session) (string, error)
	ParseSession(tenantId, payload string) (*claim.Session, error)
	SignExecutor(tenantId string, data *claim.Executor) (string, error)
	ParseExecutor(tenantId, payload string) (*claim.Executor, error)

	// auth related
	SignOauthState(tenantId string, data *claim.OauthState) (string, error)
	ParseOauthState(tenantId, payload string) (*claim.OauthState, error)
	SignPreAuthed(tenantId string, data *claim.PreAuthed) (string, error)
	ParsePreAuthed(tenantId, payload string) (*claim.PreAuthed, error)
	SignAutheFirst(tenantId string, data *claim.AuthFirst) (string, error)
	ParseAutheFirst(tenantId, payload string) (*claim.AuthFirst, error)
	SignAutheNext(tenantId string, data *claim.AuthNext) (string, error)
	ParseAutheNext(tenantId, payload string) (*claim.AuthNext, error)

	// ticket claims

	SignFolderTkt(tenantId string, data *claim.FolderTkt) (string, error)
	ParseFolderTkt(tenantId, payload string) (*claim.FolderTkt, error)

	SignRoomTagTkt(tenantId string, data *claim.RoomTagTkt) (string, error)
	ParseRoomTagTkt(tenantId, payload string) (*claim.RoomTagTkt, error)

	SignPlugDevTkt(tenantId string, data *claim.PlugDevTkt) (string, error)
	ParsePlugDevTkt(tenantId, payload string) (*claim.PlugDevTkt, error)

	SignAdviseryTkt(tenantId string, data *claim.AdviseryTkt) (string, error)
	ParseAdviseryTkt(tenantId, payload string) (*claim.AdviseryTkt, error)

	SignSockdTkt(tenantId string, data *claim.SockdTkt) (string, error)
	ParseSockdTkt(tenantId, payload string) (*claim.SockdTkt, error)

	SignDtableTkt(tenantId string, data *claim.DataTkt) (string, error)
	ParseDtableTkt(tenantId, payload string) (*claim.DataTkt, error)

	SignUserMgmtTkt(tenantId string, data *claim.UserMgmtTkt) (string, error)
	ParseUserMgmtTkt(tenantId, payload string) (*claim.UserMgmtTkt, error)
}
