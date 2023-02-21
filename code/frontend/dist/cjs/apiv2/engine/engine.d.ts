import type { ApiBase } from "../base";
export declare class EngineAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    launch_target(data: any): Promise<import("../http").HttpResponse>;
    launch_admin(data: any): Promise<import("../http").HttpResponse>;
}
