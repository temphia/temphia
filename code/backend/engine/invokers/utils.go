package invokers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

var (
	ErrInvokerActionNotImplemented = errors.New("INVOKER ACTION NOT IMPLEMENTED")
	ErrInvokerActionNotAllowed     = errors.New("INVOKER ACTION NOT ALLOWED")
)

const (
	TypeWebAdmin           = "web_admin"
	TypeWeb2RPC            = "web2rpc"
	TypeForked             = "forked"
	TypeDtableClientHook   = "dtable_client"
	TypeDtableServerHook   = "dtable_server"
	TypeUserApp            = "user_app"
	TypeAuthedServerHook   = "authed_server"
	TypeAuthedClientHook   = "authed_client"
	TypeDomainAdapter      = "domain_adapter" // web2rpc
	TypeDomainEditor       = "domain_editor"
	TypeDomainClientWidget = "domain_client_widget"
	TypeDomainServerWidget = "domain_server_widget"
)

func ParseClaim(signer service.Signer, ctx *gin.Context) (*claim.Session, error) {
	return signer.ParseSession(ctx.Param("tenant_id"), ctx.GetHeader("Authorization"))
}
