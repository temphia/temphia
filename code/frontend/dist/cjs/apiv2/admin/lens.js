"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminLensAPI = void 0;
class AdminLensAPI {
    constructor(base) {
        this.base = base;
    }
    query(qopts) {
        return this.base.post(`/admin/lens/query`, qopts);
    }
}
exports.AdminLensAPI = AdminLensAPI;
