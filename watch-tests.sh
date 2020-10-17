#!/bin/zsh

fswatch -o ./* | xargs -n1 -I{} bash -c 'clear && go test -v ./...'

