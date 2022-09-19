export const MODE_IFRAME = "IFRAME";
export const MODE_RAW_DOM = "RAW_DOM";
export const MODE_SUB_ORIGIN = "SUB_ORIGIN";

export interface LoaderOptions {
  token: string;
  entry: string;
  exec_loader: string;
  plug: string;
  agent: string;
  base_url: string;
  parent_secret?: string;
  startup_payload?: any;
}

export interface ActionResponse {
  status_ok: boolean;
  content_type?: string;
  body: any;
}

export interface Environment {
  PreformAction: (name: string, data: any) => Promise<ActionResponse>;
  PreformParentAction: (name: string, data: any) => Promise<any>;
  OnParentAction: (handler: (data: any) => {}) => void;

  GetRegistry: () => any;
  GetFolderTktAPI: (ticket: string) => any;
  GetRoomTktAPI: (room: string, ticket?: string) => Promise<any>;
  GetDtableTktAPI: (ticket: string) => any;
}

export interface FactoryOptions {
  plug: string;
  agent: string;
  entry: string;
  env: Environment;
  target: HTMLElement;
  payload?: any;
  registry: any;
}

export type Factory = (opts: FactoryOptions) => void;

// pipe stuff

export interface PipeMessage {
  action?: string;
  xid: string;
  data: any;
  parent_secret?: string;
}

export type PipeHandler = (xid: string, action: string, data: any) => void;

export interface Pipe {
  send(xid: string, action: string, data: any): void;
  set_handler(fn: PipeHandler): void;
  remove_handler(fn: PipeHandler): void;
}
