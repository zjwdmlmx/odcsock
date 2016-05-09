//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"errors"
	"log"
	"net"
)

type Application struct {
	router *Router
	conn   net.Conn
}

var TheRouter *Router = NewRouter()

func NewApplication(conn net.Conn, router *Router) *Application {
	return &Application{router, conn}
}

func (app *Application) Handle() (err error) {
	inComingMsg := NewIncomingMessage(app.conn)
	replay := NewReplay(app.conn)

	for {
		err = inComingMsg.ReadMessage()

		if err != nil {
			log.Println(err)
			if err.Error() == "EOF" {
				break
			}
			continue
		}

		ctrler := app.router.GetCtrler(inComingMsg.Params[2])

		if ctrler == nil {
			err = errors.New("unkonw command")
			log.Println(err)
			continue
		} else {
			err = ctrler.Handle(inComingMsg, replay)
		}

		if err != nil {
			log.Println(err)
		}
	}
	return
}
