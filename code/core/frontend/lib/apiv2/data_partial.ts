export interface PartialDataOptions {
  tkt: string;
  type: "QUERY" | "WRITE";
}

export class PartialDataAPI {
  tkt: string;
  constructor(opts: PartialDataOptions) {}
}