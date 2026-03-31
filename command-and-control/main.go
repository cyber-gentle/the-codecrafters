// CodeCrafters — Hackathon 002
// Squad: Gophers
// Members: David Abraham, Obeko Eunice, Ugwu Chioma, Michael Bulus, Emmanuel Eliagwu, Akatu Worthy, Samuel Jireh

package main

import (
	"fmt"
	// "math"
	//"strconv"
	"strings"
	"unicode"
)

func main() {
	var first float64
	var second float64

	fmt.Println("...GOPHER'S CALC...")
Start1:
	fmt.Println("Input first number")
	_, err := fmt.Scanln(&first)
	if err != nil {
		fmt.Print("Enter digit only!\n")
		goto Start1
	}

// Start2:
// 	fmt.Println("Input second number")
// 	fmt.Scanln(&second)
// 	if err != nil {
// 		fmt.Print("Enter digit only!\n")
// 		goto Start2
// 	}

// 	var Operator int

// 	fmt.Println("1: Addition| 2: Subtraction| 3: Multiplication| 4: Division| 5: Remainder | 6: Power | 7: Last | 8: Exit")
// 	fmt.Scanln(&Operator)

// 	if Operator < 1 || Operator > 10 {
// 		fmt.Println("Not a valid Operator")
// 	}

// 	for {
// 		switch Operator {
// 		case 1:
// 			result := first + second
// 			fmt.Println(result)
// 			continue

// 		case 2:
// 			result := first - second
// 			fmt.Println(result)
// 			continue

// 		case 3:
// 			result := first * second
// 			fmt.Println(result)
// 			continue

// 		case 4:
// 			if second == 0 {
// 				fmt.Println("Not divisible by zero")
// 			}
// 			result := first / second
// 			fmt.Println(result)
// 			continue

// 		case 5:
// 			result := int64(first) % int64(second)
// 			fmt.Println(result)
// 			continue

// 		case 6:
// 			result := math.Pow(first, second)
// 			fmt.Println(result)
// 			continue

// 		case 8:
// 			fmt.Println("Exiting...")
// 			break

// 		default:
// 			fmt.Print("Out of operation range\n", "Choose from the above\n")
// 			continue
// 		}
// 		break
// 	}
// }

// func main() {
// 	var hex, dec, bin string

// 	var base int
// start:
// 	fmt.Println("1: Dec")
// 	fmt.Println("2: Hex")
// 	fmt.Println("3: Bin")
// 	fmt.Println("4: Exit")
// 	fmt.Scanln(&base)

// 	for {
// 		switch base {
// 		case 1:
// 			fmt.Print("ENTER DEC NUMBER: ")
// 			fmt.Scanln(&dec)

// 			dec, err := strconv.ParseInt(dec, 10, 64)
// 			if err != nil {
// 				fmt.Println("Invalid dec")
// 			}
// 			hexN := strconv.FormatInt(dec, 16)
// 			binN := strconv.FormatInt(dec, 2)
// 			fmt.Printf("HexaDecimal: %v\n", strings.ToUpper(hexN))
// 			fmt.Printf("Binary: %v\n", binN)
// 			continue

// 		case 2:

// 			fmt.Print("Enter Number: ")
// 			fmt.Scanln(&hex)
// 			dec, err := strconv.ParseInt(hex, 16, 64)
// 			if err != nil {
// 				fmt.Println("Invalid Hex")
// 			}
// 			fmt.Println(dec)
// 			continue

// 		case 3:

// 			fmt.Print("Enter Number: ")
// 			fmt.Scanln(&bin)

// 			dec, err := strconv.ParseInt(bin, 2, 64)
// 			if err != nil {
// 				fmt.Println("Invalid Bin")
// 			}
// 			fmt.Println(dec)
// 			continue

// 		case 4:
// 			fmt.Println("Exiting...")
// 			return

// 		default:
// 			fmt.Println("Invalid Option")
// 			fmt.Println("Choose A Valid Number")
// 			goto start
// 		}

// 	}
// }


func reverse(word string)string{
	words := strings.Fields(word)
	for i, word := range words{
		runes := []rune(word)
		for l, r := 0, len(word)-1; l < r; l , r = l+1,r-1{
			runes[r], runes[l] = runes[l], runes[r]
		}
		words[i] = string(runes)
	}

	return strings.Join(words," ")
	
}

func snakeCase(word string)string{
	var result strings.Builder

		for _, ch := range word{
			if unicode.IsLetter(ch) || unicode.IsDigit(ch) || unicode.IsSpace(ch) || ch == '_'{
				if unicode.IsLetter(ch){
					result.WriteRune(unicode.ToLower(ch))
				}else{
					result.WriteRune(ch)
				}
			}
		}
		word = result.String()
		word = strings.Join(strings.Fields(word), "_")
		return word
		
	}

func title(word string)string{
	word = strings.ToLower(word)
	return strings.Title(word)
}

var smallWords = []string{
	"a", "an", "the", "and", "but", "or", "for", "nor",
	"on", "at", "to", "by", "in", "of", "up", "as",
	"is", "yet", "it",
}

func checkSmallWords(word string)bool{
	for _, w := range smallwords{
		if w == word {
			return  true
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

func stringTransformer() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("SENTINEL STRING TRANSFORMER — ONLINE")

		fmt.Print(">> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			fmt.Println("Please enter a valid command -- Usage: upper <text>")
			continue
		}

		textInput := strings.Fields(line)
		command := strings.ToLower(textInput[0])

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
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

func main(){
	fmt.Println(reverse("Lagos Nigeria"))
	fmt.Println(snakeCase("Alert! Level 5 detected."))
	fmt.Println(title("michael SAMUEL"))
}

