package main

import "fmt"

func fizzbuzz() {
	countTo := 100
	for i := 1; i <= countTo; i++ {
		if i%15 == 0 { //if it is divisible by both 3 and 5; same as if i%3 == 0 && i%5 == 0
			fmt.Println("Fizz Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
