import type { SelfAPI } from "../../core/api";
import type { Sockd, SockdMessage } from "../../core/sockd";
import { SockdRoom } from "../../core/sockd";

const SOCKD_NOTIFICATION_ROOM = "sys.users";
const SOCKD_DTABLE_ROOM = "sys.dtable";

export class SockdService {
  _sockd: Sockd;
  _noti_room: SockdRoom;
  _dtable_room: SockdRoom;
  _basicAPi: SelfAPI;

  constructor(basic_api: SelfAPI, sockd: Sockd) {
    this._sockd = sockd;
    this._basicAPi = basic_api;
    this._noti_room = new SockdRoom(sockd, SOCKD_NOTIFICATION_ROOM);
    this._dtable_room = new SockdRoom(sockd, SOCKD_DTABLE_ROOM);
    sockd.OnSockdMessage(this.handle);
  }

  handle = (msg: SockdMessage) => {
    switch (msg.room) {
      case SOCKD_NOTIFICATION_ROOM:
        this._noti_room.ProcessMessage(msg);
        break;
      case SOCKD_DTABLE_ROOM:
        this._dtable_room.ProcessMessage(msg);
        break;
      default:
        console.log("Room not found", msg);
        break;
    }
  };

  get_notification_room = () => {
    return this._noti_room;
  };

  get_dyn_room = () => {
    return this._dtable_room;
  };

  change_group = async (source: string, group: string, ticket: string) => {
    return this._basicAPi.dtable_change({
      group,
      source,
      ticket,
    });
  };
}
