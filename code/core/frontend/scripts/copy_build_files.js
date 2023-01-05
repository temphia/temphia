#!/usr/bin/node

var fs_Extra = require('fs-extra');
var path = require('path');

const BUILD_FOLDER = "public/build/";
const BUILD_ASSETS_FOLDER = "../backend/data/assets/build/"


const FILES = [
    "engine_iframe_guest.js",
    "auth.css",
    "auth.js",
    "auth.js.map",
    "playground.css",
    "playground.js",
    "playground.js.map",
    "portal.css",
    "portal.js",
    "portal.js.map",
    "start.css",
    "start.js",
    "start.js.map",
    "adapter_editor_loader.js"
]


FILES.forEach((file) => {

    const fromPath = path.join(BUILD_FOLDER, file)
    const toPath = path.join(BUILD_ASSETS_FOLDER, file)

    console.log(`copying file from ${fromPath} to ${toPath}`);

    fs_Extra.copy(fromPath, toPath, function (error) {
        if (error) {
            throw error;
        } else {
            console.log("success!");
        }
    });
})
