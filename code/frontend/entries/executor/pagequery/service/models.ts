export interface PgModel {
  title: string;
  stages: { [_: string]: any };
}

export interface QueryStage {
  script: string;
  about: string;
  data: { [_: string]: any };
  parameters: {[_: string]: Element};
}

export interface Element {
  name: string;
  type: string;
  info: string;
  view_opts: { [_: string]: any };
  data_opts: { [_: string]: any };
  source: string;
}
