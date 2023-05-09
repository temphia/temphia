package notz

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/logx/logid"
)

// global middleware

func (m *Notz) serve(ctx *gin.Context) {

	tenantId, hostname, err := m.extract(ctx)
	if err != nil {
		m.adapterManager.applogger.Error().
			Str("tenant_id", tenantId).
			Msg(logid.NotzHostExtractErr)
		return
	}

	m.adapterManager.Handle(tenantId, hostname, ctx)
}

func (m *Notz) servePublic(c *gin.Context, file string) {
	tenantId, _, err := m.extract(c)
	if err != nil {
		return
	}

	source := m.cabinethub.Default(tenantId)
	out, err := source.GetBlob(c.Request.Context(), "public", file)
	if err != nil {
		return
	}

	c.Writer.Write(out)
}

func (m *Notz) isAllowed(tenantId, host string) bool {
	// fixme => actually implement
	// cors check and domain_bashed checks

	// did := am.domainId(tenantId, host)
	// instance := am.get(tenantId, did)

	// err := instance.middleWare(ctx)
	// if err != nil {
	// 	return
	// }

	return false
}
