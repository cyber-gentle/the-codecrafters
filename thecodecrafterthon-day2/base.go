package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ()

func main() {
	reader := bufio.NewReader(os.Stdin)
start:
	fmt.Println("       Welcome!")
	fmt.Print(" Select Base to covert\n 1. Decimal to Binary and HexaDecimal\n 2. HexaDecimal to Decimal\n 3. Binary to Decimal\n 4. Exit\n")
	fmt.Print("Pick any one to continue: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	c_choice, _ := strconv.Atoi(choice) // c_choice stands for cleaned choice after Atoi
	fmt.Println()

	for {
		switch c_choice {
		case 1:
		case1Start:
			fmt.Print("Enter a valid Decimal Number : ")
			number, err := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Println()
				fmt.Print("It's an Empty string, Enter a valid Decimal number\n \n")
				goto case1Start
			}

			val, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Printf("%q is not a valid decimal.\n \n", number)
				continue
			}
			hexN := strconv.FormatInt(val, 16)
			binN := strconv.FormatInt(val, 2)

			fmt.Printf(" HexaDecimal: %v \n Binary: %v \n \n", strings.ToUpper(hexN), binN)

			fmt.Print("Do you have another Decimal number to convert?\n 1. Yes\n 2. No, Quit\n 3. Go to Menu\n Select choice here: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice_number, _ := strconv.Atoi(choice) // cleaned choice after atoi

			if choice_number == 1 {
				fmt.Println()
				continue
			} else if choice_number == 2 {
				return
			} else if choice_number == 3 {
				fmt.Println()
				goto start
			}

		case 2:
		case2Start:
			fmt.Print("Enter a valid HexaDecimal Number: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Println()
				fmt.Print("It's an Empty string, Enter a valid HexaDecimal number\n \n")
				goto case2Start
			}

			result, err := strconv.ParseInt(number, 16, 64)
			if err != nil {
				fmt.Printf("%q is not valid hex.\n", number)
				continue
			}
			fmt.Printf("Decimal: %v \n \n", result)

			fmt.Print("Do you have another HexaDecimal number to convert?\n 1. Yes\n 2. No, Quit\n 3. Go to Menu\n Select choice here: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice_number, _ := strconv.Atoi(choice) // cleaned choice after atoi

			if choice_number == 1 {
				fmt.Println()
				continue
			} else if choice_number == 2 {
				return
			} else if choice_number == 3 {
				fmt.Println()
				goto start
			}

		case 3:
		case3Start:
			fmt.Print("Enter a valid Binary Number: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Print("It's an Empty string, Enter a valid HexaDecimal number\n \n")
				goto case3Start
			}

			result, err := strconv.ParseInt(number, 2, 64)
			if err != nil {
				fmt.Printf("%q is not valid binary.\n \n", number)
				continue
			}
			fmt.Printf("Binary: %v \n \n", result)

		choose_choice:
			fmt.Print("Do you have another Binary number to convert\n 1. Yes\n 2. No, Quit\n 3. Go to Menu\n Select choice here: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)
			choice_number, _ := strconv.Atoi(choice) // cleaned choice after atoi
			if choice == "" {
				fmt.Print("It's an Empty string, Choose from the options above\n \n")
				goto choose_choice

			}

			if choice_number == 1 {
				fmt.Println()
				continue
			} else if choice_number == 2 {
				return
			} else if choice_number == 3 {
				fmt.Println()
				goto start
			}

		case 4:
			fmt.Println("Exiting, Goodbye!")

		default:
			fmt.Print("Select from the options above\n \n")
			goto start

		}
		break

	}

}
