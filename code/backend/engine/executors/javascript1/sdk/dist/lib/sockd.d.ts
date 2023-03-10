export declare class SockdRoom {
    _room: string;
    constructor(room: string);
    send_direct: (connIds: string[], value: string) => string;
    send_broadcast: (value: string) => string;
    send_tagged: (tags: string[], value: string, ignore?: string[]) => string;
    add_to_room: (conn: string, tags: string[]) => string;
    kick_from_room: (conn: string) => string;
    list_room_conns: () => [object, string];
    bann_conn: (conn: string) => string;
    ticket: (room: string, opts: any) => [string, string];
}
