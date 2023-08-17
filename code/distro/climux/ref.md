# REFERENCE

## data folders
- files
- logs
- db/sqlite


## temphia_package
- .bin/{nsjail, bun, webview.dll}
- temphia.desktop
- temphia_desktop.sh
- temphia_server.sh
- demo_assets
- temphia.service

## subcli
- bdev
- app {init-data[demo_seed?], init-db , start, actual-start}
- repobuild
- database {ping migrate rollback}
- desktop(local://tmp/temphia.sock, lpweb://<hash>.lpweb )
- lpweb
- log {watch["folder"]}
- remote_runner