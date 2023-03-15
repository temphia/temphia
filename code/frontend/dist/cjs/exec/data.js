"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NewDataTableApi = void 0;
const data_1 = require("../apiv2/data");
exports.NewDataTableApi = (api_base_url, token) => {
    return new data_1.DataAPI(api_base_url, token);
};
