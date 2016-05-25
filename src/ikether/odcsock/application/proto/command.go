//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package proto

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Command interface {
	/**
	 * get the command's name
	 */
	GetCmd() string

	/**
	 * parse command from string like a csv line
	 *
	 * @param cmdstr: the command in string format
	 * @return error
	 */
	ParseString(cmdstr string) error

	/**
	 * parse command from a param list
	 *
	 * @param param: the command in a param list
	 * @return error
	 */
	ParseParams(param []string) error

	/**
	 * build command to string format
	 *
	 * @return the command string and error
	 */
	BuildString() (string, error)

	/**
	 * build command to param format
	 *
	 * @return the param in []string and error
	 */
	BuildParams() ([]string, error)
}

/**
 * All command's base field
 */
type command_base struct {
	Vendor string
	Id     string
	Cmd    string
}

type V1Command struct {
	command_base
	Valid     bool
	Time      time.Time
	Latitude  float64
	Longitude float64
	State     uint32
	// now just do not need belows
	//Speed       float32
	//Direction   uint16
	//Battery     uint8
}

func (cmd *V1Command) GetCmd() string {
	return cmd.Cmd
}

func (cmd *V1Command) ParseParams(params []string) (err error) {
	if 13 > len(params) {
		err = errors.New("params' size is not enough!")
		return
	}

	/**
	 * TODO: parse the south/north latitude from params
	 */
	cmd.Latitude, err = parseLatitude(params[5])

	if err != nil {
		return
	}

	/**
	 * TODO: parse the east/west longitude from params
	 */
	cmd.Longitude, err = parseLongitude(params[7])

	if err != nil {
		return
	}

	cmd.Time, err = time.ParseInLocation("020106150405", fmt.Sprintf("%s%s", params[11], params[3]), time.UTC)

	if err != nil {
		return
	}

	if params[4] == "A" {
		cmd.Valid = true
	} else if params[4] == "V" {
		cmd.Valid = false
	} else {
		err = errors.New("Invalid data check bit")
		return
	}

	var state uint64
	state, err = strconv.ParseUint(params[12], 16, 32)

	if err != nil {
		return
	}

	cmd.State = uint32(state)
	cmd.Cmd = params[2]
	cmd.Id = params[1]
	cmd.Vendor = params[0]

	return
}

func (cmd *V1Command) ParseString(cmdstr string) (err error) {
	var params []string
	params, err = ParseStringToParams(cmdstr)

	if err != nil {
		return
	}

	err = cmd.ParseParams(params)
	return
}

func (cmd *V1Command) BuildParams() (params []string, err error) {
	return
}

func (cmd *V1Command) BuildString() (cmdstr string, err error) {
	return
}

type XTCommand struct {
	command_base
	Time time.Time
}
