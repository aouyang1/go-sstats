package sstats

import (
	"errors"
)

var (
	errorInvalidSize  = errors.New("Invalid error size for statistic, must be greater than 0")
	errorDifferentLen = errors.New("Input x and y do not have the same length")
)

// Statistic is the interface for any streaming statistic
type Statistic interface {
	Reset()
	Value() float64
	Len() int
}
