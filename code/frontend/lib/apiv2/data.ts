import { DataSheetAPI } from "./data_sheet";
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
    return this.http.post(`/data/table/${tid}/row`, data);
  }

  get_row(tid: string, rid: string) {
    return this.http.get(`/data/table/${tid}/row/${rid}`);
  }

  update_row(tid: string, rid: string, data: any) {
    return this.http.post(`/data/table/${tid}/row/${rid}`, data);
  }

  delete_row(tid: string, rid: string) {
    return this.http.delete(`/data/table/${tid}/row/${rid}`);
  }

  load_table(tid: string, view?: string) {
    return this.http.post(`/data/table/${tid}/load`, {
      view,
    });
  }

  simple_query(tid: string, query: any) {
    return this.http.post(`/data/table/${tid}/simple_query`, query);
  }

  ref_load(tid: string, data: any) {
    return this.http.post(`/data/table/${tid}/ref_load`, data);
  }

  ref_resolve(tid: string, data: any) {
    return this.http.post(`/data/table/${tid}/ref_resolve`, data);
  }

  reverse_ref_load(tid: string, data: any) {
    return this.http.post(`/data/table/${tid}/rev_ref_load`, data);
  }
  list_activity(tid: string, rid: string) {
    return this.http.get(`/data/table/${tid}/activity/${rid}`);
  }

  comment_row(tid: string, rid: string, data: any) {
    return this.http.post(`/data/table/${tid}/activity/${rid}`, data);
  }

  sockd_url = () => {
    return `${this.base_url}/data_ws/?token=${this.token}`;
  };

  sheet_api = () => {
    return new DataSheetAPI(this.base_url, this.token);
  };
}
