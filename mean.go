package sstats

// Mean implements the Statistic interface for streaming mean, 1/n*Σx
type Mean struct {
	s *Sum
}

// NewMean creates a new mean statistic with a given circular buffer size
func NewMean(size int) (*Mean, error) {
	s, err := NewSum(size)
	if err != nil {
		return nil, err
	}
	m := &Mean{s: s}
	return m, nil
}

// Update adds a new element to the mean circular buffer
func (m *Mean) Update(x float64) {
	m.s.Update(x)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (m *Mean) Reset() {
	m.s.Reset()
}

// Value computes the current mean value of the circular buffer
func (m *Mean) Value() float64 {
	n := float64(m.Len())
	if n == 0 {
		return 0
	}
	return m.s.Value() / n
}

// Len returns the number of current elements stored in the circular buffer
func (m *Mean) Len() int {
	return m.s.Len()
}
