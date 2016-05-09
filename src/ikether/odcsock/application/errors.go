//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package application

import (
	"fmt"
	"time"
)

type CommonError struct {
	When time.Time
	What string
}

func (e *CommonError) Error() string {
	return fmt.Sprintf("[%v] %s", e.When, e.What)
}
