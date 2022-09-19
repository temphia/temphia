import type { SockdMessage, SockdHandler, ISockd } from "./stypes";
import {
  LinearBackoff,
  LRUBuffer,
  Websocket,
  WebsocketBuilder,
} from "../../vendor/ws";

export class Sockd implements ISockd {
  _ws: Websocket;
  _handler: (message: SockdMessage) => void;
  _builder: WebsocketBuilder;

  constructor(url: string) {
    console.log("CONNECTING WS @ ", url);
    this._builder = new WebsocketBuilder(url);
    this._builder.onMessage(this.handleIncoming);
    this._builder.withBackoff(new LinearBackoff(1, 3));
    this._builder.withBuffer(new LRUBuffer(20));
  }

  init = async () => {
    this._ws = this._builder.build();
  };

  private handleIncoming = (_: Websocket, ev: MessageEvent) => {
    // fixme => handle system messages
    const data = JSON.parse(ev.data);
    this._handler(data);
  };

  OnSockdMessage = (h: SockdHandler): void => {
    this._handler = h;
  };

  SendSockd = (message: SockdMessage): void => {
    this._ws.send(JSON.stringify(message));
  };
}
