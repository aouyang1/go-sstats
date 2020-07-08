package sstats

// ZScore implements the Statistic interface for streaming z-score
type ZScore struct {
	xm   *Mean
	xstd *StdDev

	val float64
}

// NewZScore creates a new z-score statistic with a given circular buffer size
func NewZScore(size int) (*ZScore, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	xm, err := NewMean(size)
	if err != nil {
		return nil, err
	}

	xstd, err := NewStdDev(size)
	if err != nil {
		return nil, err
	}

	z := &ZScore{xm: xm, xstd: xstd}
	return z, nil
}

// Update adds a new element to the z-score circular buffer
func (z *ZScore) Update(x float64) {
	z.xm.Update(x)
	z.xstd.Update(x)
	std := z.xstd.Value()
	if std == 0 {
		z.val = 0
		return
	}
	z.val = (x - z.xm.Value()) / std
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (z *ZScore) Reset() {
	z.xm.Reset()
	z.xstd.Reset()
}

// Value computes the current z-score value of the circular buffer
func (z *ZScore) Value() float64 {
	return z.val
}

// Len returns the number of current elements stored in the circular buffer
func (z *ZScore) Len() int {
	return z.xm.Len()
}
