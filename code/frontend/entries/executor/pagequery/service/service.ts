import type { Environment } from "../../../../lib/engine/environment";
import type { SheetRow } from "../../../portal/pages/data/sheet/sheets";
import type {
  SheetColumn,
  SheetCell,
} from "../../../portal/pages/data/sheet/sheets";

export interface ExecData {
  invoker_type?: string;
  data_group?: string;
  sheet_id?: string;
  cells?: { [_: number]: { [_: string]: SheetCell } };
  columns?: SheetColumn[];
  rows: SheetRow[];
}

export class PageQueryService {
  env: Environment;
  exec_data: ExecData;
  constructor(env: Environment) {
    this.env = env;
    const execvars = env.GetExecVars();
    this.exec_data = execvars["exec_data"] || {};
  }

  load = () => {
    return this.env.PreformAction("load", {});
  };

  submit = (data: any) => {
    return this.env.PreformAction("submit", data);
  };
}
