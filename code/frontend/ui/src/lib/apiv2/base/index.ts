import { apiURL } from "../../utils/site";
import { Http } from "../http";

export class ApiBase {
  http: Http;
  api_base_url: string;
  tenant_id: string;
  user_token: string;

  constructor(api_base_url: string, tenant_id: string, token: string) {
    this.api_base_url = api_base_url;
    this.tenant_id = tenant_id;
    this.user_token = token;

    console.log("@api_base =>", this);

    this.http = new Http(api_base_url, {
      "Content-Type": "application/json",
      Authorization: token,
    });
  }

  async init() {
    // fixme => it should not create apiurl instead it should be passed
    // base_url or sth and build api_base_url from it
    const resp = await fetch(`${apiURL(this.tenant_id)}/auth/refresh`, {
      method: "POST",
      body: JSON.stringify({
        path: ["basic"],
        user_token: this.user_token,
      }),
    });

    const rdata = await resp.json();
  }

  async raw_fetch(
    path: string,
    method: string,
    auth: boolean,
    build_path: boolean,
    data: any
  ) {
    return this.http.rawFetch(path, method, auth, build_path, data);
  }

  async get(path: string) {
    return this.http.get(path);
  }

  async post(path: string, data: any) {
    return this.http.post(path, data);
  }

  async postForm(path: string, auth: boolean, data: any) {
    return this.http.postForm(path, auth, data);
  }

  async patchForm(path: string, auth: boolean, data: any) {
    return this.http.patchForm(path, auth, data);
  }

  async put(path: string, data: any) {
    return this.http.put(path, data);
  }

  async patch(path: string, data: any) {
    return this.http.patch(path, data);
  }

  async delete(path: string, data?: any) {
    return this.http.delete(path, data);
  }
}
