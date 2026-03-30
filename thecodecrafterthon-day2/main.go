package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ()

func decimalToOthers(s string, base int) string {
	val, _ := strconv.ParseInt(s, base, 64)
	hexaNumber := strconv.FormatInt(val, 16)
	binNumber := strconv.FormatInt(val, 2)
	output, _ := fmt.Printf("The converted Decimal is \n Binary: %v \n HexaDecimal: %v \n \n", binNumber, strings.ToUpper(hexaNumber))
	return strconv.Itoa(output)

}

func hexToDecimal(s string, base int) int64 {
	decimal, _ := strconv.ParseInt(s, base, 64)
	return decimal
}

func binToDecimal(s string, base int) int64 {
	decimal, _ := strconv.ParseInt(s, base, 64)
	return decimal
}

func main() {
	reader := bufio.NewReader(os.Stdin)
start:
	fmt.Println("       Welcome!")
	fmt.Print(" Select Base to covert\n 1. Decimal to Binary and HexaDecimal\n 2. HexaDecimal to Decimal\n 3. Binary to Decimal\n 4. Exit\n")
	fmt.Print("Pick any one to continue: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	c_choice, _ := strconv.Atoi(choice) // c_choice stands for cleaned choice after Atoi

	for {
		switch c_choice {
		case 1:
		case1Start:
			fmt.Print("Enter a valid Decimal Number to convert to HexaDecimal and Binary: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Println()
				fmt.Print("It's an Empty string, Enter a valid Decimal number\n \n")
				goto case1Start
			}

			decimalToOthers(number, 10)

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
			fmt.Println()
			fmt.Print("Enter a valid HexaDecimal Number to convert to Decimal: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Println()
				fmt.Print("It's an Empty string, Enter a valid HexaDecimal number\n \n")
				goto case2Start
			}

			base := 16
			fmt.Printf("The Decimal result of %q base %d is %d\n", number, base, hexToDecimal(number, base))

			fmt.Println()
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
			fmt.Println()
			fmt.Print("Enter a valid Binary Number to convert to Decimal: ")
			number, _ := reader.ReadString('\n')
			number = strings.TrimSpace(number)

			if number == "" {
				fmt.Print("It's an Empty string, Enter a valid HexaDecimal number\n \n")
				goto case3Start
			}

			base := 2
			fmt.Printf("The Decimal result of %q base %d is %d\n", number, base, binToDecimal(number, base))

		choose_choice:
			fmt.Println()
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
