import type { CabinetAPI } from "../../../../lib/apiv2";

interface Apm {
  get_cabinet(source: string): CabinetAPI;
}

export class CabinetService {
  apm: Apm;
  constructor(apm: Apm) {
    this.apm = apm;
  }
}
