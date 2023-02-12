import type {
  Pipe,
  PipeHandler,
  PipeMessage,
} from "../../../../lib/engine/pipe";

export class IFramePipe implements Pipe {
  secret: string;
  handlers: Set<PipeHandler>;
  port: MessagePort;

  constructor(secret: string, port: MessagePort) {
    this.secret = secret;
    this.handlers = new Set();
    this.port = port;

    window.addEventListener("message", (ev) => {
      try {
        const decoded: PipeMessage = JSON.parse(ev.data);
        this.handlers.forEach((fn) =>
          fn(decoded.xid, decoded.action, decoded.data)
        );
      } catch (error) {}
    });
  }

  set_handler = (fn: PipeHandler) => {
    this.handlers.add(fn);
  };

  remove_handler = (fn: PipeHandler) => {
    this.handlers.delete(fn);
  };

  send = (xid: string, action: string, data: any) => {
    const message = JSON.stringify({
      xid,
      data,
      action,
      parent_secret: this.secret,
    });

    this.port.postMessage(message);
  };
}
