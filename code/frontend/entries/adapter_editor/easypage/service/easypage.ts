import type { AdapterEditorEnv } from "../../../../lib/adapter/adapter";
import { AdminPlugStateTktAPI } from "../../../../lib/apiv2";

export class EasypageService {
  env: AdapterEditorEnv;
  papi: AdminPlugStateTktAPI;

  modal: {
    big_open: any;
    big_close: any;
    small_open: any;
    small_close: any;
  };

  constructor(env: AdapterEditorEnv) {
    this.env = env;
  }

  load = async () => {
    const resp = await this.env.api.perform_action("load", {});
    if (!resp.ok) {
      return resp.data;
    }
    this.papi = new AdminPlugStateTktAPI(
      this.env.api.base.api_base_url,
      resp.data
    );

    return null;
  };

  loadPages = async () => {
    const resp = await this.papi.query({
      tag1: "page",
    });

    return resp.data.map((elem) => formatData(elem));
  };

  addPage = async (slug: string, data: any) => {
    const resp = await this.papi.add(`page-${slug}`, JSON.stringify(data), {
      tag1: "page",
    });
    if (!resp.ok) {
      return resp;
    }

    const resp1 = await this.papi.add(`pdata-${slug}`, JSON.stringify(data), {
      tag1: "pdata",
    });
    if (!resp1.ok) {
      return resp1;
    }

    return resp1;
  };

  getPageData = (slug: string) => {
    return this.papi.get(`pdata-${slug}`);
  };

  setPageData = (slug: string, data: string) => {
    return this.papi.add(`pdata-${slug}`, JSON.stringify(data));
  };

  deletePage = async (slug: string) => {
    await this.papi.delete(`page-${slug}`);
    await this.papi.delete(`pdata-${slug}`);
  };
}

export const formatData = (data) => {
  const slug = data["slug"];
  const value = JSON.parse(data["value"] || "{}");

  return { slug, ...value };
};
