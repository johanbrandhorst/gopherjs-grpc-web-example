// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

//go:generate gopherjs build client.go -o html/index.js
//go:generate go-bindata -pkg compiled -nometadata -o compiled/client.go -prefix html ./html
//go:generate bash -c "rm html/*.js*"

import (
	"fmt"

	"github.com/johanbrandhorst/gopherjs-grpc-web-example/client/proto/library"
	"github.com/johanbrandhorst/gopherjs-grpc-web/logging"
)

func main() {
	rootLogger := logging.GetRootLogger()
	rootLogger.SetLevel(logging.ALL)
	rootLogger.AddHandler(logging.StandardHandler)

	client := library.NewBookServiceClient("http://localhost:8081")
	rootLogger.Info("Getting book")
	book, err := client.GetBook(library.NewGetBookRequest(140008381))
	if err != nil {
		fmt.Println("Got request error:", err)
		println("Got request error:", err)
		return
	}
	fmt.Println("Got book:", book)
	println("Got book:", book)

	return
}
