#!/usr/bin/node

const fs = require('fs');
const { execSync } = require("child_process");

const build = execSync(`ncc build entries/executor_agent_bootloader/index.ts --out public/build/executor_agent_bootloader -t`);

console.log("BUILD OUTPUT", build.toString('utf-8'));

const data = fs.readFileSync('public/build/executor_agent_bootloader/index.js', 'utf8');
const finalData = `var __dirname = ''; var module = {}; module['exports']={};${data}`


fs.writeFile('public/build/executor_agent_bootloader.js', finalData, function (err) {
    if (err) return console.log("WRITE ERR", err);
    console.log('Write file done');

    const final = execSync(`rm -rf public/build/executor_agent_bootloader`);
    console.log("@final", final.toString('utf-8'));

});