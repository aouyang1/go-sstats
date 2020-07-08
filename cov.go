package sstats

// Cov implements the Statistic interface for streaming covariance
type Cov struct {
	xy     *SumProd
	x, y   *Sum
	xm, ym *Mean
}

// NewCov creates a new covariance statistic with a given circular buffer size
func NewCov(size int) (*Cov, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	xy, err := NewSumProd(size)
	if err != nil {
		return nil, err
	}

	x, err := NewSum(size)
	if err != nil {
		return nil, err
	}

	y, err := NewSum(size)
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

	m := &Cov{
		xy: xy,
		x:  x, y: y,
		xm: xm, ym: ym,
	}
	return m, nil
}

// Update adds a new element to the covariance circular buffer
func (p *Cov) Update(x, y float64) {
	p.xy.Update(x, y)
	p.x.Update(x)
	p.y.Update(y)
	p.xm.Update(x)
	p.ym.Update(y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (p *Cov) Reset() {
	p.xy.Reset()
	p.x.Reset()
	p.y.Reset()
	p.xm.Reset()
	p.ym.Reset()
}

// Value computes the current covariance value of the circular buffer
func (p *Cov) Value() float64 {
	n := float64(p.Len())
	if n <= 1 {
		return 0
	}
	xm := p.xm.Value()
	ym := p.ym.Value()
	return (p.xy.Value() - xm*p.y.Value() - ym*p.x.Value() + n*xm*ym) / (n - 1)
}

// Len returns the number of current elements stored in the circular buffer
func (p *Cov) Len() int {
	return p.xy.Len()
}
