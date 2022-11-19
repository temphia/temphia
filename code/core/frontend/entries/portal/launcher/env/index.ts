import { ExecAPI } from "../../../../lib/apiv2/engine/exec";
import type { Environment } from "../../../../lib/engine/env";
import type { Pipe } from "../../../../lib/engine/pipe";
import type { Registry } from "../../../../lib/registry/registry";
import { generateId } from "../../../../lib/utils";

export interface EnvOptions {
  token: string;
  plug: string;
  agent: string;
  base_url: string;
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
    this._exec_api = new ExecAPI(opts.base_url, opts.token);
  }

  set_up_pipe(pipe: Pipe) {
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

  init = async () => {};

  // public

  PreformAction = async (name: string, data: any): Promise<any> => {
    return null;
  };

  startup_payload = () => {
    return this._startup_payload;
  };

  PreformParentAction = async (name: string, data: any): Promise<any> => {
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

  OnParentAction = (handler: (data: any) => {}) => {
    this._default_parent_handler = handler;
  };

  GetRegistry = (): any => {
    return this._registry;
  };

  GetFolderTktAPI = (ticket: string): any => {
    //return new FolderTktAPI(this._opts.base_url, ticket);
    return null;
  };

  GetRoomTktAPI = async (room: string, ticket?: string): Promise<any> => {
    // const sroom  = new Sockd(`${this._sockd_url}?ticket=${ticket}`);
    // return sroom
    return null;
  };

  GetDataTableTktAPI = (ticket: string): any => {};
}
