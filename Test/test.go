package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Print("-----------HELP MENU-----------\n \n")
	fmt.Print("THE CALCULATOR: \n  Use for Arithmetic operations.\n  Type 'calc' to execute the calculator.  \n \n")
	fmt.Print("THE BASE CONVERTER: \n  Used to convert from one base to another. \n  Type 'base' to execute the base converter. \n \n")
	fmt.Print("THE STRING TRANSFORMER: \n  Used for transformation of strings. \n  Type 'str' to execute string transformer \n \n")
	fmt.Print("HISTORY:  \n  Show")

}

func calculator() {
	fmt.Println("WELCOME TO SENTINEL'S CALCULATOR")
	fmt.Println("")

}

func baseConverter() {
	fmt.Println("WELCOME TO SENTINEL'S BASE CONVERTER")

	fmt.Println("")

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

		case "base":
			// baseConverter()
			// goto start

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
			break

		default:
			fmt.Print("Type the correct command \n \n")
			goto start
		}
		break
	}

}
