package sstats

import (
	"math"
	"testing"
)

func TestStdDevUpdate(t *testing.T) {
	sp, err := NewStdDev(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 2, 1, 2, 3, 2, 1}
	expected := []float64{0, 0.707, 1, 0.816, 0.837, 0.707, 0.837, 0.707, 0.837}
	for i, v := range valx {
		sp.Update(v)
		val := sp.Value()
		t.Log(val)
		if math.Abs(val-expected[i]) > 1e-3 {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], val)
			continue
		}
	}
}

func BenchmarkStdDevUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewStdDev(window)
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
