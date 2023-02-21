"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ConstantBackoff = void 0;
/**
 * ConstantBackoff always returns the same value.
 */
class ConstantBackoff {
    constructor(backoff) {
        this.reset = () => {
            // no-op
        };
        this.backoff = backoff;
    }
    next() {
        return this.backoff;
    }
}
exports.ConstantBackoff = ConstantBackoff;
