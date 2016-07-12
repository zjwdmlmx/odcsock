//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package model

import (
	"time"
)

type User struct {
	Id          uint64 `gorm:"primark_key"`
	Phone       string `gorm:"type:char(16)"`
	Password    string `gorm:"type:char(64)"`
	Code        string `gorm:"type:char(16)"`
	Imei        string `gorm:"type:char(16)"`
	Type        uint8  `gorm:"not null"`
	Email       string `gorm:"type:varchar(128)"`
	EmailValid  bool
	Name        string `gorm:"type:char(24)"`
	Cardid      string `gorm:"type:char(20)"`
	CardidValid bool
	Address     string `gorm:"type:varchar(512)"`
	Sex         bool
	Nickname    string `grom:"type:char(20)"`
	Date        time.Time
	Msglevel    uint
	Qq          string `grom:"type:char(24)"`
	Weibo       string `gorm:"type:char(24)"`
	LastLoginTm time.Time
	LastLoginIp string `gorm:"type:char(48)"`
}
