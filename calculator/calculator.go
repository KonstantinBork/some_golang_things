package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		fmt.Println("Not the right amount of arguments")
		os.Exit(1)
	}

	operator := args[0]
	arg1, err1 := strconv.ParseFloat(args[1], 64)

	if err1 != nil {
		fmt.Println("First argument is not a number")
		os.Exit(1)
	}

	arg2, err2 := strconv.ParseFloat(args[2], 64)
	if err2 != nil {
		fmt.Println("Second argument is not a number")
		os.Exit(1)
	}

	switch operator {
	case "+":
		fmt.Println(add(arg1, arg2))
	case "-":
		fmt.Println(subtract(arg1, arg2))
	case "*":
		fmt.Println(multiply(arg1, arg2))
	case "/":
		fmt.Println(divide(arg1, arg2))
	default:
		fmt.Printf("Operator %s not allowed\n", operator)
		os.Exit(1)
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Division by 0 is not allowed")
		os.Exit(1)
	}
	return a / b
}
