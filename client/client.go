// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

//go:generate gopherjs build -m client.go -o html/index.js
//go:generate go-bindata -pkg compiled -nometadata -o compiled/client.go -prefix html ./html
//go:generate bash -c "rm html/*.js*"

import (
	"github.com/johanbrandhorst/gopherjs-grpc-web-example/client/proto/library"
)

var _ = &library.Book{}

func main() {}
