#!/usr/bin/env bash
# Copyright 2017 Johan Brandhorst. All Rights Reserved.
# See LICENSE for licensing terms.

protoc proto/library/book_service.proto \
    --go_out=plugins=grpc:./server/ \
    --gopherjs_out=:./client/ \
    --grpc-web_out=out=./client/proto/library/book_service.pb.js,mode=base64:./ \
    --js_out=import_style=commonjs,binary:./client/ \
    --plugin=protoc-gen-js_service=./node_modules/.bin/protoc-gen-js_service \
    --js_service_out=:./client/
