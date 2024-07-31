package main

import "testing"

func BenchmarkEx3(b *testing.B) {
	var n = 5
	for i := 0; i < b.N; i++ {
		wanted := map[int]int{
			1: 1,
			2: 4,
			3: 9,
			4: 16,
			5: 25,
		}
		result := Ex3(n)
		for j := 0; j < n; j++ {
			if result[j] != wanted[j] {
				b.Errorf("Wanted %d but given %d", wanted[j], result[j])
			}
		}
	}
}
