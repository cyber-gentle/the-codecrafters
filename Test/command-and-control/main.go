package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func help() {
	fmt.Print("-----------HELP MENU-----------\n \n")
	fmt.Print("THE CALCULATOR: \n  Use for Arithmetic operations.\n  Type 'calc' to execute the calculator.  \n \n")
	fmt.Print("THE BASE CONVERTER: \n  Used to convert from one base to another. \n  Type 'base' to execute the base converter. \n \n")
	fmt.Print("THE STRING TRANSFORMER: \n  Used for transformation of strings. \n  Type 'str' to execute string transformer \n \n")
	fmt.Print("HISTORY:  \n  Show the last 10 executed command\n \n")
	fmt.Print("Exit: \n")
	fmt.Print("Exit the program.\n \n \n \n")

}

func calculator() {
	fmt.Println("WELCOME TO SENTINEL'S CALCULATOR")
	fmt.Println("")

	// 	scanner.Scan()
	// input := scanner.Text()
	// input = strings.ToLower(input)
	// user_input := strings.Fields(input)
	// user_operator := user_input[0]
	// user_num1, err := strconv.Atoi(user_input[1])
	// 	if err != nil {
	// 		Invalid
	// 	}
	// user_Num2, err := strconv.Atoi(user_input[2])
	// 	if err != nil {

	// 	}

}

func baseConverter() {
	scanner := bufio.NewScanner(os.Stdin)

start:
	fmt.Println(" ════════════WELCOME TO SENTINEL'S BASE CONVERTER════════════")
	fmt.Print(" Enter the base to convert from followed by the number\n E.g dec 255, bin 101, hex 1F.\n Enter 'Go to Menu' to Return to SENTINEL COMMAND AND CONTROL MENU\n  \n")

	fmt.Print("Type Here: ")

	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)
	user_input := strings.Fields(input)
	input_base := user_input[0]
	number := user_input[1]

	for {
		switch input_base {
		case "dec":
			decimal_value, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Print("Invalid Decimal Number\n \n")
				goto start
			}
			bin_Number := strconv.FormatInt(decimal_value, 2)
			hex_Number := strings.ToUpper(strconv.FormatInt(decimal_value, 16))

			fmt.Printf(" ✦ Binary : %v \n ✦ Hex    : %v\n \n", bin_Number, hex_Number)
			goto start

		case "bin":
			bin_Number, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Print("Invalid Binary Number\n \n")
				goto start
			}

			fmt.Printf(" ✦ Binary : %v \n \n",bin_Number)
			goto start

		case "go to menu":
			fmt.Print("Returning to Menu\n \n")
			continue
			//return

		}
		break
	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

start:
	fmt.Print("════════════════════════════════════════════════\n  SENTINEL — COMMAND & CONTROL CONSOLE\n     All systems nominal. Type 'help' to begin.\n════════════════════════════════════════════════\nC&C>\n \n")
	fmt.Println("     calc   <command>   → the calculator")
	fmt.Println("     base   <command>   → the base converter")
	fmt.Println("     str    <command>   → the string transformer")
	fmt.Println("     help               → shows all commands")
	fmt.Println("     history            → shows last 10 inputs")
	fmt.Println("     exit               → shuts down the console")

	fmt.Print("Type Here: ")

	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)
	fmt.Println()

	for {

		switch input {
		case "calc":
			//calculator()
			//goto start

		case "base":
			baseConverter()
			goto start

		case "str":
			// stringTransformer()
			// goto start

		case "help":

			help()
			goto start

		case "history":
			// history()
			// fmt.Println()
			// goto start

		case "exit":
			fmt.Println("Exiting...")
			fmt.Print("Goodbye!\n \n")
			return

		default:
			fmt.Print("Enter a valid command or seek help.\n \n")
			goto start
		}
		break
	}

}
