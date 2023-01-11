#!/usr/bin/node

const fs = require('fs');
const { execSync } = require("child_process");

const build = execSync(`ncc build entries/adapter_editor_loader/index.ts --out public/build/adapter_editor_loader -t`);

console.log(build.toString('utf-8'));

const data = fs.readFileSync('public/build/adapter_editor_loader/index.js', 'utf8');
const finalData = `var __dirname = ''; var module = {}; module['exports']={};${data}`


fs.writeFile('public/build/adapter_editor_loader.js', finalData, function (err) {
    if (err) return console.log(err);
    console.log('Write file done');

    const final = execSync(`rm -rf public/build/adapter_editor_loader`);
    console.log("@final", final.toString('utf-8'));

});

