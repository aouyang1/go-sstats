package sstats

// ZScore computes the streaming z-score (x-ux)/stdx
type ZScore struct {
	xstd *StdDev

	val float64
}

// NewZScore creates a new z-score statistic with a given circular buffer size
func NewZScore(size int) (*ZScore, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	xstd, err := NewStdDev(size)
	if err != nil {
		return nil, err
	}

	z := &ZScore{xstd: xstd}
	return z, nil
}

// Update adds a new element to the z-score circular buffer
func (z *ZScore) Update(x float64) {
	z.xstd.Update(x)
	z.val = x
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (z *ZScore) Reset() {
	z.xstd.Reset()
}

// Value computes the current z-score value of the circular buffer
func (z *ZScore) Value() float64 {
	std := z.xstd.Value()
	if std == 0 {
		return 0
	}
	return (z.val - z.xstd.Mean()) / std
}

// Len returns the number of current elements stored in the circular buffer
func (z *ZScore) Len() int {
	return z.xstd.Len()
}
