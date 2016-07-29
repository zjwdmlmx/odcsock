//
// Author: ikether
// Email: zjwdmlmx@126.com
//
// Copyright (c) 2016 by ikether. All Rights Reserved.
//

package global

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisPool redis connection pool
var RedisPool *redis.Pool

func newPool(server, net string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(net, server)
		},
		Wait: true,
	}
}

func initRedis() {
	var (
		net    string
		server string
		ok     bool
	)
	if net, ok = Config.Get("redis_server_net"); !ok {
		log.Println("Configure not set redis_server_net using default tcp")
		net = "tcp"
	}

	if server, ok = Config.Get("redis_server"); !ok {
		log.Println("Configure not set redis_server using default 127.0.0.1:6379")
		server = "127.0.0.1:6379"
	}

	RedisPool = newPool(server, net)
}
