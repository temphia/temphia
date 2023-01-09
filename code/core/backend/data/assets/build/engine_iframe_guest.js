var __dirname = ''; var module = {}; module['exports']={};/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ 973:
/***/ (function(__unused_webpack_module, exports) {

(function (global, factory) {
     true ? factory(exports) :
    0;
}(this, (function (exports) { 'use strict';

    var default_sort = function (item, needle) { return item - needle; };
    function binarySearch(array, search, fn) {
        if (fn === void 0) { fn = default_sort; }
        var low = 0;
        var high = array.length - 1;
        var sort = fn.length === 1
            ? function (item, needle) { return fn(item) - search; }
            : fn;
        while (low <= high) {
            var i = (high + low) >> 1;
            var d = sort(array[i], search);
            if (d < 0) {
                low = i + 1;
            }
            else if (d > 0) {
                high = i - 1;
            }
            else {
                return i;
            }
        }
        return -low - 1;
    }

    function pickRandom(array) {
        var i = ~~(Math.random() * array.length);
        return array[i];
    }

    // http://bost.ocks.org/mike/shuffle/
    function shuffle(array) {
        var m = array.length;
        // While there remain elements to shuffle…
        while (m > 0) {
            // Pick a remaining element…
            var i = Math.floor(Math.random() * m--);
            // And swap it with the current element.
            var t = array[m];
            array[m] = array[i];
            array[i] = t;
        }
        return array;
    }

    function queue(max) {
        if (max === void 0) { max = 4; }
        var items = []; // TODO
        var pending = 0;
        var closed = false;
        var fulfil_closed;
        function dequeue() {
            if (pending === 0 && items.length === 0) {
                if (fulfil_closed)
                    fulfil_closed();
            }
            if (pending >= max)
                return;
            if (items.length === 0)
                return;
            pending += 1;
            var _a = items.shift(), fn = _a.fn, fulfil = _a.fulfil, reject = _a.reject;
            var promise = fn();
            try {
                promise.then(fulfil, reject).then(function () {
                    pending -= 1;
                    dequeue();
                });
            }
            catch (err) {
                reject(err);
                pending -= 1;
                dequeue();
            }
            dequeue();
        }
        return {
            add: function (fn) {
                if (closed) {
                    throw new Error("Cannot add to a closed queue");
                }
                return new Promise(function (fulfil, reject) {
                    items.push({ fn: fn, fulfil: fulfil, reject: reject });
                    dequeue();
                });
            },
            close: function () {
                closed = true;
                return new Promise(function (fulfil, reject) {
                    if (pending === 0) {
                        fulfil();
                    }
                    else {
                        fulfil_closed = fulfil;
                    }
                });
            }
        };
    }

    function sleep(ms) {
        return new Promise(function (fulfil) {
            setTimeout(fulfil, ms);
        });
    }

    function createSprite(width, height, fn) {
        var canvas = document.createElement('canvas');
        canvas.width = width;
        canvas.height = height;
        var ctx = canvas.getContext('2d');
        fn(ctx, canvas);
        return canvas;
    }

    function clamp(num, min, max) {
        return num < min ? min : num > max ? max : num;
    }

    function random(a, b) {
        if (b === undefined)
            return Math.random() * a;
        return a + Math.random() * (b - a);
    }

    function linear(domain, range) {
        var d0 = domain[0];
        var r0 = range[0];
        var m = (range[1] - r0) / (domain[1] - d0);
        return Object.assign(function (num) {
            return r0 + (num - d0) * m;
        }, {
            inverse: function () { return linear(range, domain); }
        });
    }

    // https://stackoverflow.com/questions/2901102/how-to-print-a-number-with-commas-as-thousands-separators-in-javascript
    function commas(num) {
        var parts = String(num).split('.');
        parts[0] = parts[0].replace(/\B(?=(\d{3})+(?!\d))/g, ',');
        return parts.join('.');
    }

    // array

    exports.binarySearch = binarySearch;
    exports.pickRandom = pickRandom;
    exports.shuffle = shuffle;
    exports.queue = queue;
    exports.sleep = sleep;
    exports.createSprite = createSprite;
    exports.clamp = clamp;
    exports.random = random;
    exports.linearScale = linear;
    exports.commas = commas;

    Object.defineProperty(exports, '__esModule', { value: true });

})));


/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __nccwpck_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		var threw = true;
/******/ 		try {
/******/ 			__webpack_modules__[moduleId].call(module.exports, module, module.exports, __nccwpck_require__);
/******/ 			threw = false;
/******/ 		} finally {
/******/ 			if(threw) delete __webpack_module_cache__[moduleId];
/******/ 		}
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
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
// This entry need to be wrapped in an IIFE because it need to be in strict mode.
(() => {
"use strict";
// ESM COMPAT FLAG
__nccwpck_require__.r(__webpack_exports__);

// EXTERNAL MODULE: ./node_modules/yootils/yootils.umd.js
var yootils_umd = __nccwpck_require__(973);
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
            return this._exec_api.preform_action(name, data);
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
        this._exec_api = new ExecAPI(opts.base_url.replace("v2/", "v2"), opts.token);
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
            try {
                const decoded = JSON.parse(ev.data);
                this.handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data));
            }
            catch (error) { }
        });
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/guestentry/iframe/start.ts





/* harmony default export */ const start = (() => {
    console.log("Iframe Exec start..");
    initRegistry();
    let transfered_port;
    const handle_port_transfer = (ev) => {
        if (ev.data !== "port_transfer") {
            console.log("wrong event listener", ev);
            return;
        }
        transfered_port = ev.ports[0];
        console.log("@received_port_@guest", transfered_port);
        window.removeEventListener("message", handle_port_transfer);
        env_init(null);
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
});

;// CONCATENATED MODULE: ./entries/execute_iframe_loader/index.ts

start();

})();

module.exports = __webpack_exports__;
/******/ })()
;