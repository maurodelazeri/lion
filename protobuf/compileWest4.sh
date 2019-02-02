#!/usr/bin/env bash

rm west4API/*

protoc -I west4API-proto/ west4API-proto/*.proto  \
    --go_out=plugins=grpc:west4API \
