//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"bufio"
	"ikether/odcsock/application/proto"
	"log"
	"net"
)

const (
	// the max message size of the device client
	// TODO: if the size is bigger than that .....
	BufferSize = 2048
)

type IncomingMessage struct {
	conn   net.Conn
	buffer *bufio.Reader

	Params  []string
	Command proto.Command
}

func NewIncomingMessage(conn net.Conn) *IncomingMessage {
	res := new(IncomingMessage)
	res.conn = conn
	res.buffer = bufio.NewReaderSize(res.conn, BufferSize)
	return res
}

/**
 * Reading message from NET, And initial the IncomingMessage object
 */
func (v *IncomingMessage) ReadMessage() (err error) {
	var cmd string

	cmd, err = v.buffer.ReadString('#')

	log.Printf("recived command string: %s\n", cmd)

	if err != nil {
		return
	}

	v.Params, err = proto.ParseStringToParams(cmd)

	if err != nil {
		return
	}

	v.Command, err = proto.ParamsToCommand(v.Params)

	return
}
