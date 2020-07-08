package sstats

import (
	"math"
	"testing"
)

func TestCovUpdate(t *testing.T) {
	sp, err := NewCov(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 2, 1, 2, 3, 2, 1}
	valy := []float64{2, 4, 6, 4, 2, 1, 0, 1, 2}
	expected := []float64{0, 1, 2, 1.333, 1.4, 1, 0.35, -0.5, -0.7}
	for i, v := range valx {
		sp.Update(v, valy[i])
		val := sp.Value()
		if math.Abs(val-expected[i]) > 1e-3 {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], val)
			continue
		}
	}
}

func BenchmarkCovUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewCov(window)
	if err != nil {
		b.Fatal(err)
	}

	for j := 0; j < b.N; j++ {
		for i := 0; i < numValues; i++ {
			sp.Update(float64(i), float64(i))
		}
		sp.Reset()
	}
}
