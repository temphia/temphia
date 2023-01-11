var __dirname = ''; var module = {}; module['exports']={};/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	// The require scope
/******/ 	var __nccwpck_require__ = {};
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__nccwpck_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/compat */
/******/ 	
/******/ 	if (typeof __nccwpck_require__ !== 'undefined') __nccwpck_require__.ab = __dirname + "/";
/******/ 	
/************************************************************************/
var __webpack_exports__ = {};
// ESM COMPAT FLAG
__nccwpck_require__.r(__webpack_exports__);

;// CONCATENATED MODULE: ./lib/adapter/adapter.ts
class AdapterEditorEnv {
    constructor(api, domain_name) {
        this.api = api;
        this.domain_name = domain_name;
    }
}

;// CONCATENATED MODULE: ./lib/utils/site/routes_v2.ts
//http://localhost:4000/z/api/:tenant_id/v2
const apiURL = (tenant_id) => `${window.location.origin}/z/api/${tenant_id}/v2`;
//http://localhost:4000
const baseURL = () => window.location.origin;
const portalURL = () => `${window.location.origin}/z/portal`;
const authURL = (opts) => {
    if (!opts) {
        return `${window.location.origin}/z/auth`;
    }
    return `${window.location.origin}/z/auth?${opts.tenant_id ? "tenant_id=" + opts.tenant_id + "&" : ""}${opts.user_group ? "ugroup=" + opts.user_group : ""}`;
};

;// CONCATENATED MODULE: ./lib/apiv2/http/http.ts
class Http {
    constructor(baseURL, headers) {
        this.baseURL = baseURL;
        this.headers = headers;
    }
    replace_headers(headers) {
        this.headers = headers;
    }
    async get(path) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: "GET",
            headers: this.headers,
        });
        if (resp.ok) {
            const text = await resp.text();
            try {
                const data = JSON.parse(text);
                return Promise.resolve({
                    ok: true,
                    data,
                    status: resp.status,
                });
            }
            catch (error) {
                return Promise.resolve({
                    ok: true,
                    data: text,
                    status: resp.status,
                });
            }
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
    async post(path, data) {
        return this.jsonMethod(path, "POST", data);
    }
    async patch(path, data) {
        return this.jsonMethod(path, "PATCH", data);
    }
    async put(path, data) {
        return this.jsonMethod(path, "PUT", data);
    }
    async jsonMethod(path, method, data) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: method,
            headers: this.headers,
            body: JSON.stringify(data),
        });
        if (resp.ok) {
            return Promise.resolve({
                ok: true,
                data: await resp.json(),
                status: resp.status,
            });
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
    async postForm(path, auth, data) {
        return await fetch(`${this.baseURL}${path}`, {
            method: "POST",
            headers: auth ? { Authorization: this.headers["Authorization"] } : {},
            body: data,
        });
    }
    async patchForm(path, auth, data) {
        return await fetch(`${this.baseURL}${path}`, {
            method: "PATCH",
            headers: auth ? { Authorization: this.headers["Authorization"] } : {},
            body: data,
        });
    }
    async delete(path, data) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: "DELETE",
            headers: this.headers,
            body: data ? JSON.stringify(data) : data,
        });
        if (resp.ok) {
            return Promise.resolve({
                ok: true,
                data: await resp.json(),
                status: resp.status,
            });
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
}

;// CONCATENATED MODULE: ./lib/apiv2/base/index.ts


