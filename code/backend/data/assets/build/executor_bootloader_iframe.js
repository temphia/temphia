var __dirname = ''; var module = {}; module['exports']={};/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ 662:
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
var yootils_umd = __nccwpck_require__(662);
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
const validateSlug = (v) => /^[a-z](-?[a-z])*$/.test(v);
const validateUserId = (v) => /^[a-z]+([a-z0-9_])+/.test(v);
const validateEmail = (v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v);
const units = (/* unused pure expression or super */ null && (["bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"]));
// makes bytes value nice
const humanizeBytes = (x) => {
    let l = 0, n = parseInt(x, 10) || 0;
    while (n >= 1024 && ++l) {
        n = n / 1024;
    }
    return n.toFixed(n < 10 && l > 0 ? 1 : 0) + " " + units[l];
};
const imageTypes = (/* unused pure expression or super */ null && (["png", "jpg", "jpeg"]));
const isImage = (name) => {
    const frags = name.split(".");
    return imageTypes.includes(frags[frags.length - 1]);
};
const fromGeoJsonOrFallback = (jstr) => {
    const fallback = [27.7116, 85.3124];
    try {
        const jpoint = JSON.parse(jstr);
        return jpoint["coordinates"] || fallback;
    }
    catch (error) {
        return fallback;
    }
};
const fromGeoJson = (jstr) => {
    try {
        const jpoint = JSON.parse(jstr);
        return jpoint["coordinates"] || [0, 0];
    }
    catch (error) {
        return [0, 0];
    }
};
const toGeoJson = (_lat, _lon) => {
    return JSON.stringify({
        type: "Point",
        coordinates: [_lat, _lon],
    });
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
            mode: "cors",
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
        this.http.headers = Object.assign(Object.assign({}, this.http.headers), { "Access-Control-Allow-Origin": "*" });
        return this.http.post(`/engine/execute/${method}`, data);
    }
}

;// CONCATENATED MODULE: ./lib/apiv2/data_sheet.ts

