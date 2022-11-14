import type { AdminBprintAPI, AdminPlugAPI } from "../../../../lib/apiv2";
import type { PlugDevTktAPI } from "../../../../lib/apiv2/plug_dev_tkt";
import type { ApiManager } from "../portal/apm";

export class DevShellService {
  apm: ApiManager;
  pid: string;
  aid: string;

  plug?: object;
  agent?: string;

  dev_api: PlugDevTktAPI;
  bprint_api: AdminBprintAPI;
  plug_api: AdminPlugAPI;

  constructor(apm: ApiManager, pid: string, aid: string) {
    this.apm = apm;
    this.pid = pid;
    this.aid = aid;
  }

  async init(): Promise<true | string> {
    this.plug_api = this.apm.get_admin_plug_api();
    this.bprint_api = this.apm.get_admin_bprint_api();

    const presp = await this.plug_api.get_plug(this.pid);
    if (!presp.ok) {
      console.log("err", presp);
      return presp.data;
    }

    this.plug = presp.data;
    const resp = await this.bprint_api.issue(this.plug["bprint_id"], {
      all_plugs: false,
      plug_ids: [this.pid],
    });

    if (!resp.ok) {
      console.log("err", resp);
      return resp.data;
    }

    this.dev_api = this.apm.get_dev_plug_tkt_api(resp.data["dev_ticket"]);
    const aresp = await this.plug_api.get_agent(this.pid, this.aid);
    if (!aresp.ok) {
      return resp.data;
    }

    return true;
  }

  async load_interface() {
    const ifaceUrl: string = this.agent["iface_file"];
    if (!ifaceUrl) {
      console.log("Empty openrpc iface file");
      return;
    }

    let iface_data;
    let loaded = false;

    if (ifaceUrl.startsWith("http://") || ifaceUrl.startsWith("https://")) {
      const bresp = await fetch(ifaceUrl);
      iface_data = await bresp.json();
      loaded = true;
    } else if (ifaceUrl.startsWith("@bprint/")) {
      const bpath = ifaceUrl.replace("@bprint/", "").split("/");
      const bresp = await this.bprint_api.get_file(bpath[0], bpath[1]);
      if (bresp.status !== 200) {
        console.log("connot retrive", bresp);
        return;
      }

      iface_data =
        typeof bresp.data === "object"
          ? bresp.data
          : JSON.stringify(bresp.data);

      loaded = true;
    } else {
      const bresp = await this.bprint_api.get_file(
        this.plug["bprint_id"],
        ifaceUrl
      );
      if (bresp.status !== 200) {
        console.log("connot retrive", bresp);
        return;
      }

      iface_data =
        typeof bresp.data === "object"
          ? bresp.data
          : JSON.stringify(bresp.data);
      loaded = true;
    }
  }
}
