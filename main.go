//Mahmoud Sakalla
//Simple Calculator Program, this takes two numbers and performs the desired operation

package main

import "fmt"

func main() {

	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator using Golang")
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			//make sure that a denominator cannot be zero
			fmt.Println("Error: Division by zero.")
		}
	default:
		fmt.Println("Invalid operator, choose from the options provided")
	}
}
