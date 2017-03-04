package main

import "fmt"

func main() {
	N := 10001
	n := 1
	for x := 0; n < N; x++ {
		y := x / 2
		for i := 2; i <= y; i++ {

			if x%i == 0 {

				break

			} else if i == y {
				n++
				if n == N-1 {
					fmt.Println(x)
				}
			}
		}
	}
}

//Projekt Euler problem no. 7:
//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
//we can see that the 6th prime is 13.
//What is the 10 001st prime number?
