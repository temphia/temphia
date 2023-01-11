import {
  SockdMessage,
  SockdHandler,
  ISockd,
  MESSAGE_SERVER_SYSTEM,
  MESSAGE_CLIENT_BROADCAST,
  MESSAGE_CLIENT_DIRECT,
  MESSAGE_CLIENT_PUBLISH,
  MESSAGE_CLIENT_SYSTEM,
} from "./stypes";

import {
  LinearBackoff,
  LRUBuffer,
  Websocket,
  WebsocketBuilder,
} from "../vendor/ws";
import { generateId } from "../utils";

export interface Options {
  URL: string;
  OnHandler: (message: SockdMessage) => void;
}

export class Sockd implements ISockd {
  _ws: Websocket;
  _handler: (message: SockdMessage) => void;
  _builder: WebsocketBuilder;
  _room: string;
  _sid: number;

  constructor(opts: Options) {
    this._builder = new WebsocketBuilder(opts.URL);
    this._builder.onMessage(this.handleIncoming);
    this._builder.withBackoff(new LinearBackoff(0, 10, 100));
    this._builder.withBuffer(new LRUBuffer(20));
    this._handler = opts.OnHandler;
  }

  init = async () => {
    this._ws = this._builder.build();
  };

  private handleIncoming = (_: Websocket, ev: MessageEvent) => {
    const data: SockdMessage = JSON.parse(ev.data);
    if (data.type === MESSAGE_SERVER_SYSTEM) {
      // fixme handle this
      console.log("@data", data);
      return;
    }
    this._handler(data);
  };

  SendDirect = (data: any, target: number): void => {
    this._ws.send(
      JSON.stringify({
        type: MESSAGE_CLIENT_DIRECT,
        xid: generateId(),
        room: this._room,
        from_id: this._sid,
        targets: [target],
        payload: data,
      })
    );
  };

  SendBroadcast = (data: any): void => {
    this._ws.send(
      JSON.stringify({
        type: MESSAGE_CLIENT_BROADCAST,
        xid: generateId(),
        room: this._room,
        from_id: this._sid,
        payload: data,
      })
    );
  };

  SendTagged = (data: any, targets?: string[]): void => {
    this._ws.send(
      JSON.stringify({
        type: MESSAGE_CLIENT_PUBLISH,
        xid: generateId(),
        room: this._room,
        from_id: this._sid,
        payload: data,
        target_tags: targets,
      })
    );
  };
  UpdateToken = (token: string): void => {
    this._ws.send(
      JSON.stringify({
        type: MESSAGE_CLIENT_SYSTEM,
        xid: generateId(),
        room: this._room,
        from_id: this._sid,
        payload: token,
      })
    );
  };
  Close = (): void => {
    this._ws.close(0, "closed by client");
  };
}
