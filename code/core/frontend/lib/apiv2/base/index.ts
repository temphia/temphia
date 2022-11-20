import { apiURL } from "../../utils/site";
import { Http } from "../http";

export class ApiBase {
  http: Http;
  base_url: string;
  tenant_id: string;
  user_token: string;

  constructor(base_url: string, tenant_id: string, token: string) {
    this.base_url = base_url;
    this.tenant_id = tenant_id;
    this.user_token = token;

    this.http = new Http(base_url, {
      "Content-Type": "application/json",
      Authorization: token,
    });
  }

  async init() {
    const resp = await fetch(`${apiURL(this.tenant_id)}/auth/refresh`, {
      method: "POST",
      body: JSON.stringify({
        path: ["basic"],
        user_token: this.user_token,
      }),
    });
    
    const rdata = await resp.json();
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
