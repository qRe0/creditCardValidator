package main

import (
	b "bufio"
	f "fmt"
	io "io/ioutil"
	"os"
	r "regexp"
	s "strings"
)

// Function to validate credit card number bases on the pattern
func isCreditCardValid(ccn string) bool {
	// Remove \n from the end of the string
	ccn = s.TrimSpace(ccn)
	// Regex pattern to validate credit card number (16 digits or 4 groups of 4 digits separated by space)
	pattern := r.MustCompile(`^(\d{16})$|^(\d{4}\s\d{4}\s\d{4}\s\d{4})$`)

	// Rerun true if the credit card number matches the pattern, else return false
	return pattern.MatchString(ccn)
}

// Function to validate credit card number using Luhn algorithm
func luhnAlgorithm(ccn string) bool {
	ccn = s.TrimSpace(ccn)

	sum := 0
	alternate := false

	// Remove all spaces in the credit card number and reverse it to use Luhn algorithm
	ccn = s.ReplaceAll(ccn, " ", "")
	runes := []rune(ccn)
	for i := len(runes) - 1; i >= 0; i-- {
		n := int(runes[i] - '0')
		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		sum += n
		alternate = !alternate
	}
	ans := sum%10 == 0
	return ans
}

// Function to test the program with multiple valid credit card numbers from file
func multipleCCN(fileName string) []bool {
	content, err := io.ReadFile(fileName)
	if err != nil {
		f.Println("Unable to open file", fileName)
		return nil
	}

	lines := s.Split(string(content), "\n")
	validity := make([]bool, 0)

	for _, ccn := range lines {
		if isCreditCardValid(ccn) {
			validity = append(validity, luhnAlgorithm(ccn))
		} else {
			validity = append(validity, false)
		}
	}

	return validity
}

// Function to print the validity list
func printValidityList(list []bool) {
	for i, isValid := range list {
		f.Printf("%d. %s\n", i+1, func() string {
			if isValid {
				return "Valid"
			}
			return "NOT Valid"
		}())
	}
}

func main() {
	f.Println("This program will validate your credit card number")

	// region Maybe unused
	var fileName string
	var validity []bool
	// endregion

	option := -1
	for option != 0 {
		f.Println("Options:")
		f.Println("1. Check if the credit card number is valid using Luhn algorithm")
		f.Println("2. Check validity of CCN list (from file)")
		f.Println("0. Exit")

		f.Print("Please, choose an option: ")
		f.Scanln(&option)

		switch option {
		case 1:
			f.Print("Please enter your Credit Card number: ")
			reader := b.NewReader(os.Stdin)
			ccn, _ := reader.ReadString('\n')

			if !isCreditCardValid(ccn) {
				f.Println("------------------------------------")
				f.Println("Your credit card number is NOT valid (wrong format)")
				f.Println("------------------------------------")
				f.Println()
				return
			}

			if luhnAlgorithm(ccn) {
				f.Println("-------------------------------------------------")
				f.Println("Your credit card number is valid (Luhn algorithm)")
				f.Println("-------------------------------------------------")
			} else {
				f.Println("------------------------------------")
				f.Println("Your credit card number is NOT valid")
				f.Println("------------------------------------")
			}
			f.Println()
			break
		case 2:
			f.Print("Please enter the file name: ")
			f.Scanln(&fileName)

			validity = multipleCCN(fileName)
			f.Println("------------------------------------")
			f.Println("CCN validity (Luhn algorithm):")
			printValidityList(validity)
			f.Println()
			break
		case 0:
			f.Println("--------")
			f.Println("Goodbye!")
			f.Println("--------")
		default:
			f.Println("--------------")
			f.Println("Unknown option")
			f.Println("--------------")
			f.Println()
			break
		}
	}
}
