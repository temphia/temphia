import type { ApiBase } from "../base";
export declare class AdminTenantAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    edit(data: any): Promise<import("../http").HttpResponse>;
    get(): Promise<import("../http").HttpResponse>;
    get_domains(): Promise<import("../http").HttpResponse>;
    new_domain(data: any): Promise<import("../http").HttpResponse>;
    get_domain(did: string): Promise<import("../http").HttpResponse>;
    edit_domain(did: string, data: any): Promise<import("../http").HttpResponse>;
    delete_domain(did: string): Promise<import("../http").HttpResponse>;
    domain_issue_adapter_editor(did: string): Promise<import("../http").HttpResponse>;
    domain_adapter_reset(did: string): Promise<import("../http").HttpResponse>;
    list_system_kv({ last, etype, prefix }: {
        last: any;
        etype: any;
        prefix: any;
    }): Promise<import("../http").HttpResponse>;
    list_system_event({ last, etype }: {
        last: any;
        etype: any;
    }): Promise<import("../http").HttpResponse>;
}
