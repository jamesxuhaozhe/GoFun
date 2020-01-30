package main

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d, %d); "+"got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkTriangle(b *testing.B) {
	aa := 30000
	bb := 40000
	cc := 50000

	for i := 0; i < b.N; i++ {
		if actual := calcTriangle(aa, bb); actual != cc {
			b.Errorf("calcTriangle(%d, %d); "+"got %d; expected %d", aa, bb, actual, cc)
		}
	}
}
