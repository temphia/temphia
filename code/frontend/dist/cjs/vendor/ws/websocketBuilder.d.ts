import type { Backoff } from "./backoff/backoff";
import type { Buffer } from "./buffer/buffer";
import { RetryEventDetails, Websocket } from "./websocket";
/**
 * Used to build Websocket-instances.
 */
export declare class WebsocketBuilder {
    private readonly url;
    private ws;
    private protocols?;
    private backoff?;
    private buffer?;
    private onOpenListeners;
    private onCloseListeners;
    private onErrorListeners;
    private onMessageListeners;
    private onRetryListeners;
    constructor(url: string);
    withProtocols(p: string | string[]): WebsocketBuilder;
    withBackoff(backoff: Backoff): WebsocketBuilder;
    withBuffer(buffer: Buffer<any>): WebsocketBuilder;
    onOpen(listener: (instance: Websocket, ev: Event) => any, options?: boolean | EventListenerOptions): WebsocketBuilder;
    onClose(listener: (instance: Websocket, ev: CloseEvent) => any, options?: boolean | EventListenerOptions): WebsocketBuilder;
    onError(listener: (instance: Websocket, ev: Event) => any, options?: boolean | EventListenerOptions): WebsocketBuilder;
    onMessage(listener: (instance: Websocket, ev: MessageEvent) => any, options?: boolean | EventListenerOptions): WebsocketBuilder;
    onRetry(listener: (instance: Websocket, ev: CustomEvent<RetryEventDetails>) => any, options?: boolean | EventListenerOptions): WebsocketBuilder;
    /**
     * Multiple calls to build() will always return the same websocket-instance.
     */
    build(): Websocket;
}
