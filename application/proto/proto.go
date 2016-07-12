//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package proto

import (
	"errors"
	"fmt"
	"strings"
)

func StringToCommand(cmdstr string) (cmd Command, err error) {
	var (
		params []string
	)

	params, err = ParseStringToParams(cmdstr)

	if err != nil {
		return
	}

	return ParamsToCommand(params)
}

func ParamsToCommand(params []string) (cmd Command, err error) {
	if 4 > len(params) {
		err = errors.New("params' size is not enough!")
		return
	}

	cmdName := params[2]

	if cmdName == CMD_NAME_V1 {
		cmd = &V1Command{}
	} else if cmdName == CMD_NAME_BS {
		cmd = &BSCommand{}
	} else {
		err = errors.New(fmt.Sprintf("unregist command name: %s", cmdName))
		return
	}

	err = cmd.ParseParams(params)

	if err != nil {
		return
	}

	return
}

func ParseStringToParams(cmdstr string) (params []string, err error) {
	cmdLen := len(cmdstr)
	if cmdLen < 3 {
		err = errors.New("unacceptable command string")
		return
	}

	if cmdstr[0] != '*' || cmdstr[cmdLen-1] != '#' {
		err = errors.New("Invalid command beginer or ender")
		return
	}

	params = strings.Split(cmdstr[1:cmdLen-1], ",")

	/**
	 * as the protocol say the params must last 4
	 */
	if len(params) < 4 {
		params = nil
		err = errors.New("bad command string!")
		return
	}
	return
}

func buildParams(params []string) (cmdstr string) {
	cmdstr = "*" + strings.Join(params, ",") + "#"
	return
}
