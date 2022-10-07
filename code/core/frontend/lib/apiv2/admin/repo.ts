import type { ApiBase } from "../base";

export class RepoAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {}
  new() {}
  get() {}
  update() {}
  delete() {}
}

/*

[GIN-debug] GET    /z/api/:tenant_id/v2/admin/repo/ --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/repo/ --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/repo/:rid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/repo/:rid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/admin/repo/:rid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)


*/