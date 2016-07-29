//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package global

import (
	"log"

	"github.com/zjwdmlmx/goini"
)

// Config the golbal configure object
var Config *goini.INI

func initConfigure() {
	Config = goini.New()
	// read configure from .ini file
	err := Config.ParseFile("/etc/odcsock.cnf")

	if err != nil {
		log.Println("configure file parse failed!")
		log.Fatalln(err)
		return
	}
}
