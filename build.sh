#!/bin/sh

# check go version > 1.12

# TODO: check installed go

# TODO: check go version > 1.12

cd cmd
go build -o ../hade ./

echo "Build success. "
echo "You can find excutable command named \"hade\". "
echo "Please use ./hade"