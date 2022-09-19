import type { PlugAPI } from "../../core/api";
import type { AppBase } from "../../service/base";
import { EngineService } from "../../service/engine";
import { IframeExec } from "../../service/engine/exec_iframe";
import { generateId } from "../../utils";

export interface InstanceOpts {
  plug: string;
  agent: string;
  data?: any;
  handler: (exec_id: string, action: string, data: any) => void;
  target: HTMLElement;
}

export class EngineLauncher {
  ecore: EngineService;
  eapi?: PlugAPI;
  _base_app: AppBase;

  constructor(bapp: AppBase) {
    this.ecore = new EngineService({
      ping: (eid: string, data: any) => {
        console.log("PING RECIEVED FROM", eid, "@data=>", data);
      },
    });
    this._base_app = bapp;
  }

  instance = async (opts: InstanceOpts) => {
    if (!this.eapi) {
      this.eapi = await this._base_app.apm.get_plug_api();
    }

    const resp = await this.eapi.launch_agent(opts.plug, opts.agent, {});
    if (resp.status !== 200) {
      console.log("Error =>", resp);
      return;
    }

    const exec = new IframeExec({
      plug: opts.plug,
      agent: opts.agent,
      engine_data: resp.data,
      handler: opts.handler,
      secret: generateId(),
      target: opts.target,
    });
    exec.run(opts.data);
    return exec;
  };
}
