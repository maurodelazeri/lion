your textn#!/usr/bin/env bash
protoc -I proto/ proto/enum.proto proto/api.proto --go_out=plugins=grpc:api
protoc -I proto/ proto/events.proto --go_out=events
