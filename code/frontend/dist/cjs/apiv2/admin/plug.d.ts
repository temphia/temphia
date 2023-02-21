import type { ApiBase } from "../base";
export declare class AdminPlugAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list_plug(): Promise<import("../http").HttpResponse>;
    new_plug(data: any): Promise<import("../http").HttpResponse>;
    get_plug(pid: string): Promise<import("../http").HttpResponse>;
    update_plug(pid: string, data: any): Promise<import("../http").HttpResponse>;
    delete_plug(pid: string): Promise<import("../http").HttpResponse>;
    flowmap(pid: string): Promise<import("../http").HttpResponse>;
    list_agent(pid: string): Promise<import("../http").HttpResponse>;
    new_agent(pid: string, data: any): Promise<import("../http").HttpResponse>;
    get_agent(pid: string, aid: string): Promise<import("../http").HttpResponse>;
    update_agent(pid: string, aid: string, data: any): Promise<import("../http").HttpResponse>;
    delete_agent(pid: string, aid: string): Promise<import("../http").HttpResponse>;
    list_agent_link(pid: string, aid: string): Promise<import("../http").HttpResponse>;
    new_agent_link(pid: string, aid: string, data: any): Promise<import("../http").HttpResponse>;
    update_agent_link(pid: string, aid: string, lid: string, data: any): Promise<import("../http").HttpResponse>;
    get_agent_link(pid: string, aid: string, lid: string): Promise<import("../http").HttpResponse>;
    delete_agent_link(pid: string, aid: string, lid: string): Promise<import("../http").HttpResponse>;
    list_agent_ext(pid: string, aid: string): Promise<import("../http").HttpResponse>;
    new_agent_ext(pid: string, aid: string, data: any): Promise<import("../http").HttpResponse>;
    update_agent_ext(pid: string, aid: string, eid: string, data: any): Promise<import("../http").HttpResponse>;
    get_agent_ext(pid: string, aid: string, eid: string): Promise<import("../http").HttpResponse>;
    delete_agent_ext(pid: string, aid: string, eid: string): Promise<import("../http").HttpResponse>;
    list_agent_resource(pid: string, aid: string): Promise<import("../http").HttpResponse>;
    new_agent_resource(pid: string, aid: string, data: any): Promise<import("../http").HttpResponse>;
    update_agent_resource(pid: string, aid: string, rid: string, data: any): Promise<import("../http").HttpResponse>;
    get_agent_resource(pid: string, aid: string, rid: string): Promise<import("../http").HttpResponse>;
    delete_agent_resource(pid: string, aid: string, rid: string): Promise<import("../http").HttpResponse>;
    list_plug_state(pid: string, qparms: {
        key_cursor?: string;
        page?: number;
    }): Promise<import("../http").HttpResponse>;
    new_plug_state(pid: string, data: any): Promise<import("../http").HttpResponse>;
    get_plug_state(pid: string, skey: string): Promise<import("../http").HttpResponse>;
    update_plug_state(pid: string, skey: string, data: any): Promise<import("../http").HttpResponse>;
    delete_plug_state(pid: string, skey: string): Promise<import("../http").HttpResponse>;
}
