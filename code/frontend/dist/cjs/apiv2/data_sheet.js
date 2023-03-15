"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataSheetAPI = void 0;
const http_1 = require("./http");
class DataSheetAPI {
    constructor(base_url, token) {
        this.list_users = (opts) => {
            return this.http.post(`/data/utils/user`, opts);
        };
        this.http = new http_1.Http(base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.base_url = base_url;
    }
    list_sheet_group() {
        return this.http.post(`/data/sheet/list`, {});
    }
    load_sheet(sheetid, options) {
        return this.http.post(`/data/sheet/${sheetid}/load`, options);
    }
    // sheet
    list_sheets() {
        return this.http.get(`/data/sheet`);
    }
    get_sheet(sid) {
        return this.http.get(`/data/sheet/${sid}`);
    }
    new_sheet(data) {
        return this.http.post(`/data/sheet`, data);
    }
    update_sheet(sid, data) {
        return this.http.post(`/data/sheet/${sid}`, data);
    }
    delete_sheet(sid) {
        return this.http.delete(`/data/sheet/${sid}`);
    }
    // columns
    list_columns(sid) {
        return this.http.get(`/data/sheet/${sid}/column`);
    }
    get_column(sid, cid) {
        return this.http.get(`/data/sheet/${sid}/column/${cid}`);
    }
    new_column(sid, data) {
        return this.http.post(`/data/sheet/${sid}/column`, data);
    }
    update_column(sid, cid, data) {
        return this.http.post(`/data/sheet/${sid}/column/${cid}`, data);
    }
    delete_column(sid, cid) {
        return this.http.delete(`/data/sheet/${sid}/column/${cid}`);
    }
    // row_cells
    new_row_cell(sid, data) {
        return this.http.post(`/data/sheet/${sid}/row_cell`, data);
    }
    get_row_cell(sid, rid) {
        return this.http.get(`/data/sheet/${sid}/row_cell/${rid}`);
    }
    update_row_cell(sid, rid, data) {
        return this.http.post(`/data/sheet/${sid}/row_cell/${rid}`, data);
    }
    delete_row_cell(sid, rid) {
        return this.http.delete(`/data/sheet/${sid}/row_cell/${rid}`);
    }
}
exports.DataSheetAPI = DataSheetAPI;
