//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"bufio"
	"ikether/odcsock/application/proto"
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
	Command *proto.ClientCommand
}

func NewIncomingMessage(conn net.Conn) *IncomingMessage {
	res := new(IncomingMessage)
	res.conn = conn
	res.buffer = bufio.NewReaderSize(res.conn, BufferSize)
	return res
}

func (v *IncomingMessage) ReadMessage() (err error) {
	var cmd string

	cmd, err = v.buffer.ReadString('#')

	if err != nil {
		return
	}

	v.Params, err = proto.ParseParams(cmd)

	if err != nil {
		return
	}

	v.Command, err = proto.ParamsToClientCommand(v.Params)
	return
}
