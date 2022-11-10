import { Http } from "./http";

export class DataAPI {
  http: Http;
  constructor(base_url: string, token: string) {
    this.http = new Http(base_url, {
      "Content-Type": "application/json",
      Authorization: token,
    });
  }

  load() {
    return this.http.get(`/data`);
  }

  newRow(tid: string) {
    return this.http.get(`/data/${tid}/row`);
  }

  getRow(tid: string, rid: string) {
    return this.http.get(`/data/${tid}/row/${rid}`);
  }

  updateRow(tid: string, rid: string, data: any) {
    return this.http.post(`/data/${tid}/row/${rid}`, data);
  }

  deleteRow(tid: string, rid: string) {
    return this.http.delete(`/data/${tid}/row/${rid}`);
  }

  simpleQuery(tid: string, query: any) {
    return this.http.post(`/data/${tid}/simple_query`, query);
  }

  refLoad(tid: string, data: any) {
    return this.http.post(`/data/${tid}/ref_load`, data);
  }

  refResolve(tid: string, data: any) {
    return this.http.post(`/data/${tid}/ref_resolve`, data);
  }

  reverseRefLoad(tid: string, data: any) {
    return this.http.post(`/data/${tid}/rev_ref_load`, data);
  }
  listActivity(tid: string, rid: string) {
    return this.http.get(`/data/${tid}/activity${rid}`);
  }

  commentRow(tid: string, rid: string, data: any) {
    return this.http.post(`/data/${tid}/activity${rid}`, data);
  }
}
