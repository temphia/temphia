import type { EnvApi } from "./envapi";

export class GroupPlayer {
  parent_app: any;
  api: EnvApi;
  env: any;
  client_player?: ClientPlayer;
  constructor(app: any, api: EnvApi, env: any, start_data: any) {
    this.parent_app = app;
    this.api = api;
    this.env = env;
    if (start_data["client_side"]) {
      this.client_player = new ClientPlayer(api, env, start_data);
    }
  }
}

export class ClientPlayer {
  api: EnvApi;
  env: any;
  start_data: any;
  constructor(api: EnvApi, env: any, start_data: any) {
    this.api = api;
    this.env = env;
    this.start_data = start_data;
  }
}
