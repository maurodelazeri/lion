#!/usr/bin/env bash
rm  -rf ../../siberia/app/src/heraldsquareAPI/*
rm -rf heraldsquareAPI/*
 
protoc -I heraldsquareAPI-proto/ heraldsquareAPI-proto/*.proto  \
    --go_out=plugins=grpc:heraldsquareAPI \
    --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
    --ts_out=service=true:../../siberia/app/src/heraldsquareAPI \
    --js_out=import_style=commonjs,binary:../../siberia/app/src/heraldsquareAPI \


case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        ls ../../siberia/app/src/heraldsquareAPI/* | while read line;do
        sed -i "" '1s/^/\/* eslint-disable *\/ /' $line
        sed -i "" 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        done
     ;;
   Linux)
     echo 'Linux'
        ls ../../siberia/app/src/heraldsquareAPI/* | while read line;do
        sed -i '1s/^/\/* eslint-disable *\/ /' $line
        sed -i 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        done
     ;;
   *)
     echo 'other OS' 
     ;;
esac