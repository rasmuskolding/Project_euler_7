package main

import "fmt"

func main() {
	N := 10001
	n := 1
	for x := 0; n < N; x++ { //x is the first number we check to see if it is prime)
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
