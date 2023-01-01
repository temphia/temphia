import type { ApiBase } from "../base";

export class AdapterEditorAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  PerformAction(name: string, data: any) {}
}
