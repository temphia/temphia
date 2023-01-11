package server

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/vmodels"
)

func (s *Server) AuthIndex(ctx *gin.Context) {
	// fime => move this to template html

	tenantId := s.extractTenant(ctx)
	if tenantId == "" {
		httpx.WriteErrString(ctx, "Could not determine tenant")
		return
	}

	stoken, err := s.siteToken(tenantId, ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	data := &vmodels.SiteData{
		SiteToken: stoken,
		ApiURL:    fmt.Sprintf("http://%s/z/api/%s/v2", ctx.Request.Host, tenantId),
		TenantId:  tenantId,
		UserGroup: s.extractUserGroup(tenantId, ctx),
	}

	tdata, err := json.Marshal(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	ctx.SetCookie("tenant_id", data.TenantId, 60*60*24*30, "/", "", true, true)
	// also attach device_id

	var buf bytes.Buffer

	buf.Write([]byte(`<!DOCTYPE html>
			<html lang="en">
			<head>
			<meta charset="utf-8" />
			<meta name="viewport" content="width=device-width,initial-scale=1" />
			<title>Authed</title>
			<script>
			window.__temphia_site_data__ = `))
	buf.Write(tdata)
	buf.Write([]byte(`</script>
		<link rel="icon" type="image/png" href="/favicon.png" />
		<link rel="stylesheet"  type="text/css" href="/z/assets/build/auth.css" />
		<script defer src="/z/assets/build/auth.js"></script>
		</head>
		<body></body>
		</html>`))

	ctx.Writer.Write(buf.Bytes())
}

func (s *Server) extractTenant(ctx *gin.Context) string {
	if s.app.SingleTenant() {
		return s.app.StaticTenants()[0]
	}

	// fixme => reverse_extract from domain

	tenantId := ctx.Query("tenant_id")
	if tenantId != "" {
		return tenantId
	}
	cookie, _ := ctx.Cookie("tenant_id")
	return cookie
}

func (s *Server) extractUserGroup(tenantId string, ctx *gin.Context) string {
	ugroup := ctx.Query("ugroup")
	if ugroup != "" {
		return ugroup
	}

	return xtypes.UserGroupSuperAdmin
}

func (s *Server) siteToken(tenantId string, ctx *gin.Context) (string, error) {
	// fixme => check if host is valid

	siteClaim := claim.NewSiteClaim(tenantId, ctx.Request.Host)
	return s.signer.SignSite(siteClaim)
}
