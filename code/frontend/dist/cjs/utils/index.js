"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.toGeoJson = exports.fromGeoJson = exports.fromGeoJsonOrFallback = exports.isImage = exports.humanizeBytes = exports.validateEmail = exports.validateUserId = exports.validateSlug = exports.hslColor = exports.numHash = exports.strHash = exports.generateId = void 0;
exports.generateId = () => Math.random().toString(36).slice(2);
exports.strHash = (str) => {
    let hash = 0;
    for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i);
        hash = (hash << 5) - hash + char;
        hash &= hash; // Convert to 32bit integer
    }
    return new Uint32Array([hash])[0].toString(36);
};
const pp = ".*(D#D01e-u0_ue819g_!UJ123456789023";
exports.numHash = (str) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
        hash = str.charCodeAt(i) + ((hash << 6) - hash);
        hash = pp.charCodeAt(i) ^ hash;
    }
    return hash;
};
exports.hslColor = (str) => {
    return `background: hsl(${exports.numHash(str) % 360}, 100%, 80%)`;
};
exports.validateSlug = (v) => /^[a-z](-?[a-z])*$/.test(v);
exports.validateUserId = (v) => /^[a-z]+([a-z0-9_])+/.test(v);
exports.validateEmail = (v) => /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(v);
const units = ["bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];
// makes bytes value nice
exports.humanizeBytes = (x) => {
    let l = 0, n = parseInt(x, 10) || 0;
    while (n >= 1024 && ++l) {
        n = n / 1024;
    }
    return n.toFixed(n < 10 && l > 0 ? 1 : 0) + " " + units[l];
};
const imageTypes = ["png", "jpg", "jpeg"];
exports.isImage = (name) => {
    const frags = name.split(".");
    return imageTypes.includes(frags[frags.length - 1]);
};
exports.fromGeoJsonOrFallback = (jstr) => {
    const fallback = [27.7116, 85.3124];
    try {
        const jpoint = JSON.parse(jstr);
        return jpoint["coordinates"] || fallback;
    }
    catch (error) {
        return fallback;
    }
};
exports.fromGeoJson = (jstr) => {
    try {
        const jpoint = JSON.parse(jstr);
        return jpoint["coordinates"] || [0, 0];
    }
    catch (error) {
        return [0, 0];
    }
};
exports.toGeoJson = (_lat, _lon) => {
    return JSON.stringify({
        type: "Point",
        coordinates: [_lat, _lon],
    });
};
