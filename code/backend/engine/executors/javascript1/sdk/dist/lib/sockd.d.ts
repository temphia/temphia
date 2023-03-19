export declare class SockdRoom {
    _room: string;
    constructor(room: string);
    send_direct: (connIds: number, value: string) => string;
    send_direct_batch: (connIds: number[], value: string) => string;
    send_broadcast: (value: string, ignores?: number[]) => string;
    send_tagged: (tags: string[], value: string, ignore?: number[]) => string;
    ticket: (opts: any) => [string, string];
}
