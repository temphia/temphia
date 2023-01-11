declare function _pkv_set(tx: number, key: string, value: string, opts?: any): string
declare function _pkv_update(tx: number, key: string, value: string, opts?: any): string
declare function _pkv_get(tx: number, key: string): [PlugValue, string];
declare function _pkv_query(tx: number, opts: PlugQuery): [PlugValue[], string];
declare function _pkv_del(tx: number, key: string): string;
declare function _pkv_batch_del(tx: number, keys: string[]): string;
declare function _pkv_new_txn(): [number, string];
declare function _pkv_rollback(tx: number): string;
declare function _pkv_commit(tx: number): string;


export interface PlugValue {
    key: string
    value: string
    version: number
    // meta
    tag1?: string
    tag2?: string
    tag3?: string
    ttl?: string
}

export interface PlugQuery {
    key_prefix?: string
    load_meta?: boolean
    tag1s?: string[]
    tag2s?: string[]
    tag3s?: string[]
    page_count?: number
    page?: number
}

export interface SetOptions {
    tag1?: string
    tag2?: string
    tag3?: string
    ttl?: number
}

export interface UpdateOptions {
    force_ver?: boolean
    with_version?: boolean
    version?: number
    set_tag1?: boolean
    set_tag2?: boolean
    set_tag3?: boolean
    tag1?: string
    tag2?: string
    tag3?: string
    set_ttl?: boolean
    ttl?: number
}

export class PlugKV {
    txid: number
    constructor(txid: number) {
        this.txid = txid
    }

    quick_get = (key: string): [string, string] => {
        const [val, err] = _pkv_get(this.txid, key)
        if (err) {
            return [null, err]
        }
        return [val.value, null]
    }

    set = (key: string, value: string, opts?: SetOptions): string => {
        return _pkv_set(this.txid, key, value, opts)
    }

    update = (key: string, value: string, opts?: UpdateOptions): string => {
        return _pkv_update(this.txid, key, value, opts)
    }


    get = (key: string): [PlugValue, string] => {
        const [val, err] = _pkv_get(this.txid, key)
        if (err) {
            return [null, err]
        }
        return [val, null]
    }


    query = (opts: PlugQuery): [PlugValue[], string] => {
        return _pkv_query(this.txid, opts)
    }

    delete = (key: string): string => {
        return _pkv_del(this.txid, key)
    }

    batch_delete = (keys: string[]): string => {
        return _pkv_batch_del(this.txid, keys)
    }

    new_txn = (): [PlugKV, string] => {
        const [newtxid, err] = _pkv_new_txn()
        if (err) {
            return [null, err]
        }
        return [new PlugKV(newtxid), null]
    }

    rollback = (): string => {
        return _pkv_rollback(this.txid)
    }

    commit = (): string => {
        return _pkv_commit(this.txid)
    }
}

export const plugkv = new PlugKV(0);