//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package config

import (
	"log"

	"github.com/zjwdmlmx/goini"
)

var Config *goini.INI

func init() {
	Config = goini.New()
	// read configure from .ini file
	err := Config.ParseFile("/etc/odcsock.cnf")

	if err != nil {
		log.Println("configure file parse failed!")
		log.Fatalln(err)
		return
	}
}
