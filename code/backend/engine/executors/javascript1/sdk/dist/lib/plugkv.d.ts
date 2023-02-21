export interface PlugValue {
    key: string;
    value: string;
    version: number;
    tag1?: string;
    tag2?: string;
    tag3?: string;
    ttl?: string;
}
export interface TicketOptions {
    state_tag1: string;
    state_tag2: string;
    state_tag3: string;
}
export interface PlugQuery {
    key_prefix?: string;
    load_meta?: boolean;
    tag1s?: string[];
    tag2s?: string[];
    tag3s?: string[];
    page_count?: number;
    page?: number;
}
export interface SetOptions {
    tag1?: string;
    tag2?: string;
    tag3?: string;
    ttl?: number;
}
export interface UpdateOptions {
    force_ver?: boolean;
    with_version?: boolean;
    version?: number;
    set_tag1?: boolean;
    set_tag2?: boolean;
    set_tag3?: boolean;
    tag1?: string;
    tag2?: string;
    tag3?: string;
    set_ttl?: boolean;
    ttl?: number;
}
export declare class PlugKV {
    txid: number;
    constructor(txid: number);
    get_ticket: (opts: TicketOptions) => [string, string];
    quick_get: (key: string) => [string, string];
    set: (key: string, value: string, opts?: SetOptions) => string;
    update: (key: string, value: string, opts?: UpdateOptions) => string;
    get: (key: string) => [PlugValue, string];
    query: (opts: PlugQuery) => [PlugValue[], string];
    delete: (key: string) => string;
    batch_delete: (keys: string[]) => string;
    new_txn: () => [PlugKV, string];
    rollback: () => string;
    commit: () => string;
}
export declare const plugkv: PlugKV;
