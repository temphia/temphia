import type { ApiBase } from "../base";

export class EngineAPI {
  base: ApiBase;

  constructor(base: ApiBase) {
    this.base = base;
  }

  launch_target(data: any) {
    return this.base.post("/engine/launch/target", data);
  }

  launch_admin(data: any) {
    return this.base.post("/engine/launch/admin", data);
  }

  reset(plug_id: string, agent_id: string) {
    return this.base.post("/engine/reset", {
      plug_id,
      agent_id,
    });
  }
}
