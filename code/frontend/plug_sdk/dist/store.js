"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.getTemphiaExecKey = exports.TEMPHIA_EXEC_KEY = void 0;
exports.TEMPHIA_EXEC_KEY = "__temphia_exec_token__";
const getTemphiaExecKey = (redirect) => {
    try {
        const token = localStorage.getItem(exports.TEMPHIA_EXEC_KEY);
        if (!token) {
            if (redirect) {
                window.location.pathname = `/z/pages/agent/inject?redirect=${window.location}`;
            }
            return null;
        }
        return token;
    }
    catch (error) {
        return null;
    }
};
exports.getTemphiaExecKey = getTemphiaExecKey;
