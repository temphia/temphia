import type { ApiBase } from "../base";
export declare class AdminPlugStateTktAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    query(options: any): Promise<import("../http").HttpResponse>;
    add(key: string, value: string, opts?: any): Promise<import("../http").HttpResponse>;
    update(key: string, value: string, opts?: any): Promise<import("../http").HttpResponse>;
    delete(key: string): Promise<import("../http").HttpResponse>;
    get(key: string): Promise<import("../http").HttpResponse>;
}
