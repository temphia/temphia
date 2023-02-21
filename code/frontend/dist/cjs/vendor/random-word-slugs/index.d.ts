import { PartsOfSpeech, Categories } from "./words";
interface FixedLengthArray<T extends any, L extends number> extends Array<T> {
    0: T;
    length: L;
}
declare type Case = "kebab" | "camel" | "title" | "lower" | "sentence";
declare type Options<T, L extends number> = {
    partsOfSpeech: FixedLengthArray<T, L>;
    categories: Partial<{
        [K in PartsOfSpeech]: Categories[K][];
    }>;
    format: Case;
};
export declare type RandomWordOptions<N extends number> = Partial<Options<PartsOfSpeech, N>>;
export declare function generateSlug<N extends number>(numberOfWords?: N, options?: Partial<Options<PartsOfSpeech, N>>): string;
export declare function totalUniqueSlugs<N extends number>(numberOfWords?: N, options?: RandomWordOptions<N>): number;
export {};
