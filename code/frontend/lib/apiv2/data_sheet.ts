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

  export(sheets: number[]) {
    return this.http.post(`/data/sheet/export`, sheets);
  }

  load_sheet(sheetid: string, options: any) {
    return this.http.post(`/data/sheet/${sheetid}/load`, options);
  }

  query_sheet(sheetid: string, options: any) {
    return this.http.post(`/data/sheet/${sheetid}/query`, options);
  }

  ref_query_sheet(sheetid: string, options: any) {
    return this.http.post(`/data/sheet/${sheetid}/ref_query`, options);
  }


  get_row_relation(
    sheetid: string,
    rid: string,
    refsheet: string,
    refcol: string
  ) {
    return this.http.get(
      `/data/sheet/${sheetid}/relation/${rid}/ref/${refsheet}/column/${refcol}`
    );
  }

  // sheet

  list_sheets() {
    return this.http.get(`/data/sheet`);
  }

  get_sheet(sid: string) {
    return this.http.get(`/data/sheet/${sid}`);
  }

  new_sheet(data: any) {
    return this.http.post(`/data/sheet`, data);
  }

  update_sheet(sid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}`, data);
  }

  delete_sheet(sid: string) {
    return this.http.delete(`/data/sheet/${sid}`);
  }

  // columns

  list_columns(sid: string) {
    return this.http.get(`/data/sheet/${sid}/column`);
  }

  get_column(sid: string, cid: string) {
    return this.http.get(`/data/sheet/${sid}/column/${cid}`);
  }

  new_column(sid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/column`, data);
  }

  update_column(sid: string, cid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/column/${cid}`, data);
  }

  delete_column(sid: string, cid: string) {
    return this.http.delete(`/data/sheet/${sid}/column/${cid}`);
  }

  // row_cells

  new_row_cell(sid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/row_cell`, data);
  }

  get_row_cell(sid: string, rid: string) {
    return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
  }

  update_row_cell(sid: string, rid: string, data: any) {
    return this.http.post(`/data/sheet/${sid}/row_cell/${rid}`, data);
  }

  delete_row_cell(sid: string, rid: string) {
    return this.http.delete(`/data/sheet/${sid}/row_cell/${rid}`);
  }

  list_users = (opts: any) => {
    return this.http.post(`/data/utils/user`, opts);
  };

  search(sheetid: string, opts: object) {
    return this.http.post(`/data/sheet/${sheetid}/search`, opts);
  }
}
