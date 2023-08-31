import type { DataAPI, SelfAPI } from "../apiv2";
import type { AdminDataAPI } from "../apiv2/admin";
import type { SockdBuilder } from "../portal/sockd/builder";
import { GroupService } from "./table/group";
import { SheetGroupService } from "./sheet/sheet";

interface Apm {
  get_admin_data_api(): AdminDataAPI;
  get_data_api(source: string, group: string): Promise<DataAPI>;
}

export class DataService {
  sources: string[];
  current_group: GroupService | null;
  sockd_builder: SockdBuilder;
  apm: Apm;
  close_modal: any;
  open_modal: any;
  api_base_url: string;

  current_sheet_group: SheetGroupService;
  profile_genrator: (user: string) => string;

  constructor(opts: {
    sources: string[];
    apm: Apm;
    api_base_url: string;
    sockd_builder: SockdBuilder;
    close_modal: any;
    open_modal: any;
    profile_genrator: (string) => string;
  }) {
    this.sources = opts.sources;
    this.current_group = null;
    this.apm = opts.apm;
    this.api_base_url = opts.api_base_url;
    this.close_modal = opts.close_modal;
    this.open_modal = opts.open_modal;
    this.sockd_builder = opts.sockd_builder;
    this.profile_genrator = opts.profile_genrator;
  }

  group_service = async (source: string, group: string) => {
    if (!this.current_group) {
      this.current_group = await this.create_group(source, group);
      return this.current_group;
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

  group_sheet = async (source: string, group: string) => {
    if (!this.current_sheet_group) {
      return this.create_group_sheet(source, group);
    }

    if (
      this.current_sheet_group.source === source &&
      this.current_sheet_group.group_slug
    ) {
      return this.current_sheet_group;
    }

    return this.create_group_sheet(source, group);
  };

  private create_group_sheet = async (source: string, group: string) => {
    const dapi = await this.apm.get_data_api(source, group);
    const sgs = new SheetGroupService(
      source,
      group,
      dapi,
      this.sockd_builder,
      this.profile_genrator
    );
    await sgs.init();
    this.current_sheet_group = sgs;
    return sgs;
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
      profile_generator: this.profile_genrator,
    });

    await group_svc.init();
    return group_svc;
  };
}
