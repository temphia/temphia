declare function _sd_send_direct( room: string, connId: number, value: string): string;
declare function _sd_send_direct_batch( room: string, connIds: number[], value: string): string;
declare function _sd_send_broadcast(room: string,value: string,ignore: number[]): string;
declare function _sd_send_tagged(room: string,tags: string[],value: string,ignore: number[]): string;
declare function _sd_ticket(room: string, opts: any): [string, string];

export class SockdRoom {
  _room: string;

  constructor(room: string) {
    this._room = room;
  }

  send_direct = (connIds: number, value: string): string => {
    return _sd_send_direct(this._room, connIds, value);
  };

  send_direct_batch = (connIds: number[], value: string) => {
    return _sd_send_direct_batch(this._room, connIds, value);
  };

  send_broadcast = (value: string, ignores?: number[]): string => {
    return _sd_send_broadcast(this._room, value, ignores ? ignores : []);
  };

  send_tagged = (tags: string[], value: string, ignore?: number[]): string => {
    return _sd_send_tagged(this._room, tags, value, ignore ? ignore : []);
  };

  ticket = (opts: any) => {
    return _sd_ticket(this._room, opts);
  };
}
