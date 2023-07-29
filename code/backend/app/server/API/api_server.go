package api_server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/temphia/temphia/code/backend/app/server/API/middleware"
	"github.com/temphia/temphia/code/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/backend/controllers/dev"
	"github.com/temphia/temphia/code/backend/controllers/engine"
	"github.com/temphia/temphia/code/backend/controllers/repo"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Server struct {
	signer service.Signer

	middleware *middleware.Middleware

	// controllers
	cCabinet *cabinet.Controller
	cRepo    *repo.Controller
	cEngine  *engine.Controller
	cDev     *dev.Controller
	cSockd   *sockd.Controller

	idNode *snowflake.Node // sockdConnIdGenerator
}
