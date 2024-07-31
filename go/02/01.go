package main

import (
	"fmt"
	"math"
)

/*
Question:
Write a program that computes the value of a+aa+aaa+aaaa with a given digit as the value of a.
Suppose the following input is supplied to the program:
9
Then, the output should be:
7380

Hint: a+aa+aaa+aaaa --> a + a ^ 1 + a ^ 2 + a ^ 3 + a ^ 4
*/

func main() {
	fmt.Println(Ex5(9))
}
func Ex5(a int) int {
	var result = 0
	for i:=1; i < 5; i++ {
		value := math.Pow(float64(a), float64(i))
		result = result + int(value)
	}
	return result
}