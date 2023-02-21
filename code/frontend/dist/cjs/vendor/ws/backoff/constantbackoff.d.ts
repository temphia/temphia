import type { Backoff } from "./backoff";
/**
 * ConstantBackoff always returns the same value.
 */
export declare class ConstantBackoff implements Backoff {
    private readonly backoff;
    constructor(backoff: number);
    next(): number;
    reset: () => void;
}
