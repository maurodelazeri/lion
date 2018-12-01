#!/usr/bin/env bash
protoc -I proto/ proto/enum.proto proto/webapi.proto  \
    --go_out=plugins=grpc:../../alaska/api \
    --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:../../siberia/app/src/api \
    --js_out=import_style=commonjs,binary:../../siberia/app/src/api \