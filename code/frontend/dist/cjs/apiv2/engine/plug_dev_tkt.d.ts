import { Http } from "../http";
export declare class PlugDevTktAPI {
    http: Http;
    token: string;
    base_url: string;
    constructor(base_url: string, token: string);
    bprint_list_file(): Promise<import("../http").HttpResponse>;
    bprint_push_file(data: FormData): Promise<Response>;
    bprint_get_file(file: string): Promise<import("../http").HttpResponse>;
    bprint_del_file(file: string): Promise<import("../http").HttpResponse>;
    exec_watch_agents_url(pid: string, aid: string): string;
    exec_reset_plug(pid: string, aid: string, data: any): Promise<import("../http").HttpResponse>;
    exec_run_agent_action(pid: string, aid: string, action: string, data: any): Promise<import("../http").HttpResponse>;
}
