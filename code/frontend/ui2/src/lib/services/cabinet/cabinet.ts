import type { CabinetAPI, FolderTktAPI, SelfAPI } from "../apiv2";

interface Apm {
  get_cabinet(source: string): CabinetAPI;
  get_self_api(): SelfAPI;
  get_folder_api: (source: string, folder: string) => Promise<FolderTktAPI>
}

export class CabinetService {
  apm: Apm;
  source_apis: Map<string, CabinetAPI>;
  folder_apis: Map<string, FolderTktAPI>;
  sources: string[];

  constructor(apm: Apm) {
    this.apm = apm;
    this.source_apis = new Map();
    this.sources = null;
    this.folder_apis = new Map()
  }

  get_source_api(src: string) {
    if (this.source_apis.has(src)) {
      this.source_apis.get(src);
    }

    const capi = this.apm.get_cabinet(src);
    this.source_apis.set(src, capi);
    return capi;
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

  async get_folder_api(source: string, folder: string) {
    const key = `${source}#${folder}`

    let fapi = this.folder_apis.get(key)
    if (fapi) {
      return fapi
    }

    fapi = await this.apm.get_folder_api(source, folder)
    this.folder_apis.set(key, fapi);
    return fapi
  }
}
