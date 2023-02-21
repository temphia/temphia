import { Http } from "./http";
export declare class AuthAPI {
    http: Http;
    constructor(baseUrl: string, site_token: string);
    list_methods: (ugroup?: string) => Promise<import("./http").HttpResponse>;
    login_next: (data: any) => Promise<import("./http").HttpResponse>;
    login_submit: (data: any) => Promise<import("./http").HttpResponse>;
    altauth_generate: (id: number, data: any) => Promise<import("./http").HttpResponse>;
    altauth_next: (id: number, stage: string, data: any) => Promise<import("./http").HttpResponse>;
    altauth_submit: (id: number, data: any) => Promise<import("./http").HttpResponse>;
    finish: (data: any) => Promise<import("./http").HttpResponse>;
    signup_next: (data: any) => Promise<import("./http").HttpResponse>;
    signup_submit: (data: any) => Promise<import("./http").HttpResponse>;
    reset_submit: (data: any) => Promise<import("./http").HttpResponse>;
    reset_finish: (data: any) => Promise<import("./http").HttpResponse>;
    about: (user_token: string) => Promise<import("./http").HttpResponse>;
}
