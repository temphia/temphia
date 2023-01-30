import { Http } from "./http";

export class DataSheetAPI {
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

  list_sheet_group() {
    return this.http.post(`/data/sheet/list`, {});
  }

  load_sheet(sheetid: string, options: any) {
    return this.http.post(`/data/sheet/${sheetid}/load`, options);
  }

  // sheet

  list_sheets() {
    return this.http.get(`/data/sheet`);
  }

  get_sheet(sid: string) {
    return this.http.get(`/data/sheet/${sid}`);
  }

  new_sheet() {
    return this.http.get(`/data/sheet`);
  }

  update_sheet(sid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}`, data);
  }

  delete_sheet(sid: string) {
    return this.http.get(`/data/sheet/${sid}`);
  }

  // columns

  list_columns(sid: string) {
    return this.http.get(`/data/sheet/${sid}/column`);
  }

  get_column(sid: string, cid: string) {
    return this.http.get(`/data/sheet/${sid}/column/${cid}`);
  }

  new_column(sid: string, cid: string) {
    return this.http.get(`/data/sheet/${sid}/column/${cid}`);
  }

  update_column(sid: string, cid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/column/${cid}`, data);
  }

  delete_column(sid: string, cid: string) {
    return this.http.get(`/data/sheet/${sid}/column/${cid}`);
  }

  // row_cells

  get_row_cell(sid: string, rid: string) {
    return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
  }

  new_row_cell(sid: string, rid: string) {
    return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
  }

  update_row_cell(sid: string, rid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/row_cell/${rid}`, data);
  }

  delete_row_cell(sid: string, rid: string) {
    return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
  }
}
