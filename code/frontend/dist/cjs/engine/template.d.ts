interface iframeBuildOptions {
    api_base_url: string;
    entry_name: string;
    plug: string;
    agent: string;
    token: string;
    js_plug_script: string;
    exec_loader: string;
    style_file: string;
    ext_scripts?: object;
    parent_secret: string;
    startup_payload?: any;
    tenant_id: string;
    bootloader: string;
}
export declare const iframeTemplateBuild: (opts: iframeBuildOptions) => string;
export {};
