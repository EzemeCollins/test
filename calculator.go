package main

import "fmt"

func main() {
	var bouble1 float64
	var bouble2 float64
	var operator string
	var result float64

	fmt.Println("input first number:1")
	fmt.Scan(&bouble1)

	fmt.Println("input second number:2")
	fmt.Scan(&bouble2)

	fmt.Println("choose your arithmetic sign /,*,+,-,")
	fmt.Scan(&operator)

	switch operator {
	case "/":
		result = bouble1 / bouble2
		fmt.Printf("Your answer is %v\n", result)

	case "*":
		result = bouble1 * bouble2
		fmt.Printf("Your answer is %v\n", result)

	case "+":
		result = bouble1 + bouble2
		fmt.Printf("Your answer is %v\n", result)

	case "-":
		result = bouble1 - bouble2
		fmt.Printf("Your answer is %v\n", result)

	default:
		fmt.Printf("invalid response!!")
	}

}
