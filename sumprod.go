package sstats

// SumProd implements the Statistic interface for streaming sum product, Σx*y
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
	m := &SumProd{buffer: make([]float64, size)}
	return m, nil
}

// Update adds a new element to the sum product circular buffer
func (m *SumProd) Update(x, y float64) {
	m.sum += x*y - m.buffer[m.ptr]
	m.buffer[m.ptr] = x * y
	m.ptr = (m.ptr + 1) % len(m.buffer)
	if !m.full && m.ptr == 0 {
		m.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (m *SumProd) Reset() {
	for i := 0; i < len(m.buffer); i++ {
		m.buffer[i] = 0
	}
	m.ptr = 0
	m.full = false
	m.sum = 0
}

// Value computes the current sum product value of the circular buffer
func (m *SumProd) Value() float64 {
	return m.sum
}

// Len returns the number of current elements stored in the circular buffer
func (m *SumProd) Len() int {
	if m.full {
		return len(m.buffer)
	}
	return m.ptr
}
