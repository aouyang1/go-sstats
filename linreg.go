package sstats

// LinReg implements the Statistic interface for streaming linear regression
type LinReg struct {
	xstd   *StdDev
	xm, ym *Mean
	c      *Cov
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

	xm, err := NewMean(size)
	if err != nil {
		return nil, err
	}

	ym, err := NewMean(size)
	if err != nil {
		return nil, err
	}

	c, err := NewCov(size)
	if err != nil {
		return nil, err
	}

	l := &LinReg{
		xstd: xstd,
		xm:   xm, ym: ym, c: c,
	}
	return l, nil
}

// Update adds a new element to the linear regression circular buffer
func (l *LinReg) Update(x, y float64) {
	l.xstd.Update(x)
	l.xm.Update(x)
	l.ym.Update(y)
	l.c.Update(x, y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (l *LinReg) Reset() {
	l.xstd.Reset()
	l.xm.Reset()
	l.ym.Reset()
	l.c.Reset()
}

// Value computes the current linear regression value of the circular buffer
func (l *LinReg) Value() (float64, float64) {
	xstd := l.xstd.Value()
	if xstd == 0 {
		return 0, 0
	}
	beta := l.c.Value() / (xstd * xstd)
	alpha := l.ym.Value() - beta*l.xm.Value()
	return alpha, beta
}

// Len returns the number of current elements stored in the circular buffer
func (l *LinReg) Len() int {
	return l.xstd.Len()
}
