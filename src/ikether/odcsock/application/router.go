//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

/**
 * This a controller like MVC's Controller
 */
type Controllerable interface {
	/**
	 * When a command with the specific command-name incoming, this method Handle() will be invok
	 * You can handle get the command's information with @incomingMsg, and replay the remote device
	 * with @replay
	 *
	 * @incomingMsg The incoming message's all information
	 * @replay 		Doing replay with this
	 * @return error
	 */
	Handle(incomingMsg *IncomingMessage, replay *Replay) error

	/**
	 * Before the Controller is add to the route table, This the method Init() will be
	 * invok. And if error returned, The Controller will be not set
	 */
	Init() error
}

/**
 * A router table and manager
 */
type Router struct {
	routTable map[string]Controllerable
}

func NewRouter() *Router {
	return &Router{make(map[string]Controllerable)}
}

/**
 * Add a route path for the specific command
 * @warning this is not goroutine safe, just call this before server start
 *
 * @cmd 	The command's name
 * @ctrler 	A Controllerable interface, in which the command with command-name @cmd will be
 *			handled by @ctrler
 */
func (r *Router) Route(cmd string, ctrler Controllerable) {
	err := ctrler.Init()
	if err != nil {
		panic(err)
	}

	r.routTable[cmd] = ctrler
}

/**
 * Getting the controller instance with the command's name from route table
 *
 * @cmd The command's name
 * @return the Controller to deal with command with name @cmd
 */
func (r *Router) GetCtrler(cmd string) Controllerable {
	return r.routTable[cmd]
}
