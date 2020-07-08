package sstats

// Sum implements the Statistic interface for streaming sum, Σx
type Sum struct {
	buffer []float64
	ptr    int
	full   bool

	sum float64
}

// NewSum creates a new sum statistic with a given circular buffer size
func NewSum(size int) (*Sum, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	s := &Sum{buffer: make([]float64, size)}
	return s, nil
}

// Update adds a new element to the sum circular buffer
func (s *Sum) Update(x float64) {
	s.sum += x - s.buffer[s.ptr]
	s.buffer[s.ptr] = x
	s.ptr = (s.ptr + 1) % len(s.buffer)
	if !s.full && s.ptr == 0 {
		s.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (s *Sum) Reset() {
	for i := 0; i < len(s.buffer); i++ {
		s.buffer[i] = 0
	}
	s.ptr = 0
	s.full = false
	s.sum = 0
}

// Value computes the current sum value of the circular buffer
func (s *Sum) Value() float64 {
	return s.sum
}

// Len returns the number of current elements stored in the circular buffer
func (s *Sum) Len() int {
	if s.full {
		return len(s.buffer)
	}
	return s.ptr
}
