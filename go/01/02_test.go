package main

import (
	"fmt"
	"testing"
)

type test struct {
	value  int
	answer int
}
type tableTest struct {
	tests []test
}

func TestEx2(t *testing.T) {
	table := tableTest{
		[]test{
			{
				value:  -5,
				answer: 0,
			},
			{
				value:  0,
				answer: 0,
			},
			{
				value:  1,
				answer: 1,
			},
			{
				value:  2,
				answer: 2,
			},
			{
				value:  3,
				answer: 6,
			},
			{
				value:  10,
				answer: 3628800,
			},
		},
	}
	for _, test := range table.tests {
		answer := Ex2(test.value)
		if test.answer != answer {
			t.Errorf("Wanted %d but had given %d", test.answer, answer)
		}
	}
}

func ExampleEx2() {
	fmt.Println(Ex2(8))
	// Output: 40320
}
