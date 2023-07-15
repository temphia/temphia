# Tasks

xc -f tasks.md

## install_cli
```sh
cp ../build/temphia-cli ~/.bin
```

## backup_cli
```sh
cp ../build/temphia-cli ~/Desktop/temphia_run
```

## list_tests

```sh
cd ../cmd/tests && go test -v *.go  -args list_tests
```

## run_tests

```sh
cd ../cmd/tests && go test -v *.go
```

## run_some_tests

Inputs: TESTNAME

```sh
cd ../cmd/tests && go test -v *.go -run $TESTNAME
```
