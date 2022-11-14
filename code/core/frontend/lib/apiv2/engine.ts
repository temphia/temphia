export class EngineAPI {
    constructor() {}
}

/*

[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/data --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/user --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/admin --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/domain --> github.com/temphia/temphia/code/core/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)

[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/authd --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).launchAuthd-fm (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/engine/launch/widget --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).launchWidget-fm (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/engine/execute/:action --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).execute-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/engine/plug/:pid/agent/:aid/serve/:file --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).agentServeFile-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/engine/plug/:pid/agent/:aid/executor/:eid/:file --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).executorFile-fm (3 handlers)


  async room(token: string, on_handler: (message: SockdMessage) => void) {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: `${this.base_url}/sockd/room/ws?token=${token}`,
    });

    await sockd.init();
    return sockd;
  }


*/