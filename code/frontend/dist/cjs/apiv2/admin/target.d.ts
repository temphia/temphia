import type { ApiBase } from "../base";
export declare class AdminTargetAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    listApp(): Promise<import("../http").HttpResponse>;
    listAppByType(ttype: string, target?: string): Promise<import("../http").HttpResponse>;
    newApp(ttype: string, data: any): Promise<import("../http").HttpResponse>;
    getApp(ttype: string, id: number): Promise<import("../http").HttpResponse>;
    updateApp(ttype: string, id: number, data: any): Promise<import("../http").HttpResponse>;
    deleteApp(ttype: string, id: number): Promise<import("../http").HttpResponse>;
    listHook(): Promise<import("../http").HttpResponse>;
    listHookByType(ttype: string, target?: string): Promise<import("../http").HttpResponse>;
    newHook(ttype: string, data: any): Promise<import("../http").HttpResponse>;
    getHook(ttype: string, id: number): Promise<import("../http").HttpResponse>;
    updateHook(ttype: string, id: number, data: any): Promise<import("../http").HttpResponse>;
    deleteHook(ttype: string, id: number): Promise<import("../http").HttpResponse>;
}
