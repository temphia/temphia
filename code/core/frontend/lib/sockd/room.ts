import type { SockdHandler, SockdMessage, ISockd, ISockdRoom } from "./stypes";
import {
  MESSAGE_SERVER_DIRECT,
  MESSAGE_SERVER_BROADCAST,
  MESSAGE_SERVER_PUBLISH,
  MESSAGE_PEER_DIRECT,
  MESSAGE_PEER_BROADCAST,
  MESSAGE_PEER_PUBLISH,
} from "./stypes";

export class SockdRoom implements ISockdRoom {
  _socket: ISockd;
  _room: string;

  _onMessage?: SockdHandler;
  _onPeer?: SockdHandler;
  _onServer?: SockdHandler;

  constructor(socket: ISockd, room: string) {
    this._socket = socket;
    this._room = room;
  }

  SendDirect = (data: any) => {
    this._socket.SendSockd({
      payload: data,
      type: MESSAGE_PEER_DIRECT,
      xid: "",
      from_id: "",
      room: this._room,
    });
  };

  SendBroadcast = (data: any) => {
    this._socket.SendSockd({
      payload: data,
      type: MESSAGE_PEER_BROADCAST,
      xid: "",
      from_id: "",
      room: this._room,
    });
  };

  SendTagged = (data: any, ticket: string, targets?: string[]) => {
    this._socket.SendSockd({
      payload: data,
      type: MESSAGE_PEER_PUBLISH,
      xid: "",
      from_id: "",
      room: this._room,
      targets: targets,
      ticket: ticket,
    });
  };

  onMessage = (handler: SockdHandler): void => {
    this._onMessage = handler;
  };
  onPeer = (handler: SockdHandler): void => {
    this._onPeer = handler;
  };

  onServer = (handler: SockdHandler): void => {
    this._onServer = handler;
  };

  ProcessMessage = (message: SockdMessage) => {
    if (this._onMessage) {
      this._onMessage(message);
    }
    switch (message.type) {
      case MESSAGE_SERVER_DIRECT:
        if (this._onServer) {
          this._onServer(message);
        }
      case MESSAGE_SERVER_BROADCAST:
        if (this._onServer) {
          this._onServer(message);
        }
      case MESSAGE_SERVER_PUBLISH:
        if (this._onServer) {
          this._onServer(message);
        }
      case MESSAGE_PEER_DIRECT:
        if (this._onPeer) {
          this._onPeer(message);
        }
      case MESSAGE_PEER_BROADCAST:
        if (this._onPeer) {
          this._onPeer(message);
        }
      case MESSAGE_PEER_PUBLISH:
        if (this._onPeer) {
          this._onPeer(message);
        }
      default:
        break;
    }
  };

  IsConnected = async (): Promise<boolean> => {
    return false;
  };

  LeaveRoom = () => {
    // fixme => impl
  };
}
