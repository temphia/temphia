import type { ApiBase } from "../base";
export declare class AdminUserGroupAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(): Promise<import("../http").HttpResponse>;
    new(data: any): Promise<import("../http").HttpResponse>;
    get(ugroup: string): Promise<import("../http").HttpResponse>;
    update(ugroup: string, data: any): Promise<import("../http").HttpResponse>;
    delete(ugroup: string): Promise<import("../http").HttpResponse>;
    listData(ugroup: string): Promise<import("../http").HttpResponse>;
    newData(ugroup: string, data: any): Promise<import("../http").HttpResponse>;
    getData(ugroup: string, id: string): Promise<import("../http").HttpResponse>;
    updateData(ugroup: string, id: string, data: any): Promise<import("../http").HttpResponse>;
    deleteData(ugroup: string, id: string): Promise<import("../http").HttpResponse>;
    listAuth(ugroup: string): Promise<import("../http").HttpResponse>;
    newAuth(ugroup: string, data: any): Promise<import("../http").HttpResponse>;
    getAuth(ugroup: string, id: string): Promise<import("../http").HttpResponse>;
    updateAuth(ugroup: string, id: string, data: any): Promise<import("../http").HttpResponse>;
    deleteAuth(ugroup: string, id: string): Promise<import("../http").HttpResponse>;
}
