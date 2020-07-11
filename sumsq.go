package sstats

// SumSq computes the streaming sum squared Σx*x
type SumSq struct {
	xx *SumProd
}

// NewSumSq creates a new sum squared statistic with a given circular buffer size
func NewSumSq(size int) (*SumSq, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}

	xx, err := NewSumProd(size)
	if err != nil {
		return nil, err
	}

	s := &SumSq{xx: xx}
	return s, nil
}

// Update adds a new element to the sum squared circular buffer
func (s *SumSq) Update(x float64) {
	s.xx.Update(x, x)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (s *SumSq) Reset() {
	s.xx.Reset()
}

// Value computes the current sum squared value of the circular buffer
func (s *SumSq) Value() float64 {
	return s.xx.Value()
}

// Len returns the number of current elements stored in the circular buffer
func (s *SumSq) Len() int {
	return s.xx.Len()
}
