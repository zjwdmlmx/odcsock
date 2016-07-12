//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package proto

import (
	"errors"
	"strconv"
	"time"
)

const (
	CMD_NAME_V1 = "V1"
	CMD_NAME_BS = "BS"
)

var (
	DEVICE_STATE_SOS uint32 = 0xefffffff
	/**
	 * TODO
	 * Device state of others
	 */
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
	/**
	 * the vendor of the remote device
	 */
	Vendor string

	/**
	 * the IMEI ID number
	 */
	Id string

	/**
	 * the command name of current command
	 */
	Cmd string
}

type V1Command struct {
	command_base

	/**
	 * is the command Valid
	 */
	Valid bool

	/**
	 * the time of this command created
	 * @warning cause of the remote device's time is not always accuracy
	 * the time is reset to the command Server recived
	 */
	Time time.Time

	/**
	 * the latitude of the device's Location
	 */
	Latitude float64

	/**
	 * the longitude of the device's location
	 */
	Longitude float64

	/**
	 * the status of the device
	 */
	State uint32

	/**
	 * now just do not need belows
	 */
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

	//cmd.Time, err = time.ParseInLocation("020106150405", fmt.Sprintf("%s%s", params[11], params[3]), time.UTC)

	//if err != nil {
	//	return
	//}

	cmd.Time = time.Now()

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
	err = errors.New("V1 command not support build")
	return
}

func (cmd *V1Command) BuildString() (cmdstr string, err error) {
	err = errors.New("V1 command not support build")
	return
}

/**
 * The heartbeat command with name XT
 * @warning unused now
 */
type XTCommand struct {
	command_base
	Time time.Time
}

type BSCommand struct {
	command_base
	/**
	 * There is many unused or unknowed
	 *
	 * unknow ......
	 * unused ......
	 */

	/**
	 * the status of the device
	 */
	State uint32
}

func (cmd *BSCommand) GetCmd() string {
	return cmd.Cmd
}

func (cmd *BSCommand) ParseParams(params []string) (err error) {
	if 18 > len(params) {
		err = errors.New("params' size is not enough! when BSCommand")
		return
	}

	var state uint64
	state, err = strconv.ParseUint(params[16], 16, 32)

	if err != nil {
		return
	}

	cmd.State = uint32(state)
	cmd.Cmd = params[2]
	cmd.Id = params[1]
	cmd.Vendor = params[0]

	return
}

func (cmd *BSCommand) ParseString(cmdstr string) (err error) {
	var params []string
	params, err = ParseStringToParams(cmdstr)

	if err != nil {
		return
	}

	err = cmd.ParseParams(params)
	return
}

func (cmd *BSCommand) BuildParams() (params []string, err error) {
	err = errors.New("BS command not support build")
	return
}

func (cmd *BSCommand) BuildString() (cmdstr string, err error) {
	err = errors.New("BS command not support build")
	return
}
