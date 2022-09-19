const MESSAGE_SERVER_DIRECT = "server_direct";
const MESSAGE_SERVER_BROADCAST = "server_broadcast";
const MESSAGE_SERVER_PUBLISH = "server_publish";
const MESSAGE_PEER_DIRECT = "peer_direct";
const MESSAGE_PEER_BROADCAST = "peer_broadcast";
const MESSAGE_PEER_PUBLISH = "peer_publish";

export type SockdHandler = (message: SockdMessage) => void;

export interface SockdMessage {
  type: string;
  xid: string;
  room?: string;
  from_id?: string;
  server_ident?: string;
  ticket?: string;
  targets?: string[];
  payload: any;
}

export interface ISockd {
  init(): Promise<void>;
  OnSockdMessage(h: SockdHandler): void;
  SendSockd(message: SockdMessage): void;
}

export interface ISockdRoom {
  SendDirect(data: any): void;
  SendBroadcast(data: any): void;
  SendTagged(data: any, ticket: string, targets?: string[]): void;
  onMessage(handler: SockdHandler): void;
  onPeer(handler: SockdHandler): void;
  onServer(handler: SockdHandler): void;
  ProcessMessage(message: SockdMessage): void;
  IsConnected(): Promise<boolean>;
  LeaveRoom(): void;
}

export {
  MESSAGE_SERVER_DIRECT,
  MESSAGE_SERVER_BROADCAST,
  MESSAGE_SERVER_PUBLISH,
  MESSAGE_PEER_DIRECT,
  MESSAGE_PEER_BROADCAST,
  MESSAGE_PEER_PUBLISH,
};
