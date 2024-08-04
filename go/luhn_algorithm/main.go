package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func CheckDigits(s string) bool {
	var sr = strings.Split(s, "")
	var st []int
	var res int
	var err error
	for i := 0; i < len(sr); i++ {
		res, err = strconv.Atoi(sr[i])
		if err != nil {
			log.Fatal(err)
		}
		if i % 2 == 0 {	
			st = append(st, res * 2)
		}else{
			st = append(st, res)
		}
	}
	for index, num := range st {
		if len(fmt.Sprintf("%d", num)) > 1 {
			strNum := strconv.Itoa(st[index])
			sum := 0
			for _, digit := range strNum {
				dig, err := strconv.Atoi(string(digit))
				if err != nil {
					log.Fatal(err)
				}
				sum += dig
			}
			st[index] = sum
		}
	}
	var sum int
	for _, num := range(st) {
		sum += num
	}
	
	if sum % 10 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var cardNumber = "5892101519899518"
	result := CheckDigits(cardNumber)
	if result == true {
		fmt.Println("Its a valid card number, Welcome")
	} else {
		fmt.Println("It isnt a valid card number")
	}

}