// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

#!/usr/bin/env bash

protoc proto/library/book_service.proto --go_out=plugins=grpc:./server/ --gopherjs_out=:./client/
