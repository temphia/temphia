import { ExecAPI } from "../../../../lib/apiv2/engine/exec";
import type {
  AssetManager,
  Environment,
  ExecVariables,
} from "../../../../lib/engine/environment";
import type { Pipe } from "../../../../lib/engine/pipe";
import { ExecAM } from "../../../../lib/exec/exec_am";
import type { Registry } from "../../../../lib/registry/registry";
import { generateId } from "../../../../lib/utils";
import { EnvAssetManager } from "./asset_manager";

export interface EnvOptions {
  token: string;
  plug: string;
  agent: string;
  api_base_url: string;
  tenant_id: string;
  parent_secret?: string;
  pipe: Pipe;
  startup_payload?: any;
  registry: Registry<any>;
}

interface Pending {
  resolve: (value: unknown) => void;
  reject: (reason?: any) => void;
}

export class Env implements Environment {
  _opts: EnvOptions; // only for debug remove this
  _exec_api: ExecAPI;
  _startup_payload?: any;
  _registry: Registry<any>;

  _pipe: Pipe;
  _pending_pipe_msg: Map<string, Pending>;
  _default_parent_handler: (data: any) => {};

  constructor(opts: EnvOptions) {
    window["debug_env"] = this; // only for debug remove this

    this._registry = opts.registry;
    this._opts = opts;
    this._startup_payload = opts.startup_payload;
    this.set_up_pipe(opts.pipe);
    this._exec_api = new ExecAPI(
      opts.api_base_url.replace("v2/", "v2"),
      opts.token
    );
  }

  async init() {}

  // public

  PreformAction = async (name: string, data: any): Promise<any> => {
    return this.preformAction(name, data);
  };

  PreformParentAction = async (name: string, data: any): Promise<any> => {
    return this.preformParentAction(name, data);
  };

  OnParentAction = (handler: (data: any) => {}) => {
    this._default_parent_handler = handler;
  };

  GetRegistry = (): any => {
    return this._registry;
  };

  GetExecVars(): ExecVariables {
    return {
      agent_id: this._opts.agent,
      plug_id: this._opts.plug,
      api_base_url: this._opts.api_base_url,
      tenant_id: this._opts.tenant_id,
      exec_data: this._startup_payload,
    };
  }

  GetAssetManager(): AssetManager {
    return this.execApiManager();
  }

  GetExecApiManager(): Promise<any> {
    return Promise.resolve(new ExecAM(this._opts.api_base_url));
  }

  // private

  private set_up_pipe(pipe: Pipe) {
    this._pipe = pipe;
    this._pending_pipe_msg = new Map();

    this._pipe.set_handler((xid: string, action: string, data: any) => {
      const pending = this._pending_pipe_msg.get(xid);
      if (!pending) {
        if (this._default_parent_handler) {
          this._default_parent_handler(data);
        }
        return;
      }

      pending.resolve(data);
    });
  }

  private preformAction = (name: string, data: any) => {
    return this._exec_api.preform_action(name, data);
  };

  private preformParentAction = async (
    name: string,
    data: any
  ): Promise<any> => {
    const key = generateId();

    const p = new Promise((resolve, reject) => {
      this._pending_pipe_msg.set(key, {
        reject,
        resolve,
      });
    });

    this._pipe.send(key, name, data);
    return p;
  };

  private execApiManager = () => {
    return new EnvAssetManager(
      this._opts.api_base_url,
      this._opts.plug,
      this._opts.agent,
      "FIXME"
    );
  };
}
