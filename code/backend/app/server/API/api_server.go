package api_server

import (
	"net"

	"github.com/bwmarrin/snowflake"
	apidata "github.com/temphia/temphia/code/backend/app/server/API/data"
	"github.com/temphia/temphia/code/backend/app/server/API/middleware"
	"github.com/temphia/temphia/code/backend/app/server/API/tickets"
	"github.com/temphia/temphia/code/backend/controllers/authed"
	"github.com/temphia/temphia/code/backend/controllers/basic"
	"github.com/temphia/temphia/code/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/backend/controllers/data"
	"github.com/temphia/temphia/code/backend/controllers/dev"
	"github.com/temphia/temphia/code/backend/controllers/engine"
	"github.com/temphia/temphia/code/backend/controllers/repo"
	"github.com/temphia/temphia/code/backend/controllers/sockd"
	"github.com/temphia/temphia/code/backend/controllers/user"
	"github.com/temphia/temphia/code/backend/xtypes/service"
)

type Server struct {
	signer service.Signer

	middleware *middleware.Middleware

	listener net.Listener

	apidata *apidata.Data

	// controllers

	cAuth    *authed.Controller
	cBasic   *basic.Controller
	cUser    *user.Controller
	cData    *data.Controller
	cCabinet *cabinet.Controller
	cRepo    *repo.Controller
	cEngine  *engine.Controller
	cDev     *dev.Controller
	cSockd   *sockd.Controller

	ticketsAPI *tickets.TicketAPI
	idNode     *snowflake.Node // sockdConnIdGenerator
}
