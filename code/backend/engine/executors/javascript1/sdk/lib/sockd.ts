declare function _sd_send_direct(room: string, connIds: string[], value: string): string
declare function _sd_send_broadcast(room: string, value: string): string
declare function _sd_send_tagged(room: string, tags: string[], value: string, ignore?: string[]): string
declare function _sd_add_to_room(room: string, conn: string, tags: string[]): string
declare function _sd_kick_from_room(room: string, conn: string): string
declare function _sd_list_room_conns(room: string): [object, string]
declare function _sd_bann_conn(conn: string): string
declare function _sd_ticket(room: string, opts: any): [string, string]


export class SockdRoom {
    _room: string

    constructor(room: string) {
        this._room = room
    }

    send_direct = (connIds: string[], value: string): string => {
        return _sd_send_direct(this._room, connIds, value)
    }

    send_broadcast = (value: string): string => {
        return _sd_send_broadcast(this._room, value)
    }

    send_tagged = (tags: string[], value: string, ignore?: string[]): string => {
        return _sd_send_tagged(this._room, tags, value, ignore)
    }

    add_to_room = (conn: string, tags: string[]): string => {
        return _sd_add_to_room(this._room, conn, tags)
    }

    kick_from_room = (conn: string): string => {
        return _sd_kick_from_room(this._room, conn)
    }

    list_room_conns = (): [object, string] => {
        return _sd_list_room_conns(this._room)
    }

    bann_conn = (conn: string): string => {
        return _sd_bann_conn(conn)
    }

    ticket = (room: string, opts: any) => {
        return _sd_ticket(room, opts)
    }
}