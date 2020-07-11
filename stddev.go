package sstats

import "math"

// StdDev computes the streaming standard deviation, sqrt((Σx^2 + n*x̄^2 - 2*x̄*Σx)/n-1)
type StdDev struct {
	xx *SumSq
	xm *Mean
}

// NewStdDev creates a new standard deviation statistic with a given circular buffer size
func NewStdDev(size int) (*StdDev, error) {
	xx, err := NewSumSq(size)
	if err != nil {
		return nil, err
	}

	xm, err := NewMean(size)
	if err != nil {
		return nil, err
	}

	s := &StdDev{
		xx: xx,
		xm: xm,
	}
	return s, nil
}

// Update adds a new element to the standard deviation circular buffer
func (s *StdDev) Update(x float64) {
	s.xx.Update(x)
	s.xm.Update(x)
}

// UpdateBulk adds multiple elements to the standard deviation circular buffer
func (s *StdDev) UpdateBulk(xb []float64) error {
	for _, x := range xb {
		s.xx.Update(x)
		s.xm.Update(x)
	}
	return nil
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (s *StdDev) Reset() {
	s.xx.Reset()
	s.xm.Reset()
}

// Value computes the current standard deviation value of the circular buffer
func (s *StdDev) Value() float64 {
	n := float64(s.Len())
	if n <= 1 {
		return 0
	}
	xm := s.xm.Value()
	return math.Sqrt((s.xx.Value() + n*xm*xm - 2*xm*s.xm.Sum()) / (n - 1))
}

// Len returns the number of current elements stored in the circular buffer
func (s *StdDev) Len() int {
	return s.xm.Len()
}

func (s *StdDev) Mean() float64 {
	return s.xm.Value()
}
