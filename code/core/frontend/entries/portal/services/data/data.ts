import type { DataAPI, SelfAPI } from "../../../../lib/apiv2";
import type { AdminDataAPI } from "../../../../lib/apiv2/admin";
import type { SockdService } from "../sockd/sockd";
import { GroupService } from "./group";

interface Apm {
  get_admin_data_api(): AdminDataAPI;
  get_data_api(source: string, group: string): Promise<DataAPI>;
}

export class DataService {
  sources: string[];
  current_group: GroupService | null;
  sockd_builder: SockdService;
  apm: Apm;
  close_modal: any;
  open_modal: any;
  api_base_url: string;

  constructor(opts: {
    sources: string[];
    apm: Apm;
    api_base_url: string;
    sockd_builder: SockdService;
    close_modal: any;
    open_modal: any;
  }) {
    this.sources = opts.sources;
    this.current_group = null;
    this.apm = opts.apm;
    this.api_base_url = opts.api_base_url;
    this.close_modal = opts.close_modal;
    this.open_modal = opts.open_modal;
    this.sockd_builder = opts.sockd_builder;
  }

  group_service = async (source: string, group: string) => {
    if (!this.current_group) {
      return this.create_group(source, group);
    }

    if (
      this.current_group.source === source &&
      this.current_group.name === group
    ) {
      return this.current_group;
    }
    await this.current_group.close();
    return this.create_group(source, group);
  };

  private create_group = async (source: string, group: string) => {
    const data_api = await this.apm.get_data_api(source, group);
    if (!data_api) {
      console.log("BIG ERR");
      return;
    }

    const group_svc = new GroupService({
      api: data_api,
      name: group,
      source: source,
      close_modal: this.close_modal,
      open_modal: this.open_modal,
      api_base_url: this.api_base_url,
      sockd_builder: this.sockd_builder,
    });

    await group_svc.init();
    return group_svc;
  };
}
