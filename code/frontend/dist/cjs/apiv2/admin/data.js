"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminDataAPI = void 0;
class AdminDataAPI {
    constructor(base) {
        this.seed_table = (source, gid, tid, max) => {
            return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/seed?max=${max}`);
        };
        this.query = (source, gid, opts) => {
            return this.base.post(`/admin/data/${source}/group/${gid}/query`, opts);
        };
        this.list_table_activity = (source, gid, tid, offset) => {
            return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/activity?offset=${offset}`);
        };
        this.base = base;
    }
    list_group(source) {
        return this.base.get(`/admin/data/${source}/group`);
    }
    new_group(source, data) {
        return this.base.post(`/admin/data/${source}/group`, data);
    }
    get_group(source, gid) {
        return this.base.get(`/admin/data/${source}/group/${gid}`);
    }
    edit_group(source, gid, data) {
        return this.base.patch(`/admin/data/${source}/group/${gid}`, data);
    }
    delete_group(source, gid) {
        return this.base.delete(`/admin/data/${source}/group/${gid}`);
    }
    // table
    list_tables(source, gid) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table`);
    }
    get_table(source, gid, tid) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}`);
    }
    edit_table(source, gid, tid, data) {
        return this.base.patch(`/admin/data/${source}/group/${gid}/table/${tid}`, data);
    }
    delete_table(source, gid, tid) {
        return this.base.delete(`/admin/data/${source}/group/${gid}/table/${tid}`);
    }
    // column
    list_column(source, gid, tid) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/column`);
    }
    get_column(source, gid, tid, cid) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`);
    }
    edit_column(source, gid, tid, cid, data) {
        return this.base.patch(`/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`, data);
    }
    delete_column(source, gid, tid, cid) {
        return this.base.delete(`/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`);
    }
    // view
    list_view(source, gid, tid) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/view`);
    }
    add_view(source, gid, tid, data) {
        return this.base.post(`/admin/data/${source}/group/${gid}/table/${tid}/view`, data);
    }
    get_view(source, gid, tid, id) {
        return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`);
    }
    edit_view(source, gid, tid, id, data) {
        return this.base.post(`/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`, data);
    }
    delete_view(source, gid, tid, id) {
        return this.base.delete(`/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`);
    }
}
exports.AdminDataAPI = AdminDataAPI;
