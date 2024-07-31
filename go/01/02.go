package main

import "fmt"

/*
Write a program which can compute the factorial of a given numbers.
The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/

//func main() {
//	fmt.Println(Ex2(8))
//}

func Ex2(number int) int {
	var result = 1
	if number < 0 {
		fmt.Println("Number should be positive")
		return 0
	} else if number == 0 {
		fmt.Println("Number should be not equal to zero")
		return 0
	} else {
		for i := 2; i < number+1; i++ {
			result *= i
		}
	}
	return result
}
