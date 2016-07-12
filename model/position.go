//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package model

import (
	"time"
)

type Position struct {
	Id        uint64    `gorm:"primark_key"`
	Uid       uint64    `gorm:"not null"`
	Time      time.Time `gorm:"not null"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	State     int8
}
