import { Http } from "../http";
export declare class ExecAPI {
    http: Http;
    exec_token: string;
    base_url: string;
    constructor(base_url: string, exec_token: string);
    agent_file_url(pid: string, aid: string, file: string): string;
    executor_file_url(eid: string, pid: string, aid: string, file: string): string;
    ws_url(room_token: string): Promise<import("../http").HttpResponse>;
    ws_update(room_token: string, data: any): void;
    preform_action(method: string, data: any): Promise<import("../http").HttpResponse>;
}
