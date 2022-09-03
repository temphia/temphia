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

;// CONCATENATED MODULE: ../../services/engine/pipe.ts
class IFramePipe {
    constructor(secret) {
        this.set_handler = (fn) => {
            this._handlers.add(fn);
        };
        this.remove_handler = (fn) => {
            this._handlers.delete(fn);
        };
        this.send = (xid, action, data) => {
            const message = JSON.stringify({
                xid,
                data,
                action,
                parent_secret: this._secret,
            });
            window.parent.postMessage(message, '*');
        };
        this._secret = secret;
        this._handlers = new Set();
        window.addEventListener('message', (ev) => {
            const decoded = JSON.parse(ev.data);
            this._handlers.forEach((fn) => fn(decoded.xid, decoded.action, decoded.data));
        });
    }
}

;// CONCATENATED MODULE: ../engine/registry/index.ts
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
                const [_type, _] = key.split(',');
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
const initFactory = () => {
    if (window["__registry__"]) {
        console.warn("Registry already loaded, skipping...");
        return;
    }
    const r = new Registry();
    r.RegisterFactory("loader.factory", "std.loader", async (opts) => {
        await opts.registry.WatchLoad("plug.factory", opts.entry, 2000);
        const factory = opts.registry.Get("plug.factory", opts.entry);
        if (!factory) {
            console.warn("could not load plug factory");
            return;
        }
    });
    console.log("GLOBAL_REGISTRY =>", r);
    window["__registry__"] = r;
    window["__register_factory__"] = r.RegisterFactory;
};
const startExecFactory = async (opts) => {
    console.log("Before starting factory", opts);
    const registry = window["__registry__"];
    if (!registry) {
        console.warn("registry not found");
        return;
    }
    if (!opts.exec_loader) {
        opts.exec_loader = "std.loader";
    }
    try {
        await registry.WatchLoad("loader.factory", opts.exec_loader, 10000);
    }
    catch (error) {
        console.warn("could not load, error occured:", error);
        return;
    }
    const plugFactory = registry.Get("loader.factory", opts.exec_loader);
    if (!opts.target) {
        opts.target = document.body;
    }
    plugFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: opts.env,
        registry: registry,
        target: opts.target,
        payload: opts.payload
    });
};

;// CONCATENATED MODULE: ../api/folder.ts
class FolderTktAPI {
    constructor(base_url, ticket) {
        this.ticket = ticket;
        this.base_url = base_url;
    }
    async list() {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}`);
        return resp.json();
    }
    async upload_file(file, data) {
        const resp = await fetch(`${this.base_url}/ticket_cabinet/${this.ticket}/${file}`, {
            method: "POST",
            body: data,
        });
        return resp.json();
    }
    get_file_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/${file}`;
    }
    get_file_preview_link(file) {
        return `${this.base_url}/ticket_cabinet/${this.ticket}/preview/${file}`;
    }
}

;// CONCATENATED MODULE: ../core/ws/backoff/linearbackoff.ts
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

;// CONCATENATED MODULE: ../core/ws/buffer/lrubuffer.ts
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

;// CONCATENATED MODULE: ../core/ws/websocket.ts
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
        if (this.closedByUser)
            return;
        if (this.websocket === undefined || this.websocket.readyState !== this.websocket.OPEN)
            this.buffer?.write([data]);
        else
            this.websocket.send(data);
    }
    close(code, reason) {
        this.closedByUser = true;
        this.websocket?.close(code, reason);
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
        switch (type) {
            case WebsocketEvents.close:
                if (!this.closedByUser) // failed to connect or connection lost, try to reconnect
                    this.reconnect();
                break;
            case WebsocketEvents.open:
                this.retries = 0;
                this.backoff?.reset(); // reset backoff
                this.buffer?.forEach(this.send.bind(this)); // send all buffered messages
                this.buffer?.clear();
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

;// CONCATENATED MODULE: ../core/ws/websocketBuilder.ts

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
        this.onOpenListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.open, h.listener, h.options));
        this.onCloseListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.close, h.listener, h.options));
        this.onErrorListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.error, h.listener, h.options));
        this.onMessageListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.message, h.listener, h.options));
        this.onRetryListeners.forEach(h => this.ws?.addEventListener(WebsocketEvents.retry, h.listener, h.options));
        return this.ws;
    }
}

;// CONCATENATED MODULE: ../core/ws/index.ts








;// CONCATENATED MODULE: ../sockd/sockd.ts

class Sockd {
    constructor(url) {
        this.init = async () => {
            this._ws = this._builder.build();
        };
        this.handleIncoming = (_, ev) => {
            // fixme => handle system messages
            const data = JSON.parse(ev.data);
            this._handler(data);
        };
        this.OnSockdMessage = (h) => {
            this._handler = h;
        };
        this.SendSockd = (message) => {
            this._ws.send(JSON.stringify(message));
        };
        console.log("CONNECTING WS @ ", url);
        this._builder = new WebsocketBuilder(url);
        this._builder.onMessage(this.handleIncoming);
        this._builder.withBackoff(new LinearBackoff(1, 3));
        this._builder.withBuffer(new LRUBuffer(20));
    }
}

;// CONCATENATED MODULE: ../sockd/stypes.ts
const MESSAGE_SERVER_DIRECT = "server_direct";
const MESSAGE_SERVER_BROADCAST = "server_broadcast";
const MESSAGE_SERVER_PUBLISH = "server_publish";
const MESSAGE_PEER_DIRECT = "peer_direct";
const MESSAGE_PEER_BROADCAST = "peer_broadcast";
const MESSAGE_PEER_PUBLISH = "peer_publish";


