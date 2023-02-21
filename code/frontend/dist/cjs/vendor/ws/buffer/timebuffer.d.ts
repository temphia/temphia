import type { Buffer } from "./buffer";
/**
 * TimeBuffer keeps the elements that were written to the buffer
 * within maxAge milliseconds. For example, to  keep items in the
 * buffer that are less than a minute old, create the buffer with
 * maxAge equal to 60.000.
 *
 * When reading from the TimeBuffer, elements will be returned
 * in FIFO-order (queue).
 */
export declare class TimeBuffer<E> implements Buffer<E> {
    private readonly maxAge;
    private tail?;
    private head?;
    constructor(maxAge: number);
    cap(): number;
    len(): number;
    read(es: E[]): number;
    write(es: E[]): number;
    forEach(fn: (e: E) => any): number;
    private putElement;
    private forwardTail;
    clear(): void;
}
