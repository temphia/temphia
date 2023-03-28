export interface LoadRequest {
  exec_data: any;
  version: string;
}

export interface LoadResponse {
  name: string;
  data: { [_: string]: any };
  sources: { [_: string]: Source };
  sections: Section[];
}

export interface Section {
  name: string;
  layout: string;
  panels: Panel[];
  options: { [_: string]: any };
}

export interface Panel {
  name: string;
  width: string;
  height: number;
  interval: number;
  type: string;
  options: { [_: string]: any };
  source: string;
}

export interface Source {
  name: string;
  type: string;
  options: { [_: string]: any };
}
