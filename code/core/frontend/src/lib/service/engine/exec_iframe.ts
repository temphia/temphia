import { iframeTemplateBuild } from "../../core/engine/execbase";
import type { ExecInstance } from "./etype";

export interface IframeInstanceOptions {
  plug: string;
  agent: string;
  secret: string;
//  exec_type: string;
  engine_data: object;
  exec_source?: object;
  handler: (exec_id: string, action: string, data: any) => void;
  target: HTMLElement;
}

export class IframeExec implements ExecInstance {
  exec_id: string;
  _plug: string;
  _agent: string;
  _secret: string;
  _engine_data: object;

  _itarget: HTMLIFrameElement;
  _target: HTMLElement;

  _handler: (exec_id: string, action: string, data: any) => void;

  constructor(opts: IframeInstanceOptions) {
    this._plug = opts.plug;
    this._agent = opts.agent;
    this._secret = opts.secret;
    this._engine_data = opts.engine_data;
    this._handler = opts.handler;
    this._target = opts.target;
  }
  handle(exec_id: string, action: string, data: any): void {}
  send(action: string, data: any): void {
    const _data = JSON.stringify(data); // fixme => interface data { }
    this._itarget.contentWindow.postMessage(_data, "*");
  }

  run = (launch_data: object) => {
    this._itarget = document.createElement("iframe");
    this._target.appendChild(this._itarget);

    const src = iframeTemplateBuild({
      agent: this._agent,
      plug: this._plug,
      base_url: this._engine_data["base_url"],
      entry_name: this._engine_data["entry"],
      exec_loader: this._engine_data["exec_loader"],
      js_plug_script: this._engine_data["js_plug_script"],
      style_file: this._engine_data["style"],
      token: this._engine_data["token"] || "",
      ext_scripts: this._engine_data["ext_scripts"],
      parent_secret: this._secret,
      startup_payload: launch_data
    });

    this._itarget.setAttribute("srcdoc", src);
    this._itarget.style.height = "100%";
    this._itarget.style.width = "100%";
  };

  close(): void {
    if (this._itarget) {
      this._itarget.remove();
    }
    // this.parent.clear_exec(this.secret);
    // this.message_handler = null;
    // this.parent = null;
  }
}
