"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.actionFetch = void 0;
exports.actionFetch = (actionUrl, token) => async (name, data) => {
    const response = await fetch(`${actionUrl}/${name}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: token,
        },
        redirect: "follow",
        referrerPolicy: "strict-origin-when-cross-origin",
        body: data,
    });
    return response;
};
