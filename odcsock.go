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
	"os"

	"github.com/zjwdmlmx/odcsock/application"
	"github.com/zjwdmlmx/odcsock/controllers"
	"github.com/zjwdmlmx/odcsock/global"
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

func initLog() {
	var (
		model   string
		ok      bool
		logPath string
		err     error
	)
	if model, ok = global.Config.Get("model"); !ok {
		log.Println("configure file's model not set. Using default debug model")
		model = "debug"
	}

	if logPath, ok = global.Config.Get("log"); !ok {
		log.Println("Configure file's log not set. Using default /tmp/odcsock.log")
		logPath = "/tmp/odcsock.log"
	}

	if model == "debug" {
		var f *os.File
		if f, err = os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
			panic("open log file error")
		}
		log.SetOutput(f)
	}
}

func main() {
	var (
		port    string
		address string
		ok      bool
	)
	initLog()

	port, ok = global.Config.Get("port")

	if !ok {
		log.Println("configure file's port not set. Using default 8898")
		port = "8898"
	}

	address, ok = global.Config.Get("address")

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
