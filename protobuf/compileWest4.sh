#!/usr/bin/env bash
rm  -rf ../../siberia/app/src/west4API/*
rm -rf west4API/*
 
#protoc -I west4API-proto/ west4API-proto/*.proto  --js_out=import_style=commonjs,import_style=typescript:../../siberia/app/src/west4API/ --grpc-web_out=import_style=commonjs,import_style=typescript,mode=grpcwebtext:../../siberia/app/src/west4API/
protoc -I west4API-proto/ west4API-proto/*.proto  --js_out=import_style=commonjs:../../siberia/app/src/west4API/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../../siberia/app/src/west4API/
protoc -I west4API-proto/ west4API-proto/*.proto --go_out=plugins=grpc:west4API

case "$(uname -s)" in

   Darwin)
     echo 'Mac OS X'
        ls ../../siberia/app/src/west4API/* | while read line;do
        sed -i "" '1s/^/\/* eslint-disable *\/ /' $line
        #sed -i "" 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        #sed -i "" 's/grpc-web-client/grpc-web/' $line
        done
     ;;
   Linux)
     echo 'Linux'
        ls ../../siberia/app/src/west4API/* | while read line;do
        sed -i '1s/^/\/* eslint-disable *\/ /' $line
        #sed -i 's/grpc-web-client/@improbable-eng\/grpc-web/' $line
        #sed -i 's/grpc-web-client/grpc-web/' $line
        done
     ;;
   *)
     echo 'other OS' 
     ;;
esac
