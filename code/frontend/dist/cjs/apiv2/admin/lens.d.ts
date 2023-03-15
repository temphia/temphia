import type { ApiBase } from "../base";
export declare class AdminLensAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    query(qopts: any): Promise<import("../http").HttpResponse>;
}
