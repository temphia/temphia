import type { ApiBase } from "../base";

export class CheckAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  bprint(bid: string) {
    this.base.get(`check/bprint/${bid}`);
  }

  plug(pid: string) {
    this.base.get(`check/bprint/${pid}`);
  }

  dataGroup(gid: string) {
    this.base.get(`check/bprint/${gid}`);
  }

  dataTable() {}
}

/*

[GIN-debug] GET    /z/api/:tenant_id/v2/admin/check/bprint/:bid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/check/plug/:bid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/check/dgroup/:bid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/check/dtable/:bid --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)


*/
