import type { PortalService } from "../../core";

export interface PlugRawSchema {
  slug: string;
  name: string;
  resource_hints: { [_: string]: ResourceHint };
  agent_hints: { [_: string]: AgentHint };
}

export interface AgentHint {
  name: string;
  type: string;
  executor: string;
  iface_file: string;
  web_entry: string;
  web_script: string;
  web_style: string;
  web_loader: string;
  web_files: { [_: string]: string };
  resources: { [_: string]: string };
}

export interface ResourceHint {
  name: string;
  type: string;
  sub_type: string;
}

export interface NewTableGroup {
  name: string;
  slug: string;
  description: string;
  tables: NewTable[];
  exec_order: string[];
}

export interface DataGroupRequest {
  dyndb_source: string;
  group_name: string;
  group_slug: string;
  cabinet_source: string;
  cabinet_folder: string;
  seed_source: string; // autogen | data
}

export interface NewTable {
  name: string;
  slug: string;
  description: string;
  icon: string;
  main_column?: string;
  activity_type: string;
  sync_type: string;
  columns: object[];
  indexes: object[];
  unique_indexes: object[];
  fts_index?: object;
  column_refs: object[];
  deleted_at: boolean;
  views: object[];
  seed_data?: object;
}

export interface PlugInstanceRequest {
  new_plug_id: string;
  new_plug_name: string;
  agent_opts: { [_: string]: AgentOptions };
  resources: { [_: string]: object };
}

export interface AgentOptions {
  name: string;
  resources: { [_: string]: string };
}

export class InstanceHelper {
  app: PortalService;
  constructor(app: PortalService) {
    this.app = app;
  }

  instance_plug = async (
    bid: string,
    file: string,
    data: PlugInstanceRequest
  ) => {
    const bapi = await this.app.api_manager.get_admin_bprint_api();
    return bapi.instance(bid, {
      bprint_id: bid,
      instancer_type: "plug",
      file,
      data,
    });
  };

  instance_data_group = async (
    bid: string,
    file: string,
    data: DataGroupRequest
  ) => {
    const bapi = await this.app.api_manager.get_admin_bprint_api();
    return bapi.instance(bid, {
      bprint_id: bid,
      instancer_type: "plug",
      file,
      data,
    });
  };
}

export const instance_helper = async (
  app: PortalService,
  btype: string,
  bprint: object,
  file: string,
  bundle_compo?: any
) => {
  console.log("@bprint", bprint);

  switch (btype) {
    case "tschema":
      console.log("@tschema");
      app.nav.admin_bprint_data_instancer(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "data_group":
      console.log("@data_group");
      app.nav.admin_bprint_data_instancer(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "data_table":
      console.log("@data_table");
      app.nav.admin_bprint_data_instancer(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "plug":
      console.log("@plug");
      app.nav.admin_bprint_plug_instancer(bprint["id"], file, bprint);
      break;
    case "app_bundle":
    case "bundle":
      if (!bundle_compo) return;

      console.log("@app_bundle");


      app.utils.small_modal_open(bundle_compo, {
        app,
        bid: bprint["id"],
        bprint,
      });
    default:
      break;
  }
};


/*
const sample_plug_schema = {
  "slug": "example1",
  "name": "Example",
  "resource_hints": {},
  "agent_hints": {
      "default": {
          "executor": "goja",
          "name": "default",
          "type": "web",
          "invoke_policy": "",
          "resources": {
              "extslot": ""
          },
          "web_files": {
              "client.js": "client.js",
              "client.css": "client.css"
          },
          "web_entry": "todo.helloworld.main",
          "web_script": "client.js",
          "web_style": "client.css",
          "web_loader": "std.loader"
      }
  }
}

*/
