package sstats

import (
	"testing"
)

func TestSumSqUpdate(t *testing.T) {
	ss, err := NewSumSq(5)
	if err != nil {
		t.Fatal(err)
	}
	vals := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []float64{1, 5, 14, 30, 55, 90, 135, 190, 255}
	for i, v := range vals {
		ss.Update(v)
		mean := ss.Value()
		if mean != expected[i] {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], mean)
			continue
		}
	}
}

func BenchmarkSumSqUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	ss, err := NewSumSq(window)
	if err != nil {
		b.Fatal(err)
	}

	for j := 0; j < b.N; j++ {
		for i := 0; i < numValues; i++ {
			ss.Update(float64(i))
		}
		ss.Reset()
	}
}
