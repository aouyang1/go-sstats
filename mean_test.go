package sstats

import (
	"testing"
)

func TestMeanUpdate(t *testing.T) {
	m, err := NewMean(5)
	if err != nil {
		t.Fatal(err)
	}
	vals := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []float64{1, 1.5, 2, 2.5, 3, 4, 5, 6, 7}
	for i, v := range vals {
		m.Update(v)
		mean := m.Value()
		if mean != expected[i] {
			t.Errorf("Expected value %.3f, but got %.3f\n", expected[i], mean)
			continue
		}
	}
}

func BenchmarkMeanUpdate(b *testing.B) {
	window := 1000
	numValues := 100000
	m, err := NewMean(window)
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
