package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	num1, num2 int // integer storer
	p_choice   int // proceed choice
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("CHOOSE AN OPERATION TO EXECUTE\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit\n Type %q for more info\n", "HELP")
	fmt.Print("Select an operation to continue: ")
	op_c, _ := reader.ReadString('\n')
	op_c = strings.TrimSpace(op_c)
	op, _ := strconv.Atoi(op_c)

	for {

		switch op {
		case 1:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 + num2
			fmt.Printf("%d + %d is = %d\n", num1, num2, result)
			fmt.Println()

			fmt.Print("Do you want add another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else {
				fmt.Println("Choose from the options above")
			}

		case 2:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 - num2
			fmt.Printf("%d - %d is = %d\n", num1, num2, result)

			fmt.Print("Do you want subtract another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else {
				fmt.Println("Choose from the options above")
			}

		case 3:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 * num2
			fmt.Printf("%d x %d is = %d\n", num1, num2, result)

			fmt.Print("Do you want multiply another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else {
				fmt.Println("Choose from the options above")
			}

		case 4:
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			if num2 == 0 {
				fmt.Println("Divisor can't be zero")
				continue
			}
			result := num1 / num2
			fmt.Printf("%d ÷ %d is = %d\n", num1, num2, result)

			fmt.Print("Do you want divide another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else {
				fmt.Println("Choose from the options above")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		}
		if op_c == "HELP" {
			fmt.Println(" Addition: add your input together\n Subtraction: Minus a number from another\n Multiplication: Multiply numbers to get result\n Division: Divides a number by the other\n ")
			return
		}
	}
}
