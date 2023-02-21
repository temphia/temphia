"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.RepoAPI = void 0;
// fixme => change this to StoreAPI
class RepoAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/repo/");
    }
    load(id) {
        return this.base.get(`/repo/${id}`);
    }
    getBprint(id, group, slug) {
        return this.base.get(`/repo/${id}/${group}/${slug}`);
    }
    getBprintFile(id, group, slug, file) {
        return this.base.get(`/repo/${id}/${group}/${slug}/${file}`);
    }
}
exports.RepoAPI = RepoAPI;
