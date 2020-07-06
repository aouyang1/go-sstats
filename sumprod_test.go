package sstats

import (
	"testing"
)

func TestSumProdUpdate(t *testing.T) {
	sp, err := NewSumProd(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valy := []float64{2, 3, 4, 5, 6, 7, 8, 9, 1}
	expected := []float64{2, 8, 20, 40, 70, 110, 160, 220, 209}
	for i, v := range valx {
		sp.Update(v, valy[i])
		mean := sp.Value()
		if mean != expected[i] {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], mean)
			continue
		}
	}
}

func BenchmarkSumProdUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewSumProd(window)
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
