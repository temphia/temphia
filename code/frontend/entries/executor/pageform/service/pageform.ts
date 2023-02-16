import type { Environment } from "../../../../lib/engine/env";

export interface LoadRequest {
  data_context_type: string;
  rows: number[];
  options: { [_: string]: any };
}

export interface LoadResponse {
  items: {
    name: string;
    type: string;
    info: string;
    html_attr: { [_: string]: string };
    options: string[];
  }[];
  data: any;
  on_load: string;
  on_submit: string;
}

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

    return resp.data as LoadResponse;
  };

  save = async () => {};
}
