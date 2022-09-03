package controllers

import (
	"github.com/temphia/temphia/code/core/backend/controllers/admin"
	"github.com/temphia/temphia/code/core/backend/controllers/authed"
	"github.com/temphia/temphia/code/core/backend/controllers/basic"
	"github.com/temphia/temphia/code/core/backend/controllers/cabinet"
	"github.com/temphia/temphia/code/core/backend/controllers/engine"
	"github.com/temphia/temphia/code/core/backend/controllers/operator"
	"github.com/temphia/temphia/code/core/backend/controllers/repo"
)

type Controllers struct {
	cAdmin    *admin.Controller
	cAuthd    *authed.Controller
	cBasic    *basic.Controller
	cCabinet  *cabinet.Controller
	cEngine   *engine.Controller
	cOperator *operator.Controller
	cRepo     *repo.Controller
}

func New() *Controllers {

	return &Controllers{
		cAdmin:    nil,
		cAuthd:    nil,
		cBasic:    nil,
		cCabinet:  nil,
		cEngine:   nil,
		cOperator: nil,
		cRepo:     nil,
	}
}
