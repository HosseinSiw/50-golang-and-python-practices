package main

import (
	"strconv"
	"strings"
)

/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5,
between 2000 and 3200 (both included).  The numbers obtained should be printed in a comma-separated sequence on a single line.
*/

//func main() {
//	fmt.Println(Exercise01(2000, 3200))
//}

func Exercise01(low, high int) string {
	var myResult []string
	for i := low; i < high; i++ {
		if i%7 == 0 && i%5 != 0 {
			myResult = append(myResult, strconv.Itoa(i))
		}
	}
	return strings.Join(myResult, ",")
}
