import type { Environment } from "../../../../lib/engine/environment";
import type { SheetExecData } from "../../../portal/pages/data/sheet/sheets";
import type { TableExecData } from "../../../portal/services/data";

export type ExecData = SheetExecData | TableExecData | null;

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
