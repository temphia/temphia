export interface LoaderOptions {
    token: string;
    entry: string;
    exec_loader: string;
    plug: string;
    agent: string;
    api_base_url: string;
    parent_secret?: string;
    startup_payload?: any;
    tenant_id: string
}