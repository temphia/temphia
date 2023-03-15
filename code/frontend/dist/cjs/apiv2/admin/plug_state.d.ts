import { Http } from "../http";
export declare class AdminPlugStateTktAPI {
    http: Http;
    token: string;
    api_base_url: string;
    constructor(api_base_url: string, token: any);
    query(options: any): Promise<import("../http").HttpResponse>;
    add(key: string, value: string, opts?: any): Promise<import("../http").HttpResponse>;
    update(key: string, value: string, opts?: any): Promise<import("../http").HttpResponse>;
    delete(key: string): Promise<import("../http").HttpResponse>;
    get(key: string): Promise<import("../http").HttpResponse>;
}
