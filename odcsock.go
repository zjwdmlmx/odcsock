//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package main

import (
	"fmt"
	"log"
	"net"

	"code.csdn.net/zjwdmlmx/odcsock/application"
	"code.csdn.net/zjwdmlmx/odcsock/config"
	"code.csdn.net/zjwdmlmx/odcsock/controllers"
)

const (
	echo = `
           _      _____             _   _
  ____    | |____|  ___\ ____  ____| | / /
 / __ \ __| |  __| |___ / __ \|  __| |/ /
| |__| | _  | |__|____ | |__| | |__| |\ \
 \____/\____|____\_____|\____/|____|_| \_\
Copyright 2016 ikether. All Right reserved.

`
)

func handleConnection(conn net.Conn, router *application.Router) {
	log.Println("connection setup from", conn.RemoteAddr())

	app := application.NewApplication(conn, router)
	err := app.Handle()

	/**
	 * If bad error occured, just finished the goroutine
	 */
	if err != nil {
		log.Println("Connection will be closed!")
		conn.Close()
	}
}

func initServer(router *application.Router) {
	router.Route("V1", &controllers.LocationController{})
}

func main() {
	var (
		port    string
		address string
		ok      bool
	)
	port, ok = config.Config.Get("port")

	if !ok {
		log.Println("configure file's port not set. Using default 8898")
		port = "8898"
	}

	address, ok = config.Config.Get("address")

	if !ok {
		log.Println("configure file's address not set. Using default 0.0.0.0")
		address = "0.0.0.0"
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))

	if err != nil {
		log.Fatalln("listen failed! with error:", err)
	}

	router := application.NewRouter()
	initServer(router)

	log.Print(echo)

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Println("accept failed! with error:", err)
		}

		// foreach connection run a goroutine to handle the commands from remote
		go handleConnection(conn, router)
	}
}