"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminLensAPI = void 0;
class AdminLensAPI {
    constructor(base) {
        this.base = base;
    }
    query(index, qopts) {
        return this.base.post(`/admin/lens/${index}`, qopts);
    }
}
exports.AdminLensAPI = AdminLensAPI;
