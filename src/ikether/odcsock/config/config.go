package config

import (
	"log"

	"github.com/zjwdmlmx/goini"
)

var Config *goini.INI

func init() {
	Config = goini.New()
	err := Config.ParseFile("/etc/odcsock.cnf")

	if err != nil {
		log.Println("configure file parse failed!")
		log.Fatalln(err)
		return
	}
}
