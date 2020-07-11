package sstats

// SumProd computes the streaming sum product, Σx*y
type SumProd struct {
	buffer []float64
	ptr    int
	full   bool

	sum float64
}

// NewSumProd creates a new sum product statistic with a given circular buffer size
func NewSumProd(size int) (*SumProd, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	s := &SumProd{buffer: make([]float64, size)}
	return s, nil
}

// Update adds a new element to the sum product circular buffer
func (s *SumProd) Update(x, y float64) {
	s.sum += x*y - s.buffer[s.ptr]
	s.buffer[s.ptr] = x * y
	s.ptr = (s.ptr + 1) % len(s.buffer)
	if !s.full && s.ptr == 0 {
		s.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (s *SumProd) Reset() {
	for i := 0; i < len(s.buffer); i++ {
		s.buffer[i] = 0
	}
	s.ptr = 0
	s.full = false
	s.sum = 0
}

// Value computes the current sum product value of the circular buffer
func (s *SumProd) Value() float64 {
	return s.sum
}

// Len returns the number of current elements stored in the circular buffer
func (s *SumProd) Len() int {
	if s.full {
		return len(s.buffer)
	}
	return s.ptr
}
