import type { PortalApp } from "../../../../../lib/app/portal";

export const instance = async (
  app: PortalApp,
  btype: string,
  bprint: object,
  file: string,
  bundle_compo?: any
) => {
  console.log("@bprint", bprint);

  switch (btype) {
    case "tschema":
      console.log("@tschema");
      app.navigator.goto_admin_bprint_data_group_instance(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "data_group":
      console.log("@data_group");
      app.navigator.goto_admin_bprint_data_group_instance(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "data_table":
      console.log("@data_table");
      app.navigator.goto_admin_bprint_data_table_instance(
        bprint["id"],
        file,
        bprint
      );
      break;
    case "plug":
      console.log("@plug");
      app.navigator.goto_admin_bprint_plug_instance(bprint["id"], file, bprint);
      break;
    case "app_bundle":
    case "bundle":
      if (!bundle_compo) return;

      console.log("@app_bundle");
      app.simple_modal_open(bundle_compo, {
        app,
        bid: bprint["id"],
        bprint,
      });
    default:
      break;
  }
};
