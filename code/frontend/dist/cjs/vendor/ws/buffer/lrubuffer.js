"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.LRUBuffer = void 0;
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
exports.LRUBuffer = LRUBuffer;
