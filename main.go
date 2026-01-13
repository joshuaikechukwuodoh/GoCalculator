package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Println("Error reading first number:", err)
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Println("Error reading second number:", err)
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Error reading operator:", err)
		return
	}

	result, err := calculator(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}

// Calculator function
func calculator(num1 int, num2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator")
	}
}
