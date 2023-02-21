"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminPlugAPI = void 0;
class AdminPlugAPI {
    constructor(base) {
        this.base = base;
    }
    list_plug() {
        return this.base.get("/admin/plug/");
    }
    new_plug(data) {
        return this.base.post("/admin/plug/", data);
    }
    get_plug(pid) {
        return this.base.get(`/admin/plug/${pid}`);
    }
    update_plug(pid, data) {
        return this.base.post(`/admin/plug/${pid}`, data);
    }
    delete_plug(pid) {
        return this.base.delete(`/admin/plug/${pid}`);
    }
    flowmap(pid) {
        return this.base.get(`/admin/plug/${pid}/flowmap`);
    }
    // agent
    list_agent(pid) {
        return this.base.get(`/admin/plug/${pid}/agent`);
    }
    new_agent(pid, data) {
        return this.base.post(`/admin/plug/${pid}/agent`, data);
    }
    get_agent(pid, aid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}`);
    }
    update_agent(pid, aid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}`, data);
    }
    delete_agent(pid, aid) {
        return this.base.delete(`/admin/plug/${pid}/agent/${aid}`);
    }
    // link
    list_agent_link(pid, aid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/link`);
    }
    new_agent_link(pid, aid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/link`, data);
    }
    update_agent_link(pid, aid, lid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/link/${lid}`, data);
    }
    get_agent_link(pid, aid, lid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/link/${lid}`);
    }
    delete_agent_link(pid, aid, lid) {
        return this.base.delete(`/admin/plug/${pid}/agent/${aid}/link/${lid}`);
    }
    // extension
    list_agent_ext(pid, aid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/extension`);
    }
    new_agent_ext(pid, aid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/extension`, data);
    }
    update_agent_ext(pid, aid, eid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/extension/${eid}`, data);
    }
    get_agent_ext(pid, aid, eid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/extension/${eid}`);
    }
    delete_agent_ext(pid, aid, eid) {
        return this.base.delete(`/admin/plug/${pid}/agent/${aid}/extension/${eid}`);
    }
    // resource
    list_agent_resource(pid, aid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/resource`);
    }
    new_agent_resource(pid, aid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/resource`, data);
    }
    update_agent_resource(pid, aid, rid, data) {
        return this.base.post(`/admin/plug/${pid}/agent/${aid}/resource/${rid}`, data);
    }
    get_agent_resource(pid, aid, rid) {
        return this.base.get(`/admin/plug/${pid}/agent/${aid}/resource/${rid}`);
    }
    delete_agent_resource(pid, aid, rid) {
        return this.base.delete(`/admin/plug/${pid}/agent/${aid}/resource/${rid}`);
    }
    // state
    list_plug_state(pid, qparms) {
        const u = new URLSearchParams();
        u.set("page", String(qparms.page || 0));
        u.set("key_cursor", qparms.key_cursor || "");
        return this.base.get(`/admin/plug/${pid}/state/?${u.toString()}`);
    }
    new_plug_state(pid, data) {
        return this.base.post(`/admin/plug/${pid}/state/`, data);
    }
    get_plug_state(pid, skey) {
        return this.base.get(`/admin/plug/${pid}/state/${skey}`);
    }
    update_plug_state(pid, skey, data) {
        return this.base.post(`/admin/plug/${pid}/state/${skey}`, data);
    }
    delete_plug_state(pid, skey) {
        return this.base.delete(`/admin/plug/${pid}/state/${skey}`);
    }
}
exports.AdminPlugAPI = AdminPlugAPI;
