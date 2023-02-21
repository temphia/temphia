"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Sockd = void 0;
const stypes_1 = require("./stypes");
const ws_1 = require("../vendor/ws");
const utils_1 = require("../utils");
class Sockd {
    constructor(url) {
        this.SendDirect = (data, target) => {
            this._ws.send(JSON.stringify({
                type: stypes_1.MESSAGE_CLIENT_DIRECT,
                xid: utils_1.generateId(),
                room: this._room,
                from_id: this._sid,
                targets: [target],
                payload: data,
            }));
        };
        this.SendBroadcast = (data) => {
            this._ws.send(JSON.stringify({
                type: stypes_1.MESSAGE_CLIENT_BROADCAST,
                xid: utils_1.generateId(),
                room: this._room,
                from_id: this._sid,
                payload: data,
            }));
        };
        this.SendTagged = (data, targets) => {
            this._ws.send(JSON.stringify({
                type: stypes_1.MESSAGE_CLIENT_PUBLISH,
                xid: utils_1.generateId(),
                room: this._room,
                from_id: this._sid,
                payload: data,
                target_tags: targets,
            }));
        };
        this.UpdateToken = (token) => {
            this._ws.send(JSON.stringify({
                type: stypes_1.MESSAGE_CLIENT_SYSTEM,
                xid: utils_1.generateId(),
                room: this._room,
                from_id: this._sid,
                payload: token,
            }));
        };
        this.Close = () => {
            this._ws.close(0, "closed by client");
        };
        this.handleIncoming = (_, ev) => {
            const data = JSON.parse(ev.data);
            console.log("@incoming_message", data);
            if (data.type === stypes_1.MESSAGE_SERVER_SYSTEM) {
                // fixme => handle_server_system_message
                console.log("@handle_server_system_message");
                return;
            }
            this._handler(data);
        };
        this._builder = new ws_1.WebsocketBuilder(url);
        this._builder.onMessage(this.handleIncoming);
        this._builder.withBackoff(new ws_1.LinearBackoff(0, 10, 100));
        this._builder.withBuffer(new ws_1.LRUBuffer(20));
        this._handler = null;
    }
    async Init() {
        this._ws = this._builder.build();
    }
    SetHandler(fn) {
        this._handler = fn;
    }
}
exports.Sockd = Sockd;
