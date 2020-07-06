package sstats

import (
	"testing"
)

func TestSumUpdate(t *testing.T) {
	m, err := NewSum(5)
	if err != nil {
		t.Fatal(err)
	}
	vals := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []float64{1, 3, 6, 10, 15, 20, 25, 30, 35}
	for i, v := range vals {
		m.Update(v)
		val := m.Value()
		if val != expected[i] {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], val)
			continue
		}
	}
}

func BenchmarkSumUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	m, err := NewSum(window)
	if err != nil {
		b.Fatal(err)
	}

	for j := 0; j < b.N; j++ {
		for i := 0; i < numValues; i++ {
			m.Update(float64(i))
		}
		m.Reset()
	}
}
