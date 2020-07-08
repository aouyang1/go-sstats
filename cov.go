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

	c := &Cov{
		xy: xy,
		x:  x, y: y,
		xm: xm, ym: ym,
	}
	return c, nil
}

// Update adds a new element to the covariance circular buffer
func (c *Cov) Update(x, y float64) {
	c.xy.Update(x, y)
	c.x.Update(x)
	c.y.Update(y)
	c.xm.Update(x)
	c.ym.Update(y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (c *Cov) Reset() {
	c.xy.Reset()
	c.x.Reset()
	c.y.Reset()
	c.xm.Reset()
	c.ym.Reset()
}

// Value computes the current covariance value of the circular buffer
func (c *Cov) Value() float64 {
	n := float64(c.Len())
	if n <= 1 {
		return 0
	}
	xm := c.xm.Value()
	ym := c.ym.Value()
	return (c.xy.Value() - xm*c.y.Value() - ym*c.x.Value() + n*xm*ym) / (n - 1)
}

// Len returns the number of current elements stored in the circular buffer
func (c *Cov) Len() int {
	return c.xy.Len()
}
