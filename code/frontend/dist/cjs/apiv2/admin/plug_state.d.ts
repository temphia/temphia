import type { ApiBase } from "../base";
export declare class AdminPlugStateTktAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(qparms: {
        key_cursor?: string;
        page?: number;
        page_count?: number;
    }): Promise<import("../http").HttpResponse>;
    add(key: string, value: string): Promise<import("../http").HttpResponse>;
    update(key: string, value: string): Promise<import("../http").HttpResponse>;
    delete(key: string): Promise<import("../http").HttpResponse>;
    get(key: string): Promise<import("../http").HttpResponse>;
}
