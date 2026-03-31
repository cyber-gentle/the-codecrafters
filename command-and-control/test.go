package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Print("THE CALCULATOR: \n  Use for Arithmetic operations.\n  Type 'calc' to execute the calculator.  \n \n THE BASE CONVERTER: \n  Used to convert from one base to another. \n  Type 'base' to execute the base converter. \n \n The string transformer: \n  Used for transforming strings. \n  Type 'str' to execute string transformer \n \n")

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

	for {

		switch input {
		case "calc":
			//calculator()
			//goto start
			fmt.Println(input)

		case "base":
			// baseConverter()
			// goto start

		case "str":
			// stringTransformer()
			// goto start

		case "help":

			fmt.Println()
			goto start

		case "history":
			// history()
			// fmt.Println()
			// goto start

		case "exit":
			fmt.Print("Goodbye! \nExiting... \n")
			break

		default:
			fmt.Print("Type the correct command \n \n")
			goto start
		}
		break
	}

}