class DataSheetAPI {
    constructor(base_url, token) {
        this.list_users = (opts) => {
            return this.http.post(`/data/utils/user`, opts);
        };
        this.http = new Http(base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.base_url = base_url;
    }
    list_sheet_group() {
        return this.http.post(`/data/sheet/list`, {});
    }
    load_sheet(sheetid, options) {
        return this.http.post(`/data/sheet/${sheetid}/load`, options);
    }
    get_row_relation(sheetid, rid, refsheet, refcol) {
        return this.http.get(`/data/sheet/${sheetid}/relation/${rid}/ref/${refsheet}/column/${refcol}`);
    }
    // sheet
    list_sheets() {
        return this.http.get(`/data/sheet`);
    }
    get_sheet(sid) {
        return this.http.get(`/data/sheet/${sid}`);
    }
    new_sheet(data) {
        return this.http.post(`/data/sheet`, data);
    }
    update_sheet(sid, data) {
        return this.http.post(`/data/sheet/${sid}`, data);
    }
    delete_sheet(sid) {
        return this.http.delete(`/data/sheet/${sid}`);
    }
    // columns
    list_columns(sid) {
        return this.http.get(`/data/sheet/${sid}/column`);
    }
    get_column(sid, cid) {
        return this.http.get(`/data/sheet/${sid}/column/${cid}`);
    }
    new_column(sid, data) {
        return this.http.post(`/data/sheet/${sid}/column`, data);
    }
    update_column(sid, cid, data) {
        return this.http.post(`/data/sheet/${sid}/column/${cid}`, data);
    }
    delete_column(sid, cid) {
        return this.http.delete(`/data/sheet/${sid}/column/${cid}`);
    }
    // row_cells
    new_row_cell(sid, data) {
        return this.http.post(`/data/sheet/${sid}/row_cell`, data);
    }
    get_row_cell(sid, rid) {
        return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
    }
    update_row_cell(sid, rid, data) {
        return this.http.post(`/data/sheet/${sid}/row_cell/${rid}`, data);
    }
    delete_row_cell(sid, rid) {
        return this.http.delete(`/data/sheet/${sid}/row_cell/${rid}`);
    }
    search(sheetid, search) {
        return this.http.post(`/data/sheet/${sheetid}/search`, {
            search_term: search,
        });
    }
}

;// CONCATENATED MODULE: ./lib/apiv2/data.ts


class DataAPI {
    constructor(api_base_url, token) {
        this.sockd_url = () => {
            return `${this.api_base_url}/data_ws/?token=${this.token}`;
        };
        this.sheet_api = () => {
            return new DataSheetAPI(this.api_base_url, this.token);
        };
        this.list_users = (opts) => {
            return this.http.post(`/data/utils/user`, opts);
        };
        this.http = new Http(api_base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.api_base_url = api_base_url;
    }
    load() {
        return this.http.get(`/data`);
    }
    new_row(tid, data) {
        return this.http.post(`/data/table/${tid}/row`, data);
    }
    get_row(tid, rid) {
        return this.http.get(`/data/table/${tid}/row/${rid}`);
    }
    update_row(tid, rid, data) {
        return this.http.post(`/data/table/${tid}/row/${rid}`, data);
    }
    delete_row(tid, rid) {
        return this.http.delete(`/data/table/${tid}/row/${rid}`);
    }
    load_table(tid, opts) {
        return this.http.post(`/data/table/${tid}/load`, opts);
    }
    simple_query(tid, query) {
        return this.http.post(`/data/table/${tid}/simple_query`, query);
    }
    fts_query(tid, query) {
        return this.http.post(`/data/table/${tid}/fts_query`, query);
    }
    ref_load(tid, data) {
        return this.http.post(`/data/table/${tid}/ref_load`, data);
    }
    ref_resolve(tid, data) {
        return this.http.post(`/data/table/${tid}/ref_resolve`, data);
    }
    reverse_ref_load(tid, data) {
        return this.http.post(`/data/table/${tid}/rev_ref_load`, data);
    }
    list_activity(tid, rid) {
        return this.http.get(`/data/table/${tid}/activity/${rid}`);
    }
    comment_row(tid, rid, data) {
        return this.http.post(`/data/table/${tid}/activity/${rid}`, data);
    }
}

;// CONCATENATED MODULE: ./lib/exec/data.ts

const NewDataTableApi = (api_base_url, token) => {
    return new DataAPI(api_base_url, token);
};

;// CONCATENATED MODULE: ./lib/apiv2/cabinet.ts

class CabinetAPI {
    constructor(source, base) {
        this.source = source;
        this.base = base;
    }
    listRoot() {
        return this.base.get(`/cabinet/${this.source}/`);
    }
    listFolder(folder) {
        return this.base.get(`/cabinet/${this.source}/${folder}`);
    }
    newFolder(folder) {
        return this.base.post(`/cabinet/${this.source}/${folder}`, {});
    }
    getFile(folder, fname) {
        return this.base.get(`/cabinet/${this.source}/${folder}/file/${fname}`);
    }
    uploadFile(folder, fname, data) {
        return this.base.postForm(`/cabinet/${this.source}/${folder}/file/${fname}`, true, data);
    }
    deleteFile(folder, fname) {
        return this.base.delete(`/cabinet/${this.source}/${folder}/file/${fname}`);
    }
    getFilePreview(folder, fname) {
        return `${this.base.api_base_url}/cabinet/${this.source}/${folder}/preview/${fname}?token=${this.base.user_token}`;
    }
}
class FolderTktAPI {
    constructor(baseUrl, token) {
        this.http = new Http(baseUrl, {});
        this.ticket = token;
        this.base_url = baseUrl;
    }
    //  /folder/:ticket/
    list() {
        return this.http.get(`/folder/${this.ticket}`);
    }
    getFile(file) {
        return this.http.get(`/folder/${this.ticket}/${file}`);
    }
    getFileUrl(file) {
        return `${this.base_url}/folder/${this.ticket}/${file}`;
    }
    getFilePreviewUrl(file) {
        return `${this.base_url}/folder/${this.ticket}/${file}/preview`;
    }
    uploadFile(file, data) {
        return this.http.post(`/folder/${this.ticket}/${file}`, data);
    }
    // downgraded_ticket() {}
    deleteFile(file) {
        return this.http.delete(`/folder/${this.ticket}/${file}`);
    }
}

;// CONCATENATED MODULE: ./lib/exec/folder.ts

const NewFolderApi = (api_base_url, token) => {
    return new FolderTktAPI(api_base_url, token);
};

;// CONCATENATED MODULE: ./lib/apiv2/admin/plug_state.ts

class AdminPlugStateTktAPI {
    constructor(api_base_url, token) {
        this.http = new Http(api_base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.api_base_url = api_base_url;
    }
    query(options) {
        return this.http.post(`/admin/plug_state/query`, options);
    }
    add(key, value, opts) {
        return this.http.post(`/admin/plug_state/key`, {
            key,
            value,
            options: opts,
        });
    }
    update(key, value, opts) {
        return this.http.post(`/admin/plug_state/key/${key}`, {
            key,
            value,
            options: opts,
        });
    }
    delete(key) {
        return this.http.delete(`/admin/plug_state/key/${key}`, {
            key,
        });
    }
    get(key) {
        return this.http.get(`/admin/plug_state/key/${key}`);
    }
}

;// CONCATENATED MODULE: ./lib/exec/plug_state.ts

const NewPlugStateApi = (api_base_url, token) => {
    return new AdminPlugStateTktAPI(api_base_url, token);
};

;// CONCATENATED MODULE: ./lib/sockd/stypes.ts
const MESSAGE_SERVER_DIRECT = "server_direct";
const MESSAGE_SERVER_BROADCAST = "server_broadcast";
const MESSAGE_SERVER_PUBLISH = "server_publish";
const MESSAGE_CLIENT_DIRECT = "client_direct";
const MESSAGE_CLIENT_BROADCAST = "client_broadcast";
const MESSAGE_CLIENT_PUBLISH = "client_publish";
const MESSAGE_CLIENT_SYSTEM = "client_system";
const MESSAGE_SERVER_SYSTEM = "server_system";


;// CONCATENATED MODULE: ./lib/vendor/ws/backoff/linearbackoff.ts
/**
 * LinearBackoff increases the backoff-time by a constant number with
 * every step. An optional maximum can be provided as an upper bound
 * to the returned backoff.
 *
 * Example: for initial=0, increment=2000, maximum=8000 the Linear-
 * Backoff will produce the series [0, 2000, 4000, 6000, 8000].
 */
class LinearBackoff {
    constructor(initial, increment, maximum) {
        this.initial = initial;
        this.increment = increment;
        this.maximum = maximum;
        this.current = this.initial;
    }
    next() {
        const backoff = this.current;
        const next = this.current + this.increment;
        if (this.maximum === undefined)
            this.current = next;
        else if (next <= this.maximum)
            this.current = next;
        return backoff;
    }
    reset() {
        this.current = this.initial;
    }
}

;// CONCATENATED MODULE: ./lib/vendor/ws/buffer/lrubuffer.ts
/**
 * LRUBuffer is a buffer that keeps the last n elements. When it is
 * full and written to, the oldest element in the buffer will be
 * replaced. When reading from the LRUBuffer, elements are returned
 * in FIFO-order (queue).
 *
 * LRUBuffer has linear space- and time-requirements. Internally
 * an array is used as a circular-buffer. All memory is allocated
 * on initialization.
 */
class LRUBuffer {
    constructor(len) {
        this.writePtr = 0;
        this.wrapped = false;
        this.buffer = Array(len);
    }
    len() {
        return this.wrapped ? this.buffer.length : this.writePtr;
    }
    cap() {
        return this.buffer.length;
    }
    read(es) {
        if (es === null || es === undefined || es.length === 0 || this.buffer.length === 0)
            return 0;
        if (this.writePtr === 0 && !this.wrapped)
            return 0;
        const first = this.wrapped ? this.writePtr : 0;
        const last = (first - 1) < 0 ?
            this.buffer.length - 1 :
            first - 1;
        for (let i = 0; i < es.length; i++) {
            let r = (first + i) % this.buffer.length;
            es[i] = this.buffer[r];
            if (r === last)
                return i + 1;
        }
        return es.length;
    }
    write(es) {
        if (es === null || es === undefined || es.length === 0 || this.buffer.length === 0)
            return 0;
        const start = es.length > this.buffer.length ? es.length - this.buffer.length : 0;
        for (let i = 0; i < es.length - start; i++) {
            this.buffer[this.writePtr] = es[start + i];
            this.writePtr = (this.writePtr + 1) % this.buffer.length;
            if (this.writePtr === 0)
                this.wrapped = true;
        }
        return es.length;
    }
    forEach(fn) {
        if (this.writePtr === 0 && !this.wrapped)
            return 0;
        let cur = this.wrapped ? this.writePtr : 0;
        const last = this.wrapped ? (cur - 1) < 0 ? this.buffer.length - 1 : cur - 1 : this.writePtr - 1;
        const len = this.len();
        while (true) {
            fn(this.buffer[cur]);
            if (cur === last)
                break;
            cur = (cur + 1) % this.buffer.length;
        }
        return len;
    }
    clear() {
        this.writePtr = 0;
        this.wrapped = false;
    }
}

;// CONCATENATED MODULE: ./lib/vendor/ws/websocket.ts
var WebsocketEvents;
(function (WebsocketEvents) {
    WebsocketEvents["open"] = "open";
    WebsocketEvents["close"] = "close";
    WebsocketEvents["error"] = "error";
    WebsocketEvents["message"] = "message";
    WebsocketEvents["retry"] = "retry"; // A try to re-connect is made
})(WebsocketEvents || (WebsocketEvents = {}));
class Websocket {
    constructor(url, protocols, buffer, backoff) {
        this.eventListeners = { open: [], close: [], error: [], message: [], retry: [] };
        this.closedByUser = false;
        this.retries = 0;
        this.handleOpenEvent = (ev) => this.handleEvent(WebsocketEvents.open, ev);
        this.handleCloseEvent = (ev) => this.handleEvent(WebsocketEvents.close, ev);
        this.handleErrorEvent = (ev) => this.handleEvent(WebsocketEvents.error, ev);
        this.handleMessageEvent = (ev) => this.handleEvent(WebsocketEvents.message, ev);
        this.url = url;
        this.protocols = protocols;
        this.buffer = buffer;
        this.backoff = backoff;
        this.tryConnect();
    }
    getUnderlyingWebsocket() {
        return this.websocket;
    }
    send(data) {
        var _a;
        if (this.closedByUser)
            return;
        if (this.websocket === undefined || this.websocket.readyState !== this.websocket.OPEN)
            (_a = this.buffer) === null || _a === void 0 ? void 0 : _a.write([data]);
        else
            this.websocket.send(data);
    }
    close(code, reason) {
        var _a;
        this.closedByUser = true;
        (_a = this.websocket) === null || _a === void 0 ? void 0 : _a.close(code, reason);
    }
    addEventListener(type, listener, options) {
        const eventListener = { listener, options };
        const eventListeners = this.eventListeners[type];
        eventListeners.push(eventListener);
    }
    removeEventListener(type, listener, options) {
        this.eventListeners[type] =
            this.eventListeners[type]
                .filter(l => {
                return l.listener !== listener && (l.options === undefined || l.options !== options);
            });
    }
    dispatchEvent(type, ev) {
        const listeners = this.eventListeners[type];
        const onceListeners = [];
        listeners.forEach(l => {
            l.listener(this, ev); // call listener
            if (l.options !== undefined && l.options.once)
                onceListeners.push(l);
        });
        onceListeners.forEach(l => this.removeEventListener(type, l.listener, l.options)); // remove 'once'-listeners
    }
    tryConnect() {
        if (this.websocket !== undefined) { // remove all event-listeners from broken socket
            this.websocket.removeEventListener(WebsocketEvents.open, this.handleOpenEvent);
            this.websocket.removeEventListener(WebsocketEvents.close, this.handleCloseEvent);
            this.websocket.removeEventListener(WebsocketEvents.error, this.handleErrorEvent);
            this.websocket.removeEventListener(WebsocketEvents.message, this.handleMessageEvent);
            this.websocket.close();
        }
        this.websocket = new WebSocket(this.url, this.protocols); // create new socket and attach handlers
        this.websocket.addEventListener(WebsocketEvents.open, this.handleOpenEvent);
        this.websocket.addEventListener(WebsocketEvents.close, this.handleCloseEvent);
        this.websocket.addEventListener(WebsocketEvents.error, this.handleErrorEvent);
        this.websocket.addEventListener(WebsocketEvents.message, this.handleMessageEvent);
    }
    handleEvent(type, ev) {
        var _a, _b, _c;
        switch (type) {
            case WebsocketEvents.close:
                if (!this.closedByUser) // failed to connect or connection lost, try to reconnect
                    this.reconnect();
                break;
            case WebsocketEvents.open:
                this.retries = 0;
                (_a = this.backoff) === null || _a === void 0 ? void 0 : _a.reset(); // reset backoff
                (_b = this.buffer) === null || _b === void 0 ? void 0 : _b.forEach(this.send.bind(this)); // send all buffered messages
                (_c = this.buffer) === null || _c === void 0 ? void 0 : _c.clear();
                break;
        }
        this.dispatchEvent(type, ev); // forward to all listeners
    }
    reconnect() {
        if (this.backoff === undefined) // no backoff, we're done
            return;
        const backoff = this.backoff.next();
        setTimeout(() => {
            this.dispatchEvent(WebsocketEvents.retry, new CustomEvent(WebsocketEvents.retry, {
                detail: {
                    retries: ++this.retries,
                    backoff: backoff
                }
            }));
            this.tryConnect();
        }, backoff);
    }
}

;// CONCATENATED MODULE: ./lib/vendor/ws/websocketBuilder.ts

/**
 * Used to build Websocket-instances.
 */
class WebsocketBuilder {
    constructor(url) {
        this.ws = null;
        this.onOpenListeners = [];
        this.onCloseListeners = [];
        this.onErrorListeners = [];
        this.onMessageListeners = [];
        this.onRetryListeners = [];
        this.url = url.replace("http://", "ws://").replace("https://", "wss://");
    }
    withProtocols(p) {
        this.protocols = p;
        return this;
    }
    withBackoff(backoff) {
        this.backoff = backoff;
        return this;
    }
    withBuffer(buffer) {
        this.buffer = buffer;
        return this;
    }
    onOpen(listener, options) {
        this.onOpenListeners.push({ listener, options });
        return this;
    }
    onClose(listener, options) {
        this.onCloseListeners.push({ listener, options });
        return this;
    }
    onError(listener, options) {
        this.onErrorListeners.push({ listener, options });
        return this;
    }
    onMessage(listener, options) {
        this.onMessageListeners.push({ listener, options });
        return this;
    }
    onRetry(listener, options) {
        this.onRetryListeners.push({ listener, options });
        return this;
    }
    /**
     * Multiple calls to build() will always return the same websocket-instance.
     */
    build() {
        if (this.ws !== null)
            return this.ws;
        this.ws = new Websocket(this.url, this.protocols, this.buffer, this.backoff);
        this.onOpenListeners.forEach(h => { var _a; return (_a = this.ws) === null || _a === void 0 ? void 0 : _a.addEventListener(WebsocketEvents.open, h.listener, h.options); });
        this.onCloseListeners.forEach(h => { var _a; return (_a = this.ws) === null || _a === void 0 ? void 0 : _a.addEventListener(WebsocketEvents.close, h.listener, h.options); });
        this.onErrorListeners.forEach(h => { var _a; return (_a = this.ws) === null || _a === void 0 ? void 0 : _a.addEventListener(WebsocketEvents.error, h.listener, h.options); });
        this.onMessageListeners.forEach(h => { var _a; return (_a = this.ws) === null || _a === void 0 ? void 0 : _a.addEventListener(WebsocketEvents.message, h.listener, h.options); });
        this.onRetryListeners.forEach(h => { var _a; return (_a = this.ws) === null || _a === void 0 ? void 0 : _a.addEventListener(WebsocketEvents.retry, h.listener, h.options); });
        return this.ws;
    }
}

;// CONCATENATED MODULE: ./lib/vendor/ws/index.ts








;// CONCATENATED MODULE: ./lib/sockd/sockd.ts



class Sockd {
    constructor(url) {
        this.SendDirect = (data, target) => {
            this._ws.send(JSON.stringify({
                type: MESSAGE_CLIENT_DIRECT,
                xid: generateId(),
                room: this._room,
                from_id: this._sid,
                targets: [target],
                payload: data,
            }));
        };
        this.SendBroadcast = (data) => {
            this._ws.send(JSON.stringify({
                type: MESSAGE_CLIENT_BROADCAST,
                xid: generateId(),
                room: this._room,
                from_id: this._sid,
                payload: data,
            }));
        };
        this.SendTagged = (data, targets) => {
            this._ws.send(JSON.stringify({
                type: MESSAGE_CLIENT_PUBLISH,
                xid: generateId(),
                room: this._room,
                from_id: this._sid,
                payload: data,
                target_tags: targets,
            }));
        };
        this.UpdateToken = (token) => {
            this._ws.send(JSON.stringify({
                type: MESSAGE_CLIENT_SYSTEM,
                xid: generateId(),
                room: this._room,
                from_id: this._sid,
                payload: token,
            }));
        };
        this.Close = () => {
            this._ws.close(0, "closed by client");
        };
        this.handleIncoming = (_, ev) => {
            const data = JSON.parse(ev.data);
            console.log("@incoming_message", data);
            if (data.type === MESSAGE_SERVER_SYSTEM) {
                // fixme => handle_server_system_message
                console.log("@handle_server_system_message");
                return;
            }
            this._handler(data);
        };
        this._builder = new WebsocketBuilder(url);
        this._builder.onMessage(this.handleIncoming);
        this._builder.withBackoff(new LinearBackoff(0, 10, 100));
        this._builder.withBuffer(new LRUBuffer(20));
        this._handler = null;
    }
    async Init() {
        this._ws = this._builder.build();
    }
    SetHandler(fn) {
        this._handler = fn;
    }
}

;// CONCATENATED MODULE: ./lib/sockd/index.ts



;// CONCATENATED MODULE: ./lib/exec/sockd.ts

const NewSockdRoom = async (url, callback) => {
    const sockd = new Sockd(url);
    sockd.SetHandler(callback);
    await sockd.Init();
    return sockd;
};

;// CONCATENATED MODULE: ./lib/exec/exec_am.ts




// ExecAM stands for execution api manager
class ExecAM {
    constructor(api_base_url) {
        this.new_data_api = (token) => {
            return NewDataTableApi(this.api_base_url, token);
        };
        this.new_folder_api = (token) => {
            return NewFolderApi(this.api_base_url, token);
        };
        this.new_sockd_room = async (token, callback) => {
            return NewSockdRoom(`${this.api_base_url}/engine/ws?ticket=${token}`, callback);
        };
        this.new_sockd_room_from_url = async (url, callback) => {
            return NewSockdRoom(url, callback);
        };
        this.new_plug_state = (token) => {
            return NewPlugStateApi(this.api_base_url, token);
        };
        this.api_base_url = api_base_url;
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/env/asset_manager.ts
class EnvAssetManager {
    constructor(baseURL, plugId, agentId, exec) {
        this.agent_url = (name) => `${this.baseURL}/engine/plug/${this.plugId}/agent/${this.agentId}/serve/${name}`;
        this.baseURL = baseURL;
        this.plugId = plugId;
        this.agentId = agentId;
        this.executor = exec;
    }
    GetAgentAssetURL(name) {
        return this.agent_url(name);
    }
    GetExecutorAssetURL(name) {
        return `${this.baseURL}/engine/plug/${this.plugId}/agent/${this.agentId}/executor/${this.executor}/${name}`;
    }
    ImportDyanmic(name) {
        // fixme => impl
        return Promise.resolve();
    }
    SheduleWorker(name) {
        return new Worker(this.agent_url(name));
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/env/env.ts




class Env {
    constructor(opts) {
        // public
        this.PreformAction = async (name, data) => {
            return this.preformAction(name, data);
        };
        this.PreformParentAction = async (name, data) => {
            return this.preformParentAction(name, data);
        };
        this.OnParentAction = (handler) => {
            this._default_parent_handler = handler;
        };
        this.GetRegistry = () => {
            return this._registry;
        };
        this.preformAction = (name, data) => {
            return this._exec_api.preform_action(name, data);
        };
        this.preformParentAction = async (name, data) => {
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
        this.execApiManager = () => {
            return new EnvAssetManager(this._opts.api_base_url, this._opts.plug, this._opts.agent, "FIXME");
        };
        window["debug_env"] = this; // only for debug remove this
        this._registry = opts.registry;
        this._opts = opts;
        this._startup_payload = opts.startup_payload;
        this.set_up_pipe(opts.pipe);
        this._exec_api = new ExecAPI(opts.api_base_url.replace("v2/", "v2"), opts.token);
    }
    async init() { }
    GetExecVars() {
        return {
            agent_id: this._opts.agent,
            plug_id: this._opts.plug,
            api_base_url: this._opts.api_base_url,
            tenant_id: this._opts.tenant_id,
            exec_data: this._startup_payload,
        };
    }
    GetAssetManager() {
        return this.execApiManager();
    }
    GetExecApiManager() {
        return Promise.resolve(new ExecAM(this._opts.api_base_url));
    }
    // private
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
    Extend(s) {
        this._registry.GetAll("extensions").forEach((factory) => {
            const mods = factory({
                service: s,
                env: this,
                target: this._opts.target,
            });
            if (mods) {
                s = mods;
            }
        });
        return s;
    }
}

;// CONCATENATED MODULE: ./entries/portal/launcher/env/index.ts


;// CONCATENATED MODULE: ./entries/executor/@bootloader/iframe/iframe_pipe.ts
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

;// CONCATENATED MODULE: ./entries/executor/@bootloader/iframe/fakestorage.ts
function fakeStorage() {
    let storage = {};
    return {
        clear: function () {
            storage = {};
        },
        setItem: function (key, value) {
            storage[key] = value || "";
        },
        getItem: function (key) {
            return key in storage ? storage[key] : null;
        },
        removeItem: function (key) {
            delete storage[key];
        },
        get length() {
            return Object.keys(storage).length;
        },
        key: function (i) {
            const keys = Object.keys(storage);
            return keys[i] || null;
        },
    };
}

;// CONCATENATED MODULE: ./entries/executor/@bootloader/iframe/start.ts






/* harmony default export */ const start = (() => {
    console.log("Iframe Exec start..1111");
    try {
        if (!window["localStorage"]) {
            window["localStorage"] = fakeStorage();
        }
        else {
            // if we do not have sandbox and browser donot support
            //   credentialless/anonymous iframes #hacky_as_hell
            if (localStorage.getItem("_temphia_authed_key_")) {
                delete window["localStorage"];
                window["localStorage"] = fakeStorage();
            }
        }
    }
    catch (error) {
        console.log("@execption", error);
    }
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
        const target = document.getElementById("plugroot");
        const env = new Env({
            agent: opts.agent,
            plug: opts.plug,
            token: opts.token,
            api_base_url: opts.api_base_url,
            parent_secret: opts.parent_secret,
            pipe: pipe,
            registry: window["__registry__"],
            tenant_id: opts.tenant_id,
            target: target,
            startup_payload: opts.startup_payload,
        });
        await env.init();
        pipe.send(generateId(), "env_loaded", {});
        plugStart({
            plug: opts.plug,
            agent: opts.agent,
            entry: opts.entry,
            env: env,
            target: target,
            exec_loader: opts.exec_loader,
        });
    };
    window.addEventListener("message", handle_port_transfer, false);
});

;// CONCATENATED MODULE: ./entries/executor/@bootloader/iframe/index.ts

start();

})();

module.exports = __webpack_exports__;
/******/ })()
;