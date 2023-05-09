package notz

import "github.com/gin-gonic/gin"

// global middleware

func (m *Notz) isAllowed(tenantId, host string) bool { return false }

func (am *AdapterManager) MiddleWare(tenantId, host string, ctx *gin.Context) {

	// fixme => actually implement
	// cors check and domain_bashed checks

	// did := am.domainId(tenantId, host)
	// instance := am.get(tenantId, did)

	// err := instance.middleWare(ctx)
	// if err != nil {
	// 	return
	// }

	ctx.Next()

}
