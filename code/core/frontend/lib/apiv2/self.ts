import type { ApiBase } from "./base";

export class SelfAPI {
    base: ApiBase;
    constructor(base: ApiBase) {
      this.base = base;
    }

    list_cabinet_sources() {}
    list_data_sources() {}
    list_adapter_providers() {}
    

    load() {}
    sessions() {}
    email_change() {}
    list_message() {}
    modify_message() {}
}


/*

[GIN-debug] GET    /z/api/:tenant_id/v2/self/system/cabinet --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/self/system/datatable --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/self/system/adapter --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/self/load --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/self/session --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/self/email/change --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/self/message --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/self/message --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/self/issue --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)


*/