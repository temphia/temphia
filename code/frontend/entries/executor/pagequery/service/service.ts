import type { Environment } from "../../../../lib/engine/environment";

export class PageQueryService {
  env: Environment;
  constructor(env: Environment) {
    this.env = env;
  }

  load = () => {
    return this.env.PreformAction("load", {});
  };

  submit = (data: any) => {
    return this.env.PreformAction("submit", data);
  };

}
