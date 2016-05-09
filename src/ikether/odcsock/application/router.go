//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

type Controllerable interface {
	Handle(incomingMsg *IncomingMessage, replay *Replay) error
	Init() error
}

type Router struct {
	routTable map[string]Controllerable
}

func NewRouter() *Router {
	return &Router{make(map[string]Controllerable)}
}

func (r *Router) Route(cmd string, ctrler Controllerable) {
	err := ctrler.Init()
	if err != nil {
		panic(err)
	}

	r.routTable[cmd] = ctrler
}

func (r *Router) GetCtrler(cmd string) Controllerable {
	return r.routTable[cmd]
}
