# Hacking 

This document describes how to hack codebase and overall code structure of the project.

## Frontend

Code in frontend lives in `code/frontend` folder, unlike project of similar which tend to have multiple files managed by pnpm/yarn workspace or whatever people have invented today, instead every entry (that generates .js file out after compilation) lives in entries directory. and has corresponding build command is found in package.json. most of those entries are svelte project so package json build pass own config/entry file as env variable. if it is not svelte entry then it is probably compiled by `ncc` which generates concated self contained js file which is perfect for libs.

When entry gets built, its output is in `public/build` folder there is a copy_files script that is responsible for copying to backend/data/assets/build so it could be included in static binary using golangs `embed` package. When you are developing go dev cli serves files directly from `public/build` but in 
not dev mode it fallbacks to that `embed` package files. so if you forgot to run copy files from time to time when you modify frontend code then non dev server might run in old build file then you build with in dev. 
