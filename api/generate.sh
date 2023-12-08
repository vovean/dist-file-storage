#!/usr/bin/env bash

# this script mus be run from within its directory

buf generate --path protoc/;
mv protoc/*.go api