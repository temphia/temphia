"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExecAM = void 0;
const data_1 = require("./data");
const folder_1 = require("./folder");
const plug_state_1 = require("./plug_state");
const sockd_1 = require("./sockd");
// ExecAM stands for execution api manager
class ExecAM {
    constructor(api_base_url) {
        this.new_data_api = (token) => {
            return data_1.NewDataTableApi(this.api_base_url, token);
        };
        this.new_folder_api = (token) => {
            return folder_1.NewFolderApi(this.api_base_url, token);
        };
        this.new_sockd_room = async (token) => {
            return sockd_1.NewSockdRoom(`${this.api_base_url}/engine/ws?ticket=${token}`);
        };
        this.new_plug_state = (token) => {
            return plug_state_1.NewPlugStateApi(this.api_base_url, token);
        };
        this.api_base_url = api_base_url;
    }
}
exports.ExecAM = ExecAM;
