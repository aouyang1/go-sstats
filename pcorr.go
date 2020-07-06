package sstats

import (
	"math"
)

// PCorr implements the Statistic interface for streaming pearson correlation, (Σx*y - n*x̄*ȳ)/(sqrt(Σ(x^2-n*x̄))*sqrt(Σ(y^2-n*ȳ)))
type PCorr struct {
	xy     *SumProd
	xx, yy *SumSq
	xm, ym *Mean
}

// NewPCorr creates a new pearson correlation statistic with a given circular buffer size
func NewPCorr(size int) (*PCorr, error) {
	if size < 1 {
		return nil, errorInvalidSize
	}
	xy, err := NewSumProd(size)
	if err != nil {
		return nil, err
	}

	xx, err := NewSumSq(size)
	if err != nil {
		return nil, err
	}

	yy, err := NewSumSq(size)
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

	m := &PCorr{
		xy: xy,
		xx: xx, yy: yy,
		xm: xm, ym: ym,
	}
	return m, nil
}

// Update adds a new element to the pearson correlation circular buffer
func (p *PCorr) Update(x, y float64) {
	p.xy.Update(x, y)
	p.xx.Update(x)
	p.yy.Update(y)
	p.xm.Update(x)
	p.ym.Update(y)
}

// Reset clears out the values in the circular buffer and reset ptr and tail pointers
func (p *PCorr) Reset() {
	p.xy.Reset()
	p.xx.Reset()
	p.yy.Reset()
	p.xm.Reset()
	p.ym.Reset()
}

// Value computes the current pearson correlation value of the circular buffer
func (p *PCorr) Value() float64 {
	n := float64(p.Len())
	xm := p.xm.Value()
	ym := p.ym.Value()
	denom := (math.Sqrt(p.xx.Value()-n*xm*xm) * math.Sqrt(p.yy.Value()-n*ym*ym))
	if denom == 0 {
		return 0
	}
	return (p.xy.Value() - n*xm*ym) / denom
}

// Len returns the number of current elements stored in the circular buffer
func (p *PCorr) Len() int {
	return p.xy.Len()
}
