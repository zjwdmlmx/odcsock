//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type Application struct {
	router *Router
	conn   net.Conn
}

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
			// unacceptable errors
			if err == io.EOF || err == io.ErrUnexpectedEOF || err == io.ErrClosedPipe {
				break
			}
			continue
		}

		ctrler := app.router.GetCtrler(inComingMsg.Command.GetCmd())

		if ctrler == nil {
			err = errors.New(fmt.Sprintf("unregist Controller for the given command which name is %s", inComingMsg.Command.GetCmd()))
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
