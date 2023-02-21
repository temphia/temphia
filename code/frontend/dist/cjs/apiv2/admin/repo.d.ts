import type { ApiBase } from "../base";
export declare class AdminRepoAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(): Promise<import("../http").HttpResponse>;
    new(data: any): Promise<import("../http").HttpResponse>;
    get(rid: string): Promise<import("../http").HttpResponse>;
    update(rid: string, data: any): Promise<import("../http").HttpResponse>;
    delete(rid: string): Promise<import("../http").HttpResponse>;
}
