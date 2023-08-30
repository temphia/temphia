import { Http } from "../http";

export class AdminPlugStateTktAPI {
  http: Http;
  token: string;
  api_base_url: string;
  constructor(api_base_url: string, token) {
    this.http = new Http(api_base_url, {
      "Content-Type": "application/json",
      Authorization: token,
    });

    this.token = token;
    this.api_base_url = api_base_url;
  }

  query(options: any) {
    return this.http.post(`/admin/plug_state/query`, options);
  }

  add(key: string, value: string, opts?: any) {
    return this.http.post(`/admin/plug_state/key`, {
      key,
      value,
      options: opts,
    });
  }

  update(key: string, value: string, opts?: any) {
    return this.http.post(`/admin/plug_state/key/${key}`, {
      key,
      value,
      options: opts,
    });
  }

  delete(key: string) {
    return this.http.delete(`/admin/plug_state/key/${key}`, {
      key,
    });
  }

  get(key: string) {
    return this.http.get(`/admin/plug_state/key/${key}`);
  }
}
