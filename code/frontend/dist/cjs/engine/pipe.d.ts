export interface PipeMessage {
    action?: string;
    xid: string;
    data: any;
    parent_secret?: string;
}
export declare type PipeHandler = (xid: string, action: string, data: any) => void;
export interface Pipe {
    send(xid: string, action: string, data: any): void;
    set_handler(fn: PipeHandler): void;
    remove_handler(fn: PipeHandler): void;
}