;// CONCATENATED MODULE: ../sockd/room.ts

class SockdRoom {
    constructor(socket, room) {
        this.SendDirect = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_DIRECT,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendBroadcast = (data) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_BROADCAST,
                xid: "",
                from_id: "",
                room: this._room,
            });
        };
        this.SendTagged = (data, ticket, targets) => {
            this._socket.SendSockd({
                payload: data,
                type: MESSAGE_PEER_PUBLISH,
                xid: "",
                from_id: "",
                room: this._room,
                targets: targets,
                ticket: ticket,
            });
        };
        this.onMessage = (handler) => {
            this._onMessage = handler;
        };
        this.onPeer = (handler) => {
            this._onPeer = handler;
        };
        this.onServer = (handler) => {
            this._onServer = handler;
        };
        this.ProcessMessage = (message) => {
            if (this._onMessage) {
                this._onMessage(message);
            }
            switch (message.type) {
                case MESSAGE_SERVER_DIRECT:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_SERVER_BROADCAST:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_SERVER_PUBLISH:
                    if (this._onServer) {
                        this._onServer(message);
                    }
                case MESSAGE_PEER_DIRECT:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case MESSAGE_PEER_BROADCAST:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                case MESSAGE_PEER_PUBLISH:
                    if (this._onPeer) {
                        this._onPeer(message);
                    }
                default:
                    break;
            }
        };
        this.IsConnected = async () => {
            return false;
        };
        this.LeaveRoom = () => {
            // fixme => impl
        };
        this._socket = socket;
        this._room = room;
    }
}

;// CONCATENATED MODULE: ../engine/env/fetch.ts
const actionFetch = (actionUrl, token) => async (name, data) => {
    const response = await fetch(`${actionUrl}/${name}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: token,
        },
        redirect: "follow",
        referrerPolicy: "strict-origin-when-cross-origin",
        body: data,
    });
    return response;
};

;// CONCATENATED MODULE: ../engine/env/env.ts




class Env {
    constructor(opts) {
        this.init = async () => {
            await this._sockd.init();
        };
        this.PreformAction = async (name, data) => {
            const encoded = JSON.stringify(data);
            try {
                const resp = await this._fetch(name, encoded);
                const ctype = resp.headers.get("Content-Type");
                if (resp.status !== 200) {
                    const txt = await resp.text();
                    return {
                        status_ok: false,
                        content_type: ctype,
                        body: txt,
                    };
                }
                const respData = await resp.json();
                return {
                    body: respData,
                    content_type: ctype,
                    status_ok: true,
                };
            }
            catch (error) {
                return {
                    status_ok: false,
                    body: error,
                };
            }
        };
        this.startup_payload = () => {
            return this._startup_payload;
        };
        this.PreformParentAction = async (name, data) => {
            const key = "fixme => generate";
            const p = new Promise((resolve, reject) => {
            });
            this._pending_pipe_msg.set(key, null);
            this._pipe.send("aaa", name, data);
            // fixme => implement
        };
        this.FolderTktAPI = (ticket) => {
            return new FolderTktAPI(this._opts.base_url, ticket);
        };
        this.SockdAPI = (room) => {
            let rs = this._sockd_rooms.get(room);
            if (!rs) {
                rs = new SockdRoom(this._sockd, room);
                this._sockd_rooms.set(room, rs);
            }
            return rs;
        };
        window["debug_env"] = this; // only for debug remove this 
        this._opts = opts;
        this._sockd_rooms = new Map();
        this._pending_pipe_msg = new Map();
        this._pipe = opts.pipe;
        this._startup_payload = opts.startup_payload;
        this._fetch = actionFetch(`${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_con`, opts.token);
        const sockdUrl = `${opts.base_url}engine/${opts.plug}/${opts.agent}/exec_ws`;
        this._sockd = new Sockd(sockdUrl);
        this._sockd.OnSockdMessage((msg) => {
            if (!msg.room) {
                console.log("no room message", msg);
                return;
            }
            if (msg.room === "plugs_dev") {
                console.log("PLUG DEBUG =>", msg.payload);
                return;
            }
            const room = this._sockd_rooms.get(msg.room);
            if (!room) {
                console.log("room without handler =>");
                return;
            }
            room.ProcessMessage(msg);
        });
    }
}

;// CONCATENATED MODULE: ../engine/env/index.ts


;// CONCATENATED MODULE: ./entry/entry.ts



console.log("loader init using...");
initFactory();
window.addEventListener("load", async () => {
    const opts = window["__loader_options__"];
    if (!opts) {
        console.log("Loader Options not found");
        return;
    }
    console.log("iframe portal opts @=>", opts);
    const pipe = new IFramePipe(opts.parent_secret);
    const env = new Env({
        agent: opts.agent,
        plug: opts.plug,
        token: opts.token,
        base_url: opts.base_url,
        parent_secret: opts.parent_secret,
        pipe,
    });
    await env.init();
    pipe.send("", "env_loaded", {});
    startExecFactory({
        plug: opts.plug,
        agent: opts.agent,
        entry: opts.entry,
        env: env,
        target: document.getElementById("plugroot"),
        exec_loader: opts.exec_loader,
        payload: null,
    });
}, false);

module.exports = __webpack_exports__;
/******/ })()
;