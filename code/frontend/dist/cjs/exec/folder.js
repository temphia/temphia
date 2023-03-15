"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NewFolderApi = void 0;
const cabinet_1 = require("../apiv2/cabinet");
exports.NewFolderApi = (api_base_url, token) => {
    return new cabinet_1.FolderTktAPI(api_base_url, token);
};
