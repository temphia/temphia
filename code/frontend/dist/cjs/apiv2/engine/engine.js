"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.EngineAPI = void 0;
class EngineAPI {
    constructor(base) {
        this.base = base;
    }
    launch_target(data) {
        return this.base.post("/engine/launch/target", data);
    }
    launch_admin(data) {
        return this.base.post("/engine/launch/admin", data);
    }
}
exports.EngineAPI = EngineAPI;
