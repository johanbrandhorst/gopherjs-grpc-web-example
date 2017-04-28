#!/usr/bin/env bash

protoc proto/library/book_service.proto --go_out=plugins=grpc:./server/
