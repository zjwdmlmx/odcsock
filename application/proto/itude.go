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
	latitudeRegex  = regexp.MustCompile("^([0-9]{2})([0-9]{2}\\.[0-9]{4})$")
	longitudeRegex = regexp.MustCompile("^([0-9]{3})([0-9]{2}\\.[0-9]{4})$")
)

const (
	// for function parseItude
	longitudeType uint8 = 1
	latitudeType  uint8 = 2
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
		itudeD float64
		itudeM float64
	)

	var matchs [][]string

	if t == longitudeType {
		matchs = longitudeRegex.FindAllStringSubmatch(s, -1)
	} else if latitudeType == t {
		matchs = latitudeRegex.FindAllStringSubmatch(s, -1)
	}

	if !checkRegexSubmatch(matchs) {
		err = errors.New("parseLatitude failed with valid parse string")
		return
	}

	itudeD, err = strconv.ParseFloat(matchs[0][1], 64)

	if err != nil {
		return
	}

	itudeM, err = strconv.ParseFloat(matchs[0][2], 64)

	if err != nil {
		return
	}

	itude = itudeD + itudeM/60
	return
}

func parseLatitude(s string) (latitude float64, err error) {
	latitude, err = parseItude(latitudeType, s)
	return
}

func parseLongitude(s string) (longitude float64, err error) {
	longitude, err = parseItude(longitudeType, s)
	return
}
