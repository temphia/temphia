// syncme => vmodels/plug

export interface ExecInstanceOptions {
  api_base_url: string;
  token: string;
  entry: string;
  exec_loader: string;
  js_plug_script: string;
  style: string;
  ext_scripts?: { [_: string]: string };
  plug: string;
  agent: string;
}
