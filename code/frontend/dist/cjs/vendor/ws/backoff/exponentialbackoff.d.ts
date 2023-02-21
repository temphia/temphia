import type { Backoff } from "./backoff";
/**
 * ExponentialBackoff doubles the backoff with every step until a maximum
 * is reached. This is modelled after the binary exponential-backoff algo-
 * rithm used in computer-networking.
 *
 * The calculation-specification is:
 *          backoff = k * 2^s with s in [1, expMax].
 *
 * Example: for initial=100, expMax=7 the ExponentialBackoff will pro-
 * duce the backoff-series [100, 200, 400, 800, 1600, 3200, 6400].
 */
export declare class ExponentialBackoff implements Backoff {
    private readonly initial;
    private readonly expMax;
    private expCurrent;
    private current;
    constructor(initial: number, expMax: number);
    next(): number;
    reset(): void;
}
