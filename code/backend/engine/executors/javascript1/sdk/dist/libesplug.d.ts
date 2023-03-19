export declare class CabFolder {
    _folder: string;
    constructor(folder: string);
    add_file: (file: string) => string;
    list_folder: (file: string) => [string[], string];
    get_file: (file: string) => [ArrayBuffer, string];
    get_file_str: (file: string) => [string, string];
    del_file: (file: string) => string;
    gen_ticket: (opts: any) => [string, string];
}
export declare const core: {
    log: (message: string) => void;
    log_lazy: (message: string) => number;
    lazy_log_send: () => void;
    sleep: (t: number) => void;
    self_file: (file: string) => [string, string];
};
export declare class Response {
    status: number;
    headers: object;
    body: string;
    constructor(status: number, header: object, body: string);
    ok: () => boolean;
    get_header: (key: string) => any;
    json_body: () => any;
}
export declare class Request {
    _url: string;
    _headers: object;
    _body: string;
    constructor(url: string);
    set_header: (key: string, value: string) => void;
    set_body: (body: string) => void;
    set_json_body: (value: object) => void;
    get: () => Response;
    post: () => Response;
    put: () => Response;
    patch: () => Response;
    delete: () => Response;
}
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
export declare class SockdRoom {
    _room: string;
    constructor(room: string);
    send_direct: (connIds: number, value: string) => string;
    send_direct_batch: (connIds: number[], value: string) => string;
    send_broadcast: (value: string, ignores?: number[]) => string;
    send_tagged: (tags: string[], value: string, ignore?: number[]) => string;
    ticket: (opts: any) => [string, string];
}
interface Response {
    payload: {
        ok: boolean;
        data?: any;
        message?: string;
    };
}
export declare const utils: {
    is_db_not_found: (err: string) => boolean;
    is_db_already_exists: (err: string) => boolean;
    ok_response: (data: any) => Response;
    err_response: (message: string) => Response;
    ab2str: (buf: ArrayBuffer) => any;
    str2ab: (str: string) => ArrayBuffer;
    is_arraybuffer(value: any): boolean;
    generate_str_id: () => string;
};
export {};
