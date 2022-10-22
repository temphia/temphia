import { Sockd } from "../../../sockd/sockd";
import type { ActionResponse, Environment, Pipe } from "../ecore";
import { actionFetch } from "./fetch";
import { FolderTktAPI } from "../../../apiv2";


export interface EnvOptions {
  token: string;
  plug: string;
  agent: string;
  base_url: string;
  parent_secret?: string;
  pipe: Pipe;
  startup_payload?: any;
}

interface Pending {
  resolve: (value: unknown) => void;
  reject: (reason?: any) => void;
}

export class Env implements Environment {
  _opts: EnvOptions; // only for debug remove this
  _fetch: (name: string, data: string) => Promise<Response>;

  _sockd_url: string

  _startup_payload?: any;
  _pipe: Pipe;
  _pending_pipe_msg: Map<string, Pending>;
  _default_parent_handler: (data: any) => {};

  constructor(opts: EnvOptions) {
    window["debug_env"] = this; // only for debug remove this

    this._opts = opts;
    this._pending_pipe_msg = new Map();

    this._pipe = opts.pipe;
    this._startup_payload = opts.startup_payload;
    this._fetch = actionFetch(
      `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_con`,
      opts.token
    );

    this._sockd_url = `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_ws`;
   
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

  init = async () => {
    
  };

  PreformAction = async (name: string, data: any): Promise<ActionResponse> => {
    const encoded = JSON.stringify(data);
    try {
      const resp = await this._fetch(name, encoded);
      const ctype = resp.headers.get("Content-Type");

      if (resp.status !== 200) {
        const txt = await resp.text();
        return {
          status_ok: false,
          content_type: ctype,
          body: txt,
        };
      }

      const respData = await resp.json();
      return {
        body: respData,
        content_type: ctype,
        status_ok: true,
      };
    } catch (error) {
      return {
        status_ok: false,
        body: error,
      };
    }
  };

  startup_payload = () => {
    return this._startup_payload;
  };

  PreformParentAction = async (name: string, data: any): Promise<any> => {
    const key = "fixme => generate";

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

  GetRegistry = (): any => {};

  GetFolderTktAPI = (ticket: string): any => {
    return new FolderTktAPI(this._opts.base_url, ticket);
  };

  GetRoomTktAPI = async (room: string, ticket?: string): Promise<any> => {
    const sroom  = new Sockd( {
      OnHandler: null,
      URL: `${this._sockd_url}?ticket=${ticket}`,
    });
    return sroom
  };

  GetDtableTktAPI = (ticket: string): any => {};
}
