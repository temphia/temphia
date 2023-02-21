"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminUserAPI = void 0;
class AdminUserAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/user/");
    }
    get(uid) {
        return this.base.get(`/admin/user/${uid}`);
    }
    new(data) {
        return this.base.post(`/admin/user/`, data);
    }
    update(uid, data) {
        return this.base.post(`/admin/user/${uid}`, data);
    }
    delete(uid) {
        return this.base.delete(`/admin/user/${uid}`);
    }
}
exports.AdminUserAPI = AdminUserAPI;
/*

[GIN-debug] GET    /z/api/:tenant_id/v2/admin/user/perm --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/user/perm --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/user/perm/:perm --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/user/perm/:perm --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/admin/user/perm/:perm --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)

[GIN-debug] GET    /z/api/:tenant_id/v2/admin/user/role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/user/role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] GET    /z/api/:tenant_id/v2/admin/user/role/:role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/user/role/:role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/admin/user/role/:role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)

[GIN-debug] GET    /z/api/:tenant_id/v2/admin/user/user_role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] POST   /z/api/:tenant_id/v2/admin/user/user_role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)
[GIN-debug] DELETE /z/api/:tenant_id/v2/admin/user/user_role --> github.com/temphia/temphia/code/backend/app/server/middleware.(*Middleware).Authed.func1 (3 handlers)


*/
