//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package proto

import (
    "strings"
    "errors"
    "time"
    "fmt"
    "regexp"
    "strconv"
)

var (
    latitude_regex = regexp.MustCompile("^([0-9]{2})([0-9]{2}\\.[0-9]{4})$")
    longitude_regex = regexp.MustCompile("^([0-9]{3})([0-9]{2}\\.[0-9]{4})$")
)

func parseLatitude(s string) (latitude float64, err error) {
    var (
        latitude_D      float64
        latitude_M      float64
    )

    matchs := latitude_regex.FindAllStringSubmatch(s, -1)

    latitude_D, err = strconv.ParseFloat(matchs[0][1], 64)

    if err != nil {
        return
    }

    latitude_M, err = strconv.ParseFloat(matchs[0][2], 64)

    if err != nil {
        return
    }

    latitude = latitude_D + latitude_M/60
    return
}

func parseLongitude(s string) (longitude float64, err error) {
    var (
        longitude_D     float64
        longitude_M     float64
    )

    matchs := longitude_regex.FindAllStringSubmatch(s, -1)

    longitude_D, err = strconv.ParseFloat(matchs[0][1], 64)

    if err != nil {
        return
    }

    longitude_M, err = strconv.ParseFloat(matchs[0][2], 64)

    if err != nil {
        return
    }

    longitude = longitude_D + longitude_M/60
    return
}


type ClientCommand struct {
    Valid       bool
    Cmd         string
    Id          string
    Vendor      string
    Time        time.Time
    Latitude    float64
    Longitude   float64
    // now just do not need belows
    //Speed       float32
    //Direction   uint16
    //Battery     uint8
    State       uint32
}

func StringToClientCommand(cmdstr string) (cmd *ClientCommand, err error) {
    var (
        params      []string
    )

    params, err = ParseParams(cmdstr)

    if err != nil {
        return
    }

    return ParamsToClientCommand(params)
}

func ParamsToClientCommand(params []string) (cmd *ClientCommand, err error) {
    if 13 > len(params) {
        err = errors.New("params' size is not enough!")
        return
    }

    res := new(ClientCommand)

    /**
     * TODO: parse the south/north latitude from params
     */
    res.Latitude, err = parseLatitude(params[5])

    if err != nil {
        return
    }

    /**
     * TODO: parse the east/west longitude from params
     */
    res.Longitude, err = parseLongitude(params[7])

    if err != nil {
        return
    }

    res.Time, err = time.ParseInLocation("020106150405", fmt.Sprintf("%s%s", params[11], params[3]), time.UTC)

    if err != nil {
        return
    }

    if params[4] == "A" {
        res.Valid = true
    } else if params[4] == "V" {
        res.Valid = false
    } else {
        err = errors.New("Invalid data check bit")
        return
    }

    var state uint64
    state, err = strconv.ParseUint(params[12], 16, 32)

    if err != nil {
        return
    }

    res.State       = uint32(state)
    res.Cmd         = params[2]
    res.Id          = params[1]
    res.Vendor      = params[0]

    cmd = res
    return
}

func ParseParams(cmdstr string) (params []string, err error) {
    cmdLen := len(cmdstr)
    if cmdLen < 3 {
        err = errors.New("unacceptable command string")
        return
    }

    if cmdstr[0] != '*' || cmdstr[cmdLen - 1] != '#' {
        err = errors.New("Invalid command beginer or ender")
        return
    }

    params = strings.Split(cmdstr[1:cmdLen - 1], ",")

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

func BuildParams(params []string) (cmdstr string) {
    cmdstr = "*" + strings.Join(params, ",") + "#"
    return
}
