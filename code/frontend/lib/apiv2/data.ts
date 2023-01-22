import { Http } from "./http";

export class DataAPI {
  http: Http;
  base_url: string;
  token: string;
  constructor(base_url: string, token: string) {
    this.http = new Http(base_url, {
      "Content-Type": "application/json",
      Authorization: token,
    });

    this.token = token;
    this.base_url = base_url;
  }

  load() {
    return this.http.get(`/data`);
  }

  new_row(tid: string, data: any) {
    return this.http.post(`/data/${tid}/row`, data);
  }

  get_row(tid: string, rid: string) {
    return this.http.get(`/data/${tid}/row/${rid}`);
  }

  update_row(tid: string, rid: string, data: any) {
    return this.http.post(`/data/${tid}/row/${rid}`, data);
  }

  delete_row(tid: string, rid: string) {
    return this.http.delete(`/data/${tid}/row/${rid}`);
  }

  load_table(tid: string, view?: string) {
    return this.http.post(`/data/${tid}/load`, {
      view,
    });
  }

  simple_query(tid: string, query: any) {
    return this.http.post(`/data/${tid}/simple_query`, query);
  }

  ref_load(tid: string, data: any) {
    return this.http.post(`/data/${tid}/ref_load`, data);
  }

  ref_resolve(tid: string, data: any) {
    return this.http.post(`/data/${tid}/ref_resolve`, data);
  }

  reverse_ref_load(tid: string, data: any) {
    return this.http.post(`/data/${tid}/rev_ref_load`, data);
  }
  list_activity(tid: string, rid: string) {
    return this.http.get(`/data/${tid}/activity/${rid}`);
  }

  comment_row(tid: string, rid: string, data: any) {
    return this.http.post(`/data/${tid}/activity/${rid}`, data);
  }

  sockd_url = () => {
    return `${this.base_url}/data_ws/?token=${this.token}`;
  };
}
