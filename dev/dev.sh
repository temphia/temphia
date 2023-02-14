#! /usr/bin/bash

if [[ -f "go.mod" ]]
then
    cd dev
fi


tmux new "echo database; make docker_run" ';'  \
    split -h "echo server_backend; make server_run" ';'  \
    split "cd ../code/frontend && bash" ';' \
    select-pane -L ';'  \
    split "bash" ';' \