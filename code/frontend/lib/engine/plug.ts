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

export interface FactoryOptions {
  plug: string;
  agent: string;
  entry: string;
  env: any; //Environment;
  target: HTMLElement;
  payload?: any;
  registry: any;
}

export interface PipeMessage {
  action?: string;
  xid: string;
  data: any;
  parent_secret?: string;
}
