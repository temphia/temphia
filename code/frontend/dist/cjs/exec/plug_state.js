"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NewPlugStateApi = void 0;
const plug_state_1 = require("../apiv2/admin/plug_state");
exports.NewPlugStateApi = (api_base_url, token) => {
    return new plug_state_1.AdminPlugStateTktAPI(api_base_url, token);
};
