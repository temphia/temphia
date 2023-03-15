"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NewSockdRoom = void 0;
const sockd_1 = require("../sockd");
exports.NewSockdRoom = async (url) => {
    const sockd = new sockd_1.Sockd(url);
    await sockd.Init();
    return sockd;
};
