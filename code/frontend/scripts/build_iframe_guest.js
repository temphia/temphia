#!/usr/bin/node

const fs = require('fs');
const { execSync } = require("child_process");

const build = execSync(`ncc build entries/execute_iframe_loader/index.ts --out public/build/guest_iframe`);

console.log("BUILD OUTPUT", build.toString('utf-8'));

const data = fs.readFileSync('public/build/guest_iframe/index.js', 'utf8');
const finalData = `var __dirname = ''; var module = {}; module['exports']={};${data}`


fs.writeFile('public/build/engine_iframe_guest.js', finalData, function (err) {
    if (err) return console.log("WRITE ERR", err);
    console.log('Write file done');

    const final = execSync(`rm -rf public/build/guest_iframe`);
    console.log("@final", final.toString('utf-8'));

});





