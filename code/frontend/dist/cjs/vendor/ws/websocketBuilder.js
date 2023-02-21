"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.WebsocketBuilder = void 0;
const websocket_1 = require("./websocket");
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
        this.ws = new websocket_1.Websocket(this.url, this.protocols, this.buffer, this.backoff);
        this.onOpenListeners.forEach(h => this.ws?.addEventListener(websocket_1.WebsocketEvents.open, h.listener, h.options));
        this.onCloseListeners.forEach(h => this.ws?.addEventListener(websocket_1.WebsocketEvents.close, h.listener, h.options));
        this.onErrorListeners.forEach(h => this.ws?.addEventListener(websocket_1.WebsocketEvents.error, h.listener, h.options));
        this.onMessageListeners.forEach(h => this.ws?.addEventListener(websocket_1.WebsocketEvents.message, h.listener, h.options));
        this.onRetryListeners.forEach(h => this.ws?.addEventListener(websocket_1.WebsocketEvents.retry, h.listener, h.options));
        return this.ws;
    }
}
exports.WebsocketBuilder = WebsocketBuilder;
