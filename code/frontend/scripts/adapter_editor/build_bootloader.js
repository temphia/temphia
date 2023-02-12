#!/usr/bin/node

const fs = require('fs');
const { execSync } = require("child_process");

const build = execSync(`ncc build entries/adapter_editor/@bootloader/index.ts --out public/build/adapter_editor_bootloader -t`);

console.log(build.toString('utf-8'));

const data = fs.readFileSync('public/build/adapter_editor_bootloader/index.js', 'utf8');
const finalData = `var __dirname = ''; var module = {}; module['exports']={};${data}`


fs.writeFile('public/build/adapter_editor_bootloader.js', finalData, function (err) {
    if (err) return console.log(err);
    console.log('Write file done');

    const final = execSync(`rm -rf public/build/adapter_editor_bootloader`);
    console.log("@final", final.toString('utf-8'));
});

