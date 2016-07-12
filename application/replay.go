//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"bufio"
	"net"
)

type Replay struct {
	conn   net.Conn
	buffer *bufio.Writer
}

func NewReplay(conn net.Conn) *Replay {
	res := new(Replay)
	res.conn = conn
	res.buffer = bufio.NewWriter(conn)
	return res
}

func (rep *Replay) Send(str string) (err error) {
	var ct int

	ct, err = rep.buffer.WriteString(str)

	if err != nil || len(str) != ct {
		return
	}

	err = rep.buffer.Flush()
	return
}

func (rep *Replay) SendBytes(bs []byte) (err error) {
	var ct int

	ct, err = rep.buffer.Write(bs)

	if err != nil || len(bs) != ct {
		return
	}

	err = rep.buffer.Flush()
	return
}
