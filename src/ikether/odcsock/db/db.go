//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package db

import (
	"ikether/odcsock/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var (
		err              error
		dbType           string
		connectingString string
		ok               bool
	)

	dbType, ok = config.Config.Get("type")

	if !ok {
		log.Println("configure file get \"type\" failed! Using default value \"mysql\"")
		dbType = "mysql"
	}

	connectingString, ok = config.Config.Get("connectingString")

	if !ok {
		log.Println("configure file get \"connectingString\" failed! Using default value \"test:test@\\test?charset=utf8&parseTime=True\"")
		connectingString = "test:test@\\test?charset=utf8&parseTime=True"
	}

	DB, err = gorm.Open(dbType, connectingString)

	if err != nil {
		log.Println("database initial failed!")
		log.Fatalln(err)
		return
	}
}
