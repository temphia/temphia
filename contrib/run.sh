#!/bin/sh

echo 'Hello, World'

echo $PWD

BINARY="${PWD}/temphia"

mkdir -p ~/.local/share/temphia

cd ~/.local/share/temphia

exec $BINARY