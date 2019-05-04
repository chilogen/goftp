// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This is a very simple ftpd server using this library as an example
// and as something to run tests against.
package main

import (
	"flag"
	"github.com/goftp/server/auth"
	"log"

	filedriver "github.com/goftp/file-driver"
	"github.com/goftp/server"
)

var (
	RootPath *string
)

func main() {
	RootPath = flag.String("root", "~/testftp", "Root directory for server")
	var (
		port = flag.Int("port", 2122, "Port")
		host = flag.String("host", "", "Port")
	)
	flag.Parse()
	if *RootPath == "" {
		log.Fatalf("Please set a root to serve with -root")
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: *RootPath,
		Perm:     server.NewSimplePerm("user", "group"),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     *port,
		Hostname: *host,
		//Auth:     &server.SimpleAuth{Name: *user, Password: *pass},
		Auth: auth.GetAuth(),
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
