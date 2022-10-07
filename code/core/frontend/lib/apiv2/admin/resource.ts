import type { ApiBase } from "../base";

export class ResourceAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }
}
/*


[GIN-debug] GET    /z/api/:tenant_id/v2/admin/resource/ --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/resource/ --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/resource/:slug --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/resource/:slug --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/admin/resource/:slug --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)


*/