#!/usr/bin/node

var fs_Extra = require('fs-extra');
var path = require('path');

const FILES = [
    ["public/build/engine_iframe_guest.js", "../backend/data/assets/build/engine_iframe_guest.js"]
]


FILES.forEach(([from, to]) => {

    console.log(`copying file from ${from} to ${to}`);

    fs_Extra.copy(from, to, function (error) {
        if (error) {
            throw error;
        } else {
            console.log("success!");
        }
    });
})
