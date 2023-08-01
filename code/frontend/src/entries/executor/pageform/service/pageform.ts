import type { Environment } from "../../../../lib/engine/environment";
import type { LoadRequest, Response } from "./pfmodels";

export class PageFormService {
  env: Environment;

  constructor(env: Environment) {
    this.env = env;
  }

  load = async (opts: LoadRequest) => {
    const resp = await this.env.PreformAction("load", opts);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    return resp.data as Response;
  };

  submit = async (stage: string, data: any) => {
    const resp = await this.env.PreformAction("submit", {
      stage,
      data,
    });
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }
    return resp.data as Response;
  };
}
