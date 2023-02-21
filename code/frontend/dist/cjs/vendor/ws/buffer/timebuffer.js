"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.TimeBuffer = void 0;
/**
 * TimeBuffer keeps the elements that were written to the buffer
 * within maxAge milliseconds. For example, to  keep items in the
 * buffer that are less than a minute old, create the buffer with
 * maxAge equal to 60.000.
 *
 * When reading from the TimeBuffer, elements will be returned
 * in FIFO-order (queue).
 */
class TimeBuffer {
    constructor(maxAge) {
        this.maxAge = maxAge;
    }
    cap() {
        return Number.POSITIVE_INFINITY;
    }
    len() {
        this.forwardTail();
        let cur = this.tail;
        let i = 0;
        while (cur !== undefined) {
            i++;
            cur = cur.n;
        }
        return i;
    }
    read(es) {
        this.forwardTail();
        if (es.length === 0)
            return 0;
        let cur = this.tail;
        let i = 0;
        while (cur !== undefined) {
            es[i++] = cur.e;
            if (i === es.length)
                break;
            cur = cur.n;
        }
        return i;
    }
    write(es) {
        for (let i = 0; i < es.length; i++)
            this.putElement(es[i]);
        return es.length;
    }
    forEach(fn) {
        this.forwardTail();
        let cur = this.tail;
        let i = 0;
        while (cur !== undefined) {
            fn(cur.e);
            i++;
            cur = cur.n;
        }
        return i;
    }
    putElement(e) {
        const newElement = { e, t: Date.now(), n: undefined };
        if (this.tail === undefined)
            this.tail = newElement;
        if (this.head === undefined)
            this.head = newElement;
        else {
            this.head.n = newElement;
            this.head = newElement;
        }
    }
    forwardTail() {
        if (this.tail === undefined)
            return;
        const d = Date.now();
        while (d - this.tail.t > this.maxAge) {
            if (this.tail === this.head) {
                this.tail = undefined;
                this.head = undefined;
            }
            else
                this.tail = this.tail.n;
            if (this.tail === undefined)
                break;
        }
    }
    clear() {
        // TODO
    }
}
exports.TimeBuffer = TimeBuffer;
