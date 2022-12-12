import type { Writable } from "svelte/store";
import { EnvApi } from "./envapi";
import type { GroupPlayer } from "./player";

export class WizardApp {
  status: Writable<string>;
  group_player: GroupPlayer;
  env: any;
  env_api: EnvApi;
  exec_data?: any;
  constructor(env: any, exec_data?: any) {
    this.env = env;
    this.env_api = new EnvApi(env);
    this.exec_data = exec_data;
  }

  init = async () => {
    const resp = await this.env_api.get_splash(!!this.exec_data);
  };
}
