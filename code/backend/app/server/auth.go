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

	stoken, err := s.siteToken(s.opts.TenantId, ctx)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	data := &vmodels.SiteData{
		SiteToken: stoken,
		ApiURL:    fmt.Sprintf("http://%s/z/api/%s/v2", ctx.Request.Host, s.opts.TenantId),
		TenantId:  s.opts.TenantId,
		UserGroup: s.extractUserGroup(ctx),
	}

	tdata, err := json.Marshal(data)
	if err != nil {
		httpx.WriteErr(ctx, err)
		return
	}

	// fixme => also attach device_id

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

		<link rel="icon" type="image/png" href="/z/assets/static/logo.png">
		<link rel="stylesheet"  type="text/css" href="/z/assets/build/auth.css" />
		<script defer src="/z/assets/build/auth.js"></script>
		</head>
		<body></body>
		</html>`))

	ctx.Writer.Write(buf.Bytes())
}

func (s *Server) extractUserGroup(ctx *gin.Context) string {
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
