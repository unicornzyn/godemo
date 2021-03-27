package main

import "fmt"

func sqrt(num float64) float64 {
	currguess := 1.0
	prevguess := 0.0

	for count := 1; count <= 10; count++ {
		prevguess = currguess
		currguess = prevguess - (prevguess*prevguess-num)/(2*prevguess)
		if currguess == prevguess {
			break
		}
		fmt.Println("A guess for square root is ", currguess)
	}
	return currguess
}

func main() {
	var num float64 = 25
	fmt.Println("Square root is:", sqrt(num))
}
