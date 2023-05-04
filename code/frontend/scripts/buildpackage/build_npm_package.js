#!/usr/bin/node

const { execSync } = require("child_process");

execSync("cp -f scripts/buildpackage/package.json package.json")

execSync("tsc -p scripts/tsconfig.package.json")