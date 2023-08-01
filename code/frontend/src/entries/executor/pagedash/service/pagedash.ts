import type { Environment } from "../../../../lib/engine/environment";

export class PageDashService {
  env: Environment;

  constructor(env: Environment) {
    this.env = env;
  }

  load = (opts: any) => {
    return this.env.PreformAction("load", opts);
  };
}
