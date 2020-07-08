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

	m := &LinReg{
		xstd: xstd,
		xm:   xm, ym: ym, c: c,
	}
	return m, nil
}

// Update adds a new element to the linear regression circular buffer
func (p *LinReg) Update(x, y float64) {
	p.xstd.Update(x)
	p.xm.Update(x)
	p.ym.Update(y)
	p.c.Update(x, y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (p *LinReg) Reset() {
	p.xstd.Reset()
	p.xm.Reset()
	p.ym.Reset()
	p.c.Reset()
}

// Value computes the current linear regression value of the circular buffer
func (p *LinReg) Value() (float64, float64) {
	xstd := p.xstd.Value()
	if xstd == 0 {
		return 0, 0
	}
	beta := p.c.Value() / (xstd * xstd)
	alpha := p.ym.Value() - beta*p.xm.Value()
	return alpha, beta
}

// Len returns the number of current elements stored in the circular buffer
func (p *LinReg) Len() int {
	return p.xstd.Len()
}
