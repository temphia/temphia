"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Websocket = exports.WebsocketEvents = void 0;
var WebsocketEvents;
(function (WebsocketEvents) {
    WebsocketEvents["open"] = "open";
    WebsocketEvents["close"] = "close";
    WebsocketEvents["error"] = "error";
    WebsocketEvents["message"] = "message";
    WebsocketEvents["retry"] = "retry"; // A try to re-connect is made
})(WebsocketEvents = exports.WebsocketEvents || (exports.WebsocketEvents = {}));
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
exports.Websocket = Websocket;
