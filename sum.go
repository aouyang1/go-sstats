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
	m := &Sum{buffer: make([]float64, size)}
	return m, nil
}

// Update adds a new element to the sum circular buffer
func (m *Sum) Update(x float64) {
	m.sum += x - m.buffer[m.ptr]
	m.buffer[m.ptr] = x
	m.ptr = (m.ptr + 1) % len(m.buffer)
	if !m.full && m.ptr == 0 {
		m.full = true
	}
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (m *Sum) Reset() {
	for i := 0; i < len(m.buffer); i++ {
		m.buffer[i] = 0
	}
	m.ptr = 0
	m.full = false
	m.sum = 0
}

// Value computes the current sum value of the circular buffer
func (m *Sum) Value() float64 {
	return m.sum
}

// Len returns the number of current elements stored in the circular buffer
func (m *Sum) Len() int {
	if m.full {
		return len(m.buffer)
	}
	return m.ptr
}
