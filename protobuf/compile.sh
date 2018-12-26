#!/usr/bin/env bash

protoc -I proto/ proto/*.proto  \
    --go_out=plugins=grpc:api \
    --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:../../siberia/app/src/api \
    --js_out=import_style=commonjs,binary:../../siberia/app/src/api \


case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        ls ../../siberia/app/src/api/* | while read line;do
        sed -i "" '1s/^/\/* eslint-disable *\/ /' $line
        done
     ;;
   Linux)
     echo 'Linux'
        ls ../../siberia/app/src/api/* | while read line;do
        sed -i '1s/^/\/* eslint-disable *\/ /' $line
        done
     ;;
   *)
     echo 'other OS' 
     ;;
esac




