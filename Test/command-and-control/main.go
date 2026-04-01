package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
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

	if input == "" {
		fmt.Print("Enter base and a valid number of the base\n \n")
		goto start
	}

	if len(user_input) != 2 {
		fmt.Print("Invalid base command\n Enter a base and a valid number\n \n")
		goto start
	}

	for {
		switch input_base {
		case "dec":
			decimal_value, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				fmt.Printf(" %q is not a valid Decimal.\n Enter a valid Decimal Number.\n \n", number)
				goto start
			}
			bin_Number := strconv.FormatInt(decimal_value, 2)
			hex_Number := strings.ToUpper(strconv.FormatInt(decimal_value, 16))

			fmt.Printf(" ✦ Binary : %v \n ✦ Hex    : %v\n \n", bin_Number, hex_Number)
			goto start

		case "bin":
			bin_Number, err := strconv.ParseInt(number, 2, 64)
			if err != nil {
				fmt.Printf(" %q is not a valid binary.\n Enter a valid Binary Number.\n \n", number)
				goto start
			}

			fmt.Printf(" ✦ Decimal : %v \n \n", bin_Number)
			goto start

		case "hex":
			hex_Number, err := strconv.ParseInt(number, 16, 64)
			if err != nil {
				fmt.Printf(" %q is not a valid hex.\n Enter a valid Hexa-Decimal Number.\n \n", strings.ToUpper(number))
				goto start
			}

			fmt.Printf(" ✦ Decimal : %v \n \n", hex_Number)
			goto start

		case "go to menu":
			fmt.Print("Returning to Menu\n \n")
			continue
			//return

		default:
			fmt.Print("Enter a base and a valid number\n \n")
			goto start

		}
		break
	}

}

func upper(word string) string {
	return strings.ToUpper(word)
}

func lower(word string) string {
	return strings.ToLower(word)
}

func cap(word string) string {
	words := strings.Fields(word)

	for i, ch := range words {
		runes := []rune(strings.ToLower(ch))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)

	}
	return strings.Join(words, " ")
}

var smallWords = []string{
	"a", "an", "the", "and", "but", "or", "for", "nor",
	"on", "at", "to", "by", "in", "of", "up", "as",
	"is", "yet", "it",
}

func checkSmallWords(word string) bool {
	for _, w := range smallWords {
		if w == word {
			return true
		}
	}
	return false
}

func title(word string) string {
	words := strings.Fields(word)

	for i, w := range words {
		lower := strings.ToLower(w)

		if i == 0 || !checkSmallWords(lower) {
			runes := []rune(lower)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		} else {
			words[i] = lower
		}
	}

	return strings.Join(words, " ")
}

func snake(word string) string {
	var result strings.Builder

	for _, ch := range word {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || unicode.IsSpace(ch) {
			if unicode.IsLetter(ch) {
				result.WriteRune(unicode.ToLower(ch))
			} else {
				result.WriteRune(ch)
			}
		}
	}

	word = result.String()
	word = strings.Join(strings.Fields(word), "_")
	return word
}

func reverse(word string) string {
	words := strings.Fields(word)
	for i, word := range words {
		runes := []rune(word)
		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}
		words[i] = string(runes)

	}
	return strings.Join(words, " ")
}

func stringTransformer() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")
		fmt.Print("Enter 'Go to Menu' to Return to SENTINEL COMMAND AND CONTROL MENU\n \n")

		fmt.Print(">> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Println("Please enter a valid command -- Usage: upper <text>")
			continue
		}

		line = strings.ToLower(line)
		textInput := strings.Fields(line)
		command := strings.ToLower(textInput[0])

		if line == "go to menu" {
			fmt.Print("Shutting down String Transformer. Goodbye.\n \n")
			return
		}

		if len(textInput) < 2 {
			fmt.Printf("No text provided. Usage: %s text\n\n", command)
			continue
		}

		text := strings.Join(textInput[1:], " ")

		switch command {
		case "upper":
			fmt.Println("→ ", upper(text))

		case "snake":
			fmt.Println("→ ", snake(text))

		case "lower":
			fmt.Println("→ ", lower(text))

		case "reverse":
			fmt.Println("→ ", reverse(text))

		case "cap":
			fmt.Println("→ ", cap(text))

		case "title":
			fmt.Println("→ ", title(text))

		default:
			fmt.Printf("Unknown command: %q\nValid commands: upper, lower, cap, title, snake, reverse, exit\n", command)
		}
		fmt.Println()
	}
}

func calculator() {

	var history [5]string
	var index int = 0
	var lastResult float64

start:
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("WELCOME TO SENTINEL'S CALCULATOR")
	fmt.Print("Enter 'Go to Menu' to Return to SENTINEL COMMAND AND CONTROL MENU\n  \n")

	fmt.Print("Type Here: ")

	scanner.Scan()
	input := scanner.Text()
	input = strings.ToLower(input)

	if input == "" {
		fmt.Print("Enter operator and a valid set of number\n \n")
		goto start
	}

	if input == "go to menu" {
		fmt.Print(" Shutting down  SENTINEL CONTROL CALCULATOR.\n Goodbye.\n \n")
		return
	} else if input == "last" {
		fmt.Println("Last Result: ", lastResult)
	} else if input == "history" {
		fmt.Println("Last 5 Calculations: ")
		for i := 0; i < 5; i++ {
			if history[i] != "" {
				fmt.Println(history[i])
				fmt.Println()
			}
		}
		goto start
	}

	user_input := strings.Fields(input)
	user_operator := user_input[0]
	user_operator = strings.ToLower(user_operator)

	user_Num1, err := strconv.Atoi(user_input[1])
	if err != nil {
		fmt.Print("Enter Digit Only!\n \n")
	}
	user_Num2, err := strconv.Atoi(user_input[2])
	if err != nil {
		fmt.Print("Enter Digit Only!\n \n")
	}

	if len(user_input) != 3 {
		fmt.Print("Enter operator and a valid set of number\n \n")
		goto start
	}

	for {

		num1 := float64(user_Num1)
		num2 := float64(user_Num2)

		switch user_operator {
		case "add":
			result := num1 + num2
			fmt.Printf("✦ Result: %g\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %g", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = result
			goto start

		case "sub":
			result := num1 - num2
			fmt.Printf("✦ Result: %g\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %g", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = result
			goto start

		case "mul":
			result := num1 * num2
			fmt.Printf("✦ Result: %g\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %g", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = result
			goto start

		case "div":
			if user_Num2 == 0 {
				fmt.Print("Not divisible by zero\n \n")
				goto start
			}

			result := num1 / num2
			fmt.Printf("✦ Result: %g\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %g", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = result
			goto start

		case "mod":
			result := int64(num1) % int64(num2)
			fmt.Printf("✦ Result: %d\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %d", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = float64(result)
			goto start

		case "pow":
			result := math.Pow(num1, num2)
			fmt.Printf("✦ Result: %g\n \n", result)

			entry := fmt.Sprintf("%s %g  %g = %g", user_operator, num1, num2, result)
			history[index%5] = entry
			index++
			lastResult = result
			goto start

		default:
			fmt.Print("Out of operation range\n", "Enter a valid command \n")
			goto start
		}

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
	input = strings.TrimSpace(input)
	fmt.Println()

	for {

		switch input {
		case "calc":
			calculator()
			goto start

		case "base":
			baseConverter()
			goto start

		case "str":
			stringTransformer()
			goto start

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
