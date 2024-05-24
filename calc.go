package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter First Number:")
	scanner.Scan()
	var num1 float64
	fmt.Sscan(scanner.Text(), &num1) 

	fmt.Println("Enter Operator (+, -, *, /):")
	scanner.Scan()
	var operator string
	operator = scanner.Text()

	fmt.Println("Enter Second Number:")
	scanner.Scan()
	var num2 float64
	fmt.Sscan(scanner.Text(), &num2) 

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid Operator")
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
