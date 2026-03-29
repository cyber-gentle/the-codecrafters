package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	num1, num2 float64 // integer storer
	p_choice   int     // proceed choice whether to do another calculation or not
	// op_c // operation choice
)

func main() {

	reader := bufio.NewReader(os.Stdin)
start:
	fmt.Printf("CHOOSE AN OPERATION TO EXECUTE\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit\n 6. More info\n")
	fmt.Print("Select an operation to continue: ")
	op_c, _ := reader.ReadString('\n')
	op_c = strings.TrimSpace(op_c)
	op, _ := strconv.Atoi(op_c)

	for {

		switch op {
		case 1:

		case1FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case1FirstNumber
			}
		case1SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case1SecondNumber
			}
			num1 = float64(num1)
			num2 = float64(num2)
			result := num1 + num2
			fmt.Printf("The sum of %g + %g is %g\n \n", num1, num2, result)

			fmt.Print("Do you want add another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 2:
		case2FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case2FirstNumber
			}
		case2SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case2SecondNumber
			}
			num1 = float64(num1)
			num2 = float64(num2)
			result := num1 - num2
			fmt.Printf("The difference of %g and %g is %g\n \n", num1, num2, result)

			fmt.Print("Do you want subtract another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 3:
		case3FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case3FirstNumber
			}
		case3SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case3SecondNumber
			}
			num1 = float64(num1)
			num2 = float64(num2)
			result := num1 * num2
			fmt.Printf("The product of %g and %g is %g\n \n", num1, num2, result)

			fmt.Print("Do you want multiply another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 4:
		case4FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case4FirstNumber
			}
		case4SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Print("Enter a number and not alphabet\n \n")
				goto case4SecondNumber
			}
			if num2 == 0 {
				fmt.Print("Divisor can't be zero\n \n")
				goto case4SecondNumber
			}
			num1 = float64(num1)
			num2 = float64(num2)
			result := num1 / num2
			fmt.Printf("The qoutient of %g and %g is %g\n \n", num1, num2, result)

			fmt.Print("Do you want divide another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		case 6:
			fmt.Println()
			fmt.Print(" Addition: add your input together\n Subtraction: Minus a number from another\n Multiplication: Multiply numbers to get result\n Division: Divides a number by the other\n \n ")
			goto start

		default:
			fmt.Println("Enter a valid operation: ")
			fmt.Println()
			goto start

		}

	}
}
