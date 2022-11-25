import type { CabinetAPI, SelfAPI } from "../../../../lib/apiv2";

interface Apm {
  get_cabinet(source: string): CabinetAPI;
  get_self_api(): SelfAPI;
}

export class CabinetService {
  apm: Apm;
  source_apis: Map<string, CabinetAPI>;
  folder_apis: Map<string, any>;
  sources: string[];

  constructor(apm: Apm) {
    this.apm = apm;
    this.source_apis = new Map();
    this.sources = null;
  }

  get_source_api(src: string) {
    if (this.source_apis.has(src)) {
      this.source_apis.get(src);
    }

    const capi = this.apm.get_cabinet(src);
    this.source_apis.set(src, capi);
    return capi
  }

  async get_cab_sources() {
    if (this.sources) {
      return this.sources;
    }

    const sapi = this.apm.get_self_api();

    const resp = await sapi.list_cabinet_sources();
    if (!resp.ok) {
      return [];
    }
    this.sources = resp.data;
    return this.sources;
  }
}
