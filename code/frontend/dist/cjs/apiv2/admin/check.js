"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminCheckAPI = void 0;
class AdminCheckAPI {
    constructor(base) {
        this.base = base;
    }
    bprint(bid) {
        this.base.get(`check/bprint/${bid}`);
    }
    plug(pid) {
        this.base.get(`check/plug/${pid}`);
    }
    dataGroup(gid) {
        this.base.get(`check/dgroup/${gid}`);
    }
}
exports.AdminCheckAPI = AdminCheckAPI;
