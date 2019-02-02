#!/usr/bin/env bash

protoc -I west4API-proto/ west4API-proto/*.proto  \
    --go_out=plugins=grpc:west4API \
    --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:../../siberia/app/src/west4API \
    --js_out=import_style=commonjs,binary:../../siberia/app/src/west4API \


case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        ls ../../siberia/app/src/west4API/* | while read line;do
        sed -i "" '1s/^/\/* eslint-disable *\/ /' $line
        done
     ;;
   Linux)
     echo 'Linux'
        ls ../../siberia/app/src/west4API/* | while read line;do
        sed -i '1s/^/\/* eslint-disable *\/ /' $line
        done
     ;;
   *)
     echo 'other OS' 
     ;;
esac




