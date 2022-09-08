import type { PipeMessage } from "../../core/engine";
import type { ExecInstance } from "./etype";

export type Commands = {
  [_: string]: (exec_id: string,  data: any) => void;
};

export class EngineService {
  _instances: Map<string, ExecInstance>;
  _commands: Commands;

  constructor(commands: Commands) {
    this._instances = new Map();
    this._commands = commands;
    window.addEventListener("message", this.on_message);
  }

  on_message = (ev) => {
    try {
      const decoded: PipeMessage = JSON.parse(ev.data);
      console.log("EVENT@parent", decoded)
      const exec = this._instances.get(decoded["parent_secret"]);


      const chandler = this._commands[decoded.action];
      if (chandler) {
        chandler(decoded.xid, decoded.data);
        return;
      }
      exec.handle(decoded.xid, decoded.action, decoded.data);
    } catch (error) {
      console.log("engine interframe communication error", error);
    }
  };

  get(exec_id: string): ExecInstance {
    return this._instances.get(exec_id);
  }

  set(exec_id: string, instance: ExecInstance): void {
    this._instances.set(exec_id, instance);
  }

  del = (exec_id: string) => {
    this._instances.delete(exec_id);
  };
}
