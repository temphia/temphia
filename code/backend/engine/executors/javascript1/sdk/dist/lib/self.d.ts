export declare const self: {
    list_resource: () => [any, string];
    get_resource: (name: string) => [any, string];
    inlinks: () => [any, string];
    outlinks: () => [any, string];
    new_module: (name: string, data: any) => [number, string];
    module_ticket: (name: string, opts: any) => [number, string];
    module_execute: (mid: number, method: string, data: any) => [any, string];
    link_execute: (name: string, method: string, path: string, data: any, async: boolean, detached: boolean) => [any, string];
    fork_execute: (method: string, data: any) => string;
};
