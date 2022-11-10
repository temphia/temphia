import type { AdminDataAPI } from "../../../../lib/apiv2/admin";
import type { GroupService } from "./group";

export class DataService {
  sources: string[];
  admin_api: AdminDataAPI;
  current_group: GroupService | null;

  constructor(opts: { sources: string[]; admin_api: AdminDataAPI }) {
    this.sources = opts.sources;
    this.admin_api = opts.admin_api;
    this.current_group = null;
  }

  activate_group = async (source: string, group: string) => {
    if (!this.current_group) {
      return this.create_group(source, group);
    }

    if (
      this.current_group.source === source &&
      this.current_group.name === group
    ) {
      return;
    }
  };


  private create_group = async (source: string, group: string) => {};
}
