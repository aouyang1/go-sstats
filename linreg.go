package sstats

// LinReg computes the streaming linear regression
type LinReg struct {
	xstd *StdDev
	c    *Cov
}

// NewLinReg creates a new linear regression statistic with a given circular buffer size
func NewLinReg(size int) (*LinReg, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	xstd, err := NewStdDev(size)
	if err != nil {
		return nil, err
	}

	c, err := NewCov(size)
	if err != nil {
		return nil, err
	}

	l := &LinReg{
		xstd: xstd,
		c:    c,
	}
	return l, nil
}

// Update adds a new element to the linear regression circular buffer
func (l *LinReg) Update(x, y float64) {
	l.xstd.Update(x)
	l.c.Update(x, y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (l *LinReg) Reset() {
	l.xstd.Reset()
	l.c.Reset()
}

// Value computes the current linear regression value of the circular buffer
func (l *LinReg) Value() (float64, float64) {
	xstd := l.xstd.Value()
	if xstd == 0 {
		return 0, 0
	}
	beta := l.c.Value() / (xstd * xstd)
	alpha := l.c.YMean() - beta*l.xstd.Mean()
	return alpha, beta
}

// Len returns the number of current elements stored in the circular buffer
func (l *LinReg) Len() int {
	return l.xstd.Len()
}
