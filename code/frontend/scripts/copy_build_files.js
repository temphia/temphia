#!/usr/bin/node

var fs_Extra = require('fs-extra');
var path = require('path');

const BUILD_FOLDER = "public/build/";
const BUILD_ASSETS_FOLDER = "../backend/data/assets/build/"


const FILES = [
    "auth.css",
    "auth.js",
    "auth.js.map",
    // "playground.css",
    // "playground.js",
    // "playground.js.map",
    "portal.css",
    "portal.js",
    "portal.js.map",
    // "start.css",
    // "start.js",
    // "start.js.map",
    "operator.css",
    "operator.js",
    "operator.js.map",

    
    "adapter_editor_bootloader.js",
    "adapter_editor_noop.js",
    "adapter_editor_easypage.js",
    "adapter_editor_easypage.css",

    "executor_bootloader_iframe.js",
    "executor_pageform.js",
    "executor_pageform.css",
    "executor_pagedash.js",
    "executor_pagedash.css",
    "executor_pagequery.js",
    "executor_pagequery.css"
]

fs_Extra.readdirSync(BUILD_FOLDER).forEach((file) => {
    if (!FILES.includes(file)) {
        console.log("@file not in copied list =>", file)
    }
})

fs_Extra.readdirSync(BUILD_ASSETS_FOLDER).forEach((file) => {
    if (!FILES.includes(file)) {
        console.log("@file not from build folder =>", file)
    }
})



FILES.forEach((file) => {

    const fromPath = path.join(BUILD_FOLDER, file)
    const toPath = path.join(BUILD_ASSETS_FOLDER, file)

    console.log(`copying file from ${fromPath} to ${toPath}`);

    fs_Extra.copy(fromPath, toPath, function (error) {
        if (error) {
            console.log(`@err [${file}] =>`, error)
            throw error;
        } else {
            console.log("@success copy", file);
        }
    });
})
