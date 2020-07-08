package sstats

// SumSq implements the Statistic interface for streaming sum squared Σx*x
type SumSq struct {
	buffer []float64
	ptr    int
	full   bool

	sum float64
}

// NewSumSq creates a new sum squared statistic with a given circular buffer size
func NewSumSq(size int) (*SumSq, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	s := &SumSq{buffer: make([]float64, size)}
	return s, nil
}

// Update adds a new element to the sum squared circular buffer
func (s *SumSq) Update(x float64) {
	s.sum += x*x - s.buffer[s.ptr]
	s.buffer[s.ptr] = x * x
	s.ptr = (s.ptr + 1) % len(s.buffer)
	if !s.full && s.ptr == 0 {
		s.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (s *SumSq) Reset() {
	for i := 0; i < len(s.buffer); i++ {
		s.buffer[i] = 0
	}
	s.ptr = 0
	s.full = false
	s.sum = 0
}

// Value computes the current sum squared value of the circular buffer
func (s *SumSq) Value() float64 {
	return s.sum
}

// Len returns the number of current elements stored in the circular buffer
func (s *SumSq) Len() int {
	if s.full {
		return len(s.buffer)
	}
	return s.ptr
}
