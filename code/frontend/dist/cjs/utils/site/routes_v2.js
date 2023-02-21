"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.authURL = exports.portalURL = exports.baseURL = exports.apiURL = void 0;
//http://localhost:4000/z/api/:tenant_id/v2
exports.apiURL = (tenant_id) => `${window.location.origin}/z/api/${tenant_id}/v2`;
//http://localhost:4000
exports.baseURL = () => window.location.origin;
exports.portalURL = () => `${window.location.origin}/z/portal`;
exports.authURL = (opts) => {
    if (!opts) {
        return `${window.location.origin}/z/auth`;
    }
    return `${window.location.origin}/z/auth?${opts.tenant_id ? "tenant_id=" + opts.tenant_id + "&" : ""}${opts.user_group ? "ugroup=" + opts.user_group : ""}`;
};
