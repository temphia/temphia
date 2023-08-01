import type { XtMgr } from "../../../../services/portal/xtmgr";
import DataMigrator from "./data_migrator/index.svelte";
import DataSchema from "./data_schema/index.svelte";

export const registerEditors = (x: XtMgr) => {
  x.register_bprint_editor("data_migrator", DataMigrator);
  x.register_bprint_editor("data_schema", DataSchema);
};
