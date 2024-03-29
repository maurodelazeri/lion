#!/usr/bin/env bash
rm  -rf ../../siberia/app/src/api/*
rm -rf api/*
 
protoc -I proto/ proto/*.proto  --js_out=import_style=commonjs:../../siberia/app/src/api/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../../siberia/app/src/api/
protoc -I proto/ proto/*.proto --go_out=plugins=grpc:api


# protoc -I proto/ proto/*.proto  \
#     --go_out=plugins=grpc:api \
#     --plugin=protoc-gen-ts=../../siberia/app/node_modules/.bin/protoc-gen-ts \
#     --ts_out=service=true:../../siberia/app/src/api \
#     --js_out=import_style=commonjs,binary:../../siberia/app/src/api \

case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        ls ../../siberia/app/src/api/* | while read line;do
        sed -i "" '1s/^/\/* eslint-disable *\/ /' $line
        #sed -i "" 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        #sed -i "" 's/grpc-web-client/grpc-web/' $line
        done
     ;;
   Linux)
     echo 'Linux'
        ls ../../siberia/app/src/api/* | while read line;do
        sed -i '1s/^/\/* eslint-disable *\/ /' $line
        #sed -i 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        #sed -i 's/grpc-web-client/grpc-web/' $line
        done
     ;;
   *)
     echo 'other OS' 
     ;;
esac




