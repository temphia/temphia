import type { ApiBase } from "./base";

export class CabinetAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }
}

/*

[GIN-debug] GET    /z/api/:tenant_id/v2/cabinet/ --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/cabinet/:folder --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/cabinet/:folder --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/cabinet/:folder/file/:fname --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/cabinet/:folder/file/:fname --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/cabinet/:folder/file/:fname --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/cabinet/:folder/preview/:fname --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/cabinet/:folder/ticket --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)



[GIN-debug] GET    /z/api/:tenant_id/v2/tkt/cabinet/:ticket/ --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).folderTktList-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/tkt/cabinet/:ticket/:name --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).folderTktFile-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/tkt/cabinet/:ticket/:name/preview --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).folderTktPreview-fm (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/tkt/cabinet/:ticket/:name --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).folderTktUpload-fm (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/tkt/cabinet/:ticket/:name --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).folderTktDelete-fm (3 handlers)



*/
