// syncme => xbprint/plug.go

export interface NewPlug {
  slug: string;
  name: string;
  agents: NewAgent[];
}

export interface NewAgent {
  name: string;
  type: string;
  executor: string;
  iface_file: string;
  entry_file: string;
  web_entry: string;
  web_script: string;
  web_style: string;
  web_loader: string;
  web_files: { [_: string]: string };

  resources: NewAgentResource[];
}

export interface NewAgentResource {
  name: string;
  type: string;
  ref_name: string;
  ref_data: any;
}
