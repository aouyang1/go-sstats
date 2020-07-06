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
	m := &SumSq{buffer: make([]float64, size)}
	return m, nil
}

// Update adds a new element to the sum squared circular buffer
func (m *SumSq) Update(x float64) {
	m.sum += x*x - m.buffer[m.ptr]
	m.buffer[m.ptr] = x * x
	m.ptr = (m.ptr + 1) % len(m.buffer)
	if !m.full && m.ptr == 0 {
		m.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (m *SumSq) Reset() {
	for i := 0; i < len(m.buffer); i++ {
		m.buffer[i] = 0
	}
	m.ptr = 0
	m.full = false
	m.sum = 0
}

// Value computes the current sum squared value of the circular buffer
func (m *SumSq) Value() float64 {
	return m.sum
}

// Len returns the number of current elements stored in the circular buffer
func (m *SumSq) Len() int {
	if m.full {
		return len(m.buffer)
	}
	return m.ptr
}
