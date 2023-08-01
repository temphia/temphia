const MESSAGE_SERVER_DIRECT = "server_direct";
const MESSAGE_SERVER_BROADCAST = "server_broadcast";
const MESSAGE_SERVER_PUBLISH = "server_publish";
const MESSAGE_CLIENT_DIRECT = "client_direct";
const MESSAGE_CLIENT_BROADCAST = "client_broadcast";
const MESSAGE_CLIENT_PUBLISH = "client_publish";
const MESSAGE_CLIENT_SYSTEM = "client_system";
const MESSAGE_SERVER_SYSTEM = "server_system";

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

export {
  MESSAGE_SERVER_DIRECT,
  MESSAGE_SERVER_BROADCAST,
  MESSAGE_SERVER_PUBLISH,
  MESSAGE_CLIENT_DIRECT,
  MESSAGE_CLIENT_BROADCAST,
  MESSAGE_CLIENT_PUBLISH,
  MESSAGE_CLIENT_SYSTEM,
  MESSAGE_SERVER_SYSTEM,
};
