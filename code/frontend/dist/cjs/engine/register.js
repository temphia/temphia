"use strict";
// register factor helper functions 
Object.defineProperty(exports, "__esModule", { value: true });
exports.registerFactory = exports.registerExecLoaderFactory = exports.registerPlugFactory = void 0;
exports.registerPlugFactory = (entryName, factory) => exports.registerFactory("plug.factory", entryName, factory);
exports.registerExecLoaderFactory = (name, factory) => exports.registerFactory("loader.factory", name, factory);
exports.registerFactory = (ftype, name, factory) => {
    const pf = window["__register_factory__"];
    if (!pf) {
        console.warn("factory registry not found");
        return;
    }
    pf(ftype, name, factory);
};
