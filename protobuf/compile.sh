#!/usr/bin/env bash

protoc -I proto/ proto/enum.proto proto/api.proto proto/webapi.proto  \
    --go_out=plugins=grpc:api \
    --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:../../siberia/app/src/api \
    --js_out=import_style=commonjs,binary:../../siberia/app/src/api \


case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        sed -i "" '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/webapi_pb.js
        sed -i "" '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/api_pb_service.js
        sed -i "" '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/enum_pb.js
        sed -i "" '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/api_pb.js
     ;;
   Linux)
     echo 'Linux'
        sed -i '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/webapi_pb.js
        sed -i '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/api_pb_service.js
        sed -i '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/enum_pb.js
        sed -i '1s/^/\/* eslint-disable *\/ /' ../../siberia/app/src/api/api_pb.js
     ;;
   *)
     echo 'other OS' 
     ;;
esac




