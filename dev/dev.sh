#! /usr/bin/bash

if [[ -f "go.mod" ]]
then
    cd dev
fi


tmux new "echo database; make backend_docker_run" ';'  \
    split -h "echo server_backend; make backend_server_run" ';'  \
    split "cd ../code/frontend && bash" ';' \
    select-pane -L ';'  \
    split "bash" ';' \