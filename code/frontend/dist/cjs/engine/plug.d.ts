export interface LoaderOptions {
    token: string;
    entry: string;
    exec_loader: string;
    plug: string;
    agent: string;
    api_base_url: string;
    parent_secret?: string;
    startup_payload: any;
    tenant_id: string;
}
export interface FactoryOptions {
    plug: string;
    agent: string;
    entry: string;
    env: any;
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
