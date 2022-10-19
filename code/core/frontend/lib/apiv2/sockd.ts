import { Websocket } from "../vendor/ws";

export class SockdAPI {
  constructor(baseUrl: string) {
    
  }

  userSocket() {}
  dataSocket() {}
  roomSocket() {}
  devSocket() {}
}

/*

[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/user/ws --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdUserWS-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/data/ws --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdDataWS-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/data/update --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdDataUpdateWS-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/room/ws --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdRoomWS-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/room/update --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdRoomUpdateWS-fm (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/sockd/dev/room/ws --> github.com/temphia/temphia/code/core/backend/app/server.(*Server).sockdDevWS-fm (3 handlers)



*/
