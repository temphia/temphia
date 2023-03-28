export interface LoadRequest {
  data_context_type: string;
  rows: number[];
  options: { [_: string]: any };
}

export interface Response {
  ok: boolean;
  final: boolean;
  message: string;
  items: {
    name: string;
    type: string;
    info: string;
    html_attr: { [_: string]: string };
    options: string[];
  }[];
  data: any;
  on_load: string;
  on_submit: string;
  stage: string;
}
