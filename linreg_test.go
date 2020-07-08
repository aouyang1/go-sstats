package sstats

import (
	"math"
	"testing"
)

func TestLinRegUpdate(t *testing.T) {
	sp, err := NewLinReg(5)
	if err != nil {
		t.Fatal(err)
	}
	valx := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valy := []float64{2, 4, 6, 4, 2, 1, 0, 1, 2}
	expected := [][]float64{
		{0, 0},
		{0, 2},
		{0, 2},
		{2, 0.8},
		{3.6, 0},
		{7.4, -1},
		{10.1, -1.5},
		{6.4, -0.8},
		{1.2, 0},
	}
	for i, v := range valx {
		sp.Update(v, valy[i])
		alpha, beta := sp.Value()
		if math.Abs(alpha-expected[i][0]) > 1e-3 {
			t.Errorf("Expected alpha value %.3f, but got %.3f, at index %d\n", expected[i][0], alpha, i)
		}
		if math.Abs(beta-expected[i][1]) > 1e-3 {
			t.Errorf("Expected beta value %.3f, but got %.3f, at index %d\n", expected[i][1], beta, i)
		}
	}
}

func BenchmarkLinRegUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	sp, err := NewLinReg(window)
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
