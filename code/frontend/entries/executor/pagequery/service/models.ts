export interface PgModel {
  title: string;
  stages: { [_: string]: any };
  first_stage: string;
}

export interface QueryStage {
  script: string;
  about: string;
}

export interface Element {
  name: string;
  type: string;
  info: string;
  view_opts: { [_: string]: any };
  data_opts: { [_: string]: any };
  source: string;
}

export interface LoadRequest {
  exec_data: { [_: string]: any };
}

export interface LoadResponse {
  title: string;
  stages: { [_: string]: QueryStage };
  first_stage: string;
}

export interface SubmitRequest {
  stage: string;
  param_data: { [_: string]: any };
  script: string;
}

export interface SubmitResponse {
  stage: string;
  data: { [_: string]: any };
  elements: Element[];
}


