// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/johanbrandhorst/gopherjs-grpc-web-example/client/compiled"
	"github.com/johanbrandhorst/gopherjs-grpc-web-example/server"
	"github.com/johanbrandhorst/gopherjs-grpc-web-example/server/proto/library"
)

var logger *logrus.Logger

// If you change this, you'll need to change the cert as well
const gRPCAddr = "0.0.0.0:8081"
const staticAddr = "0.0.0.0:8080"

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.Kitchen,
		DisableSorting:  true,
	})
	grpclog.SetLogger(logger)
}

func main() {
	gs := grpc.NewServer()
	library.RegisterBookServiceServer(gs, &server.BookService{})

	gRPCServer := http.Server{
		Addr:    gRPCAddr,
		Handler: gs,
	}

	staticServer := http.Server{
		Addr: staticAddr,
		Handler: http.FileServer(&assetfs.AssetFS{
			Asset:     compiled.Asset,
			AssetDir:  compiled.AssetDir,
			AssetInfo: compiled.AssetInfo,
		}),
	}

	logger.Warn("Serving gRPC on http://", gRPCAddr)
	go gRPCServer.ListenAndServe()

	logger.Warn("Serving static on http://", staticAddr)
	logger.Fatal(staticServer.ListenAndServe())
}
