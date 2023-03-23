import type { Environment } from "../../../../lib/engine/environment";

export class PageDashService {
  env: Environment;

  constructor(env: Environment) {
    this.env = env;
  }

  load = async (opts: any) => {
    const resp = await this.env.PreformAction("load", opts);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    return resp.data as Response;
  };
}
