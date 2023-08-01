import type { PortalService } from "../../core";
import type { DataGroupRequest } from "./xmodels/data";

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
      app.nav.admin_bprint_data_instancer(bprint["id"], file, bprint);
      break;
    case "data_group":
      console.log("@data_group");
      app.nav.admin_bprint_data_instancer(bprint["id"], file, bprint);
      break;
    case "data_table":
      console.log("@data_table");
      app.nav.admin_bprint_data_instancer(bprint["id"], file, bprint);
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
