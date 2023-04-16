import type { Environment } from "../../../../lib/engine/environment";

export class PageQueryService {
  env: Environment;
  constructor(env: Environment) {
    this.env = env;
  }
}
