import { SockdMessage } from "./stypes";
import { Websocket, WebsocketBuilder } from "../vendor/ws";
export declare class Sockd {
    _ws: Websocket;
    _handler: (message: SockdMessage) => void;
    _builder: WebsocketBuilder;
    _room: string;
    _sid: number;
    constructor(url: string);
    Init(): Promise<void>;
    SetHandler(fn: (msg: SockdMessage) => void): void;
    SendDirect: (data: any, target: number) => void;
    SendBroadcast: (data: any) => void;
    SendTagged: (data: any, targets?: string[]) => void;
    UpdateToken: (token: string) => void;
    Close: () => void;
    private handleIncoming;
}
