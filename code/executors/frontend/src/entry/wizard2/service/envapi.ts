export class EnvApi {
  env: any;
  constructor(env: any) {
    this.env = env;
  }

  get_splash = async (has_exec_data: boolean) => {
    return this.env.PreformAction("get_splash", {
      has_exec_data,
    });
  };

  run_start = async (splash_data: any, start_raw_data: any) => {
    return this.env.PreformAction("run_start", {
      splash_data,
      start_raw_data,
    });
  };

  run_nested_start = async (
    parent_odata: string,
    field: string,
    start_raw_data: any
  ) => {
    return this.env.PreformAction("run_nested_start", {
      parent_odata,
      field,
      start_raw_data,
    });
  };

  run_back = async (odata: string) => {
    return this.env.PreformAction("run_back", {
      odata,
    });
  };

  run_next = async (odata: any, data: any) => {
    return this.env.PreformAction("run_next", {
      odata,
      data,
    });
  };

  get_source = async (odata: string, stage: string, source: string) => {
    return this.env.PreformAction("run_next", {
      odata,
      stage,
      source,
    });
  };
}
