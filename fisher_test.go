package sstats

import (
	"math"
	"testing"
)

func TestFisherUpdate(t *testing.T) {
	sp, err := NewFisher(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 2, 1, 2, 3, 2, 1}
	expected := []float64{0, 0.881, math.Inf(1), 0, -1.899, 0, 1.899, 0, -1.899}
	for i, v := range valx {
		sp.Update(v)
		val := sp.Value()
		if math.Abs(val-expected[i]) > 1e-3 {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], val)
			continue
		}
	}
}

func BenchmarkFisherUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewFisher(window)
	if err != nil {
		b.Fatal(err)
	}

	for j := 0; j < b.N; j++ {
		for i := 0; i < numValues; i++ {
			sp.Update(float64(i))
		}
		sp.Reset()
	}
}
