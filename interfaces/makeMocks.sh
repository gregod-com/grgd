#!/bin/sh

for file in *.go; do
echo $file
mockgen --source=$file -destination mocks/mock$file -package mocks
done

for file in ../../grgdplugincontracts/*.go; do
echo $file
mockgen --source=$file -destination mocks/mock$(basename $file) -package mocks
done

