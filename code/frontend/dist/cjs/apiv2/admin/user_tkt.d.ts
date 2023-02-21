import type { ApiBase } from "../base";
export declare class AdminUserTktAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(): Promise<import("../http").HttpResponse>;
    get(uid: string): Promise<import("../http").HttpResponse>;
    new(data: any): Promise<import("../http").HttpResponse>;
    update(uid: string, data: any): Promise<import("../http").HttpResponse>;
    delete(uid: string): Promise<import("../http").HttpResponse>;
}
