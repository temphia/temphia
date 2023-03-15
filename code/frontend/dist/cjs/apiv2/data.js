"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataAPI = void 0;
const data_sheet_1 = require("./data_sheet");
const http_1 = require("./http");
class DataAPI {
    constructor(api_base_url, token) {
        this.sockd_url = () => {
            return `${this.api_base_url}/data_ws/?token=${this.token}`;
        };
        this.sheet_api = () => {
            return new data_sheet_1.DataSheetAPI(this.api_base_url, this.token);
        };
        this.list_users = (opts) => {
            return this.http.post(`/data/utils/user`, opts);
        };
        this.http = new http_1.Http(api_base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.api_base_url = api_base_url;
    }
    load() {
        return this.http.get(`/data`);
    }
    new_row(tid, data) {
        return this.http.post(`/data/table/${tid}/row`, data);
    }
    get_row(tid, rid) {
        return this.http.get(`/data/table/${tid}/row/${rid}`);
    }
    update_row(tid, rid, data) {
        return this.http.post(`/data/table/${tid}/row/${rid}`, data);
    }
    delete_row(tid, rid) {
        return this.http.delete(`/data/table/${tid}/row/${rid}`);
    }
    load_table(tid, view) {
        return this.http.post(`/data/table/${tid}/load`, {
            view,
        });
    }
    simple_query(tid, query) {
        return this.http.post(`/data/table/${tid}/simple_query`, query);
    }
    ref_load(tid, data) {
        return this.http.post(`/data/table/${tid}/ref_load`, data);
    }
    ref_resolve(tid, data) {
        return this.http.post(`/data/table/${tid}/ref_resolve`, data);
    }
    reverse_ref_load(tid, data) {
        return this.http.post(`/data/table/${tid}/rev_ref_load`, data);
    }
    list_activity(tid, rid) {
        return this.http.get(`/data/table/${tid}/activity/${rid}`);
    }
    comment_row(tid, rid, data) {
        return this.http.post(`/data/table/${tid}/activity/${rid}`, data);
    }
}
exports.DataAPI = DataAPI;
