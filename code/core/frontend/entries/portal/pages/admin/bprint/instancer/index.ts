import type { PortalService } from "../../core";

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
