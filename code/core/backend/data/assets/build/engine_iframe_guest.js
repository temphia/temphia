/******/ (() => { // webpackBootstrap
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
    console.log("GLOBAL_REGISTRY =>", r);
    window["__registry__"] = r;
    window["__register_factory__"] = r.RegisterFactory;
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

;// CONCATENATED MODULE: ./lib/utils/index.ts
const generateId = () => Math.random().toString(36).slice(2);
const strHash = (str) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i);
        hash = (hash << 5) - hash + char;
        hash &= hash; // Convert to 32bit integer
    }
    return new Uint32Array([hash])[0].toString(36);
};
const pp = ".*(D#D01e-u0_ue819g_!UJ123456789023";
const numHash = (str) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 6) - hash);
        hash = pp.charCodeAt(i) ^ hash;
    }
    return hash;
};
const hslColor = (str) => {
    return `background: hsl(${numHash(str) % 360}, 100%, 80%)`;
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

;// CONCATENATED MODULE: ./lib/apiv2/engine/exec.ts

class ExecAPI {
    constructor(base_url, exec_token) {
        this.http = new Http(base_url, {
            Authorization: exec_token,
        });
    }
    agent_file_url(pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/serve/${file}`;
    }
    executor_file_url(eid, pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/executor/${eid}/${file}`;
    }
    ws_url(room_token) {
        return this.http.get(`/engine/ws?room_token=${room_token}`);
    }
    ws_update(room_token, data) {
        this.http.post(`/engine/ws?room_token=${room_token}`, data);
    }
    preform_action(method, data) {
        return this.http.post(`/engine/execute/${method}`, data);
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/env/index.ts


class Env {
    constructor(opts) {
        this.init = async () => { };
        // public
        this.PreformAction = async (name, data) => {
            return null;
        };
        this.startup_payload = () => {
            return this._startup_payload;
        };
        this.PreformParentAction = async (name, data) => {
            const key = generateId();
            const p = new Promise((resolve, reject) => {
                this._pending_pipe_msg.set(key, {
                    reject,
                    resolve,
                });
            });
            this._pipe.send(key, name, data);
            return p;
        };
        this.OnParentAction = (handler) => {
            this._default_parent_handler = handler;
        };
        this.GetRegistry = () => {
            return this._registry;
        };
        this.GetFolderTktAPI = (ticket) => {
            //return new FolderTktAPI(this._opts.base_url, ticket);
            return null;
        };
        this.GetRoomTktAPI = async (room, ticket) => {
            // const sroom  = new Sockd(`${this._sockd_url}?ticket=${ticket}`);
            // return sroom
            return null;
        };
        this.GetDataTableTktAPI = (ticket) => { };
        window["debug_env"] = this; // only for debug remove this
        this._registry = opts.registry;
        this._opts = opts;
        this._startup_payload = opts.startup_payload;
        this.set_up_pipe(opts.pipe);
        this._exec_api = new ExecAPI(opts.base_url, opts.token);
    }
    set_up_pipe(pipe) {
        this._pipe = pipe;
        this._pending_pipe_msg = new Map();
        this._pipe.set_handler((xid, action, data) => {
            const pending = this._pending_pipe_msg.get(xid);
            if (!pending) {
                if (this._default_parent_handler) {
                    this._default_parent_handler(data);
                }
                return;
            }
            pending.resolve(data);
        });
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/guestentry/iframe/iframe_pipe.ts
class IFramePipe {
    constructor(secret, port) {
        this.set_handler = (fn) => {
            this.handlers.add(fn);
        };
        this.remove_handler = (fn) => {
            this.handlers.delete(fn);
        };
        this.send = (xid, action, data) => {
            const message = JSON.stringify({
                xid,
                data,
                action,
                parent_secret: this.secret,
            });
            this.port.postMessage(message);
        };
        this.secret = secret;
        this.handlers = new Set();
        this.port = port;
        window.addEventListener("message", (ev) => {
            const decoded = JSON.parse(ev.data);
            this.handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data));
        });
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/guestentry/iframe/start.ts




/* harmony default export */ const start = (() => {
    console.log("Iframe Exec start..");
    initRegistry();
    let transfered_port;
    const handle_port_transfer = (ev) => {
        transfered_port = ev.ports[0];
        window.removeEventListener("message", handle_port_transfer);
    };
    const env_init = async (ev) => {
        const opts = window["__loader_options__"];
        if (!opts) {
            console.log("Loader Options not found");
            return;
        }
        console.log("iframe portal opts @=>", opts);
        const pipe = new IFramePipe(opts.parent_secret, transfered_port);
        const env = new Env({
            agent: opts.agent,
            plug: opts.plug,
            token: opts.token,
            base_url: opts.base_url,
            parent_secret: opts.parent_secret,
            pipe: pipe,
            registry: window["__registry__"],
        });
        await env.init();
        pipe.send(generateId(), "env_loaded", {});
        plugStart({
            plug: opts.plug,
            agent: opts.agent,
            entry: opts.entry,
            env: env,
            target: document.getElementById("plugroot"),
            exec_loader: opts.exec_loader,
            payload: null,
        });
    };
    window.addEventListener("message", handle_port_transfer, false);
    window.addEventListener("load", env_init, false);
});

;// CONCATENATED MODULE: ./entries/portal/launcher/guestentry/iframe/index.ts

start();

module.exports = __webpack_exports__;
/******/ })()
;