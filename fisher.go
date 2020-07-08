package sstats

import (
	"math"
)

// Fisher implements the Statistic interface for streaming fisher transform
type Fisher struct {
	zs *ZScore
}

// NewFisher creates a new fisher transform statistic with a given circular buffer size
func NewFisher(size int) (*Fisher, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	zs, err := NewZScore(size)
	if err != nil {
		return nil, err
	}

	f := &Fisher{zs: zs}
	return f, nil
}

// Update adds a new element to the fisher transform circular buffer
func (f *Fisher) Update(x float64) {
	f.zs.Update(x)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (f *Fisher) Reset() {
	f.zs.Reset()
}

// Value computes the current fisher transform value of the circular buffer
func (f *Fisher) Value() float64 {
	zs := f.zs.Value()
	return 0.5 * math.Log((1+zs)/(1-zs))
}

// Len returns the number of current elements stored in the circular buffer
func (f *Fisher) Len() int {
	return f.zs.Len()
}
