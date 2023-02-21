import type { Buffer } from "./buffer";
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
export declare class LRUBuffer<E> implements Buffer<E> {
    private readonly buffer;
    private writePtr;
    private wrapped;
    constructor(len: number);
    len(): number;
    cap(): number;
    read(es: E[]): number;
    write(es: E[]): number;
    forEach(fn: (e: E) => any): number;
    clear(): void;
}