class ApiBase {
    constructor(base_url, tenant_id, token) {
        this.base_url = base_url;
        this.tenant_id = tenant_id;
        this.user_token = token;
        this.http = new Http(base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
    }
    async init() {
        const resp = await fetch(`${apiURL(this.tenant_id)}/auth/refresh`, {
            method: "POST",
            body: JSON.stringify({
                path: ["basic"],
                user_token: this.user_token,
            }),
        });
        const rdata = await resp.json();
    }
    async get(path) {
        return this.http.get(path);
    }
    async post(path, data) {
        return this.http.post(path, data);
    }
    async postForm(path, auth, data) {
        return this.http.postForm(path, auth, data);
    }
    async patchForm(path, auth, data) {
        return this.http.patchForm(path, auth, data);
    }
    async put(path, data) {
        return this.http.put(path, data);
    }
    async patch(path, data) {
        return this.http.patch(path, data);
    }
    async delete(path, data) {
        return this.http.delete(path, data);
    }
}

;// CONCATENATED MODULE: ./lib/apiv2/admin/adapter_editor.ts

class AdapterEditorAPI {
    constructor(base_url, tenant_id, token) {
        this.base = new ApiBase(base_url, tenant_id, token);
    }
    perform_action(name, data) {
        return this.base.post(`/admin/adapter_editor/action/${name}`, data);
    }
    self_update(data) {
        return this.base.post("/admin/adapter_editor/", data);
    }
    self_reset() {
        return this.base.post("/admin/adapter_editor/reset", {});
    }
    // app
    list_apps() {
        return this.base.get("/admin/adapter_editor/app");
    }
    new_app(data) {
        return this.base.post("/admin/adapter_editor/app", data);
    }
    get_app(id) {
        return this.base.get(`/admin/adapter_editor/app/${id}`);
    }
    update_app(id, data) {
        return this.base.post(`/admin/adapter_editor/app/${id}`, data);
    }
    delete_app(id) {
        return this.base.delete(`/admin/adapter_editor/app/${id}`);
    }
    // hook
    list_hooks() {
        return this.base.get("/admin/adapter_editor/hook");
    }
    new_hook(data) {
        return this.base.post("/admin/adapter_editor/hook", data);
    }
    get_hook(id) {
        return this.base.get(`/admin/adapter_editor/hook/${id}`);
    }
    update_hook(id, data) {
        return this.base.post(`/admin/adapter_editor/hook/${id}`, data);
    }
    delete_hook(id) {
        return this.base.delete(`/admin/adapter_editor/hook/${id}`);
    }
}

;// CONCATENATED MODULE: ./lib/registry/registry.ts
class Registry {
    constructor() {
        this.RegisterFactory = (type, name, factory) => {
            console.log(`START REGISTER FACTORY => type(${type}) name(${name})`);
            const key = [type, name].toString();
            this._factories.set(key, factory);
            const watchers = this._watchers.get(key);
            if (watchers) {
                console.log("Found watchers ", watchers);
                watchers.forEach((watcher) => watcher());
            }
            const typeWatchers = this._type_watchers.get(type);
            if (typeWatchers) {
                typeWatchers.forEach((f) => f(factory));
            }
            console.log(`END REGISTER FACTORY => type(${type}) name(${name})`);
        };
        this.WatchLoad = async (type, name, timeout) => {
            console.log("before Watching");
            const key = [type, name].toString();
            if (this._factories.has(key)) {
                console.log("found factories already");
                return Promise.resolve();
            }
            const p = new Promise((resolve, reject) => {
                console.log("making promise");
                let oldwatcher = this._watchers.get(key);
                if (!oldwatcher) {
                    oldwatcher = new Array(0);
                }
                oldwatcher.push(() => {
                    resolve();
                });
                this._watchers.set(key, oldwatcher);
                setTimeout(() => {
                    reject(`TimeOut loading type ${type} & name ${name}`);
                }, timeout);
            });
            return p;
        };
        this.OnTypeLoad = (typename, callback) => {
            let oldwatcher = this._type_watchers.get(typename);
            if (!oldwatcher) {
                oldwatcher = new Array(0);
            }
            oldwatcher.push(callback);
        };
        this.Get = (type, name) => {
            const key = [type, name].toString();
            return this._factories.get(key.toString());
        };
        this.GetAll = (type) => {
            const facts = Array(0);
            this._factories.forEach((fact, [_type, _]) => {
                if (type !== _type) {
                    return;
                }
                facts.push(fact);
            });
            return facts;
        };
        this.InstanceAll = (type, opts) => {
            this._factories.forEach((fact, key) => {
                const [_type, _] = key.split(",");
                if (type !== _type) {
                    return;
                }
                fact(opts);
            });
        };
        this.Instance = (type, name, opts) => {
            const key = [type, name].toString();
            this._factories.get(key)(opts);
        };
        this._factories = new Map();
        this._watchers = new Map();
        this._type_watchers = new Map();
    }
}

;// CONCATENATED MODULE: ./lib/engine/putils.ts

const initRegistry = () => {
    console.log("init registry");
    if (window["__registry__"]) {
        console.warn("Registry already loaded, skipping...");
        return;
    }
    const r = new Registry();
    r.RegisterFactory("loader.factory", "std.loader", async (opts) => {
        await opts.registry.WatchLoad("plug.factory", opts.entry, 200000);
        const factory = opts.registry.Get("plug.factory", opts.entry);
        if (!factory) {
            console.warn("could not load plug factory");
            return;
        }
        factory({
            plug: opts.plug,
            agent: opts.agent,
            entry: opts.entry,
            env: opts.env,
            target: opts.target,
            payload: opts.payload,
            registry: opts.registry,
        });
    });
    window["__registry__"] = r;
    window["__register_factory__"] = r.RegisterFactory;
    window["RegisterFactory"] = r.RegisterFactory;
};
// it will find appoprate loader and call loader
// then its loader responsibility to start registered factories
// plugStart => loader => actual_plug_factory_start (using entry_name)
const plugStart = async (opts) => {
    console.log("let there be light", opts);
    const registry = window["__registry__"];
    if (!registry) {
        console.warn("registry not found");
        return;
    }
    if (!opts.exec_loader) {
        opts.exec_loader = "std.loader";
    }
    try {
        await registry.WatchLoad("loader.factory", opts.exec_loader, 100000);
    }
    catch (error) {
        console.warn("could not load, error occured:", error);
        return;
    }
    const loaderFactory = registry.Get("loader.factory", opts.exec_loader);
    if (!opts.target) {
        opts.target = document.body;
    }
    loaderFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        registry: registry,
        target: opts.target,
        payload: opts.payload,
    });
};

;// CONCATENATED MODULE: ./entries/adapter_editor_loader/index.ts



initRegistry();
window.addEventListener("load", async (ev) => {
    const registry = window["__registry__"];
    const loaderOpts = window["__loader_options__"] || {};
    const api = new AdapterEditorAPI(loaderOpts["base_url"], loaderOpts["tenant_id"], loaderOpts["adapter_editor_token"]);
    const env = new AdapterEditorEnv(api, loaderOpts["domain_name"] || "");
    console.log("@adapter_editor_loader", loaderOpts, env, registry);
    const adapterType = loaderOpts["adapter_type"] || "";
    let factory = registry.Get("temphia.adapter_editor.loader", `${adapterType}.main`);
    if (!factory) {
        await registry.WatchLoad("temphia.adapter_editor.loader", `${adapterType}.main`, 20000);
        factory = registry.Get("temphia.adapter_editor.loader", `${adapterType}.main`);
    }
    if (factory) {
        factory({
            env: env,
            target: document.getElementById("adapter-editor-root"),
        });
    }
    else {
        console.warn("@adapter_editor_loader", "factory not found");
    }
});

module.exports = __webpack_exports__;
/******/ })()
;