#!/bin/bash

set -e
cd .. 

# All paths relative to project root
PB_SRCDIR=./nbproto
PB_SRC=nagbot.proto

# TODO: this should likely be moved into the Golang generate process.
echo "Generating protobuf code..."
protoc -I . "$PB_SRCDIR/$PB_SRC" --go_out=plugins=grpc:.

if type gnorm 2>&1 > /dev/null; then
    echo "Generating GNORM code..."
    (cd db/gnorm && gnorm gen)
else
    echo "Skipping GNORM generation, command not found."
fi
