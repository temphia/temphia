import type { ApiBase } from "./base";
export declare class RepoAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(): Promise<import("./http").HttpResponse>;
    load(id: string): Promise<import("./http").HttpResponse>;
    getBprint(id: string, group: string, slug: string): Promise<import("./http").HttpResponse>;
    getBprintFile(id: string, group: string, slug: string, file: string): Promise<import("./http").HttpResponse>;
}
