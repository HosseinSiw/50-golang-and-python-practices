package main

import (
	"fmt"
	"math"
)

/*
Exercise 003

With a given integral number n,
write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included),
and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8
Then, the output should be:
map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]
*/
func main() {
	fmt.Println(Ex3(8))

}
func Ex3(n int) map[int]int {
	var result = make(map[int]int, n)
	for i := 1; i < n+1; i++ {
		result[i] = int(math.Pow(float64(i), float64(2)))
	}

	return result
}
