import type { Backoff } from "./backoff";
/**
 * LinearBackoff increases the backoff-time by a constant number with
 * every step. An optional maximum can be provided as an upper bound
 * to the returned backoff.
 *
 * Example: for initial=0, increment=2000, maximum=8000 the Linear-
 * Backoff will produce the series [0, 2000, 4000, 6000, 8000].
 */
export declare class LinearBackoff implements Backoff {
    private readonly initial;
    private readonly increment;
    private readonly maximum?;
    private current;
    constructor(initial: number, increment: number, maximum?: number);
    next(): number;
    reset(): void;
}
