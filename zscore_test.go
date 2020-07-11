package sstats

import (
	"math"
	"testing"
)

func TestZScoreUpdate(t *testing.T) {
	sp, err := NewZScore(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 2, 1, 2, 3, 2, 1}
	expected := []float64{0, 0.707, 1, 0, -0.956, 0, 0.956, 0, -0.956}
	for i, v := range valx {
		sp.Update(v)
		val := sp.Value()
		if math.Abs(val-expected[i]) > 1e-3 {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], val)
			continue
		}
	}
}

func BenchmarkZScoreUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewZScore(window)
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
