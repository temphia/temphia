import type { ApiBase } from "../base";
export declare class AdminCheckAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    bprint(bid: string): void;
    plug(pid: string): void;
    dataGroup(gid: string): void;
}
