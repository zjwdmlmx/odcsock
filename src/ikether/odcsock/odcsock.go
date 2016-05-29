//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package main

import (
	"fmt"
	"ikether/odcsock/application"
	"ikether/odcsock/config"
	"ikether/odcsock/controllers"
	"log"
	"net"
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
	if err != nil {
		log.Println(err)
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

		go handleConnection(conn, router)
	}
}
