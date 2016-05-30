//
// Author: ikether
// Email: zjwdmlmx@126.com
//
// Copyright (c) 2016 by ikether. All Rights Reserved.

package proto

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	latitude_regex  = regexp.MustCompile("^([0-9]{2})([0-9]{2}\\.[0-9]{4})$")
	longitude_regex = regexp.MustCompile("^([0-9]{3})([0-9]{2}\\.[0-9]{4})$")
)

const (
	// for function parseItude
	longitude_type uint8 = 1
	latitude_type  uint8 = 2
)

func checkRegexSubmatch(matchs [][]string) (valid bool) {
	if len(matchs) != 1 {
		valid = false
		return
	}

	if len(matchs[0]) != 3 {
		valid = false
		return
	}

	valid = true
	return
}

func parseItude(t uint8, s string) (itude float64, err error) {
	var (
		itude_D float64
		itude_M float64
	)

	var matchs [][]string

	if t == longitude_type {
		matchs = longitude_regex.FindAllStringSubmatch(s, -1)
	} else if latitude_type == t {
		matchs = latitude_regex.FindAllStringSubmatch(s, -1)
	}

	if !checkRegexSubmatch(matchs) {
		err = errors.New("parseLatitude failed with valid parse string!")
		return
	}

	itude_D, err = strconv.ParseFloat(matchs[0][1], 64)

	if err != nil {
		return
	}

	itude_M, err = strconv.ParseFloat(matchs[0][2], 64)

	if err != nil {
		return
	}

	itude = itude_D + itude_M/60
	return
}

func parseLatitude(s string) (latitude float64, err error) {
	latitude, err = parseItude(latitude_type, s)
	return
}

func parseLongitude(s string) (longitude float64, err error) {
	longitude, err = parseItude(longitude_type, s)
	return
}
