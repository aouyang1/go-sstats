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

	m := &ZScore{xm: xm, xstd: xstd}
	return m, nil
}

// Update adds a new element to the z-score circular buffer
func (p *ZScore) Update(x float64) {
	p.xm.Update(x)
	p.xstd.Update(x)
	std := p.xstd.Value()
	if std == 0 {
		p.val = 0
		return
	}
	p.val = (x - p.xm.Value()) / std
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (p *ZScore) Reset() {
	p.xm.Reset()
	p.xstd.Reset()
}

// Value computes the current z-score value of the circular buffer
func (p *ZScore) Value() float64 {
	return p.val
}

// Len returns the number of current elements stored in the circular buffer
func (p *ZScore) Len() int {
	return p.xm.Len()
}
