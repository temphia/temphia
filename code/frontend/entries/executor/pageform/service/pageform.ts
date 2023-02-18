import type { Environment } from "../../../../lib/engine/env";

export interface LoadRequest {
  data_context_type: string;
  rows: number[];
  options: { [_: string]: any };
}

export interface Response {
  ok: boolean;
  message: string;
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
  stage: string;
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
