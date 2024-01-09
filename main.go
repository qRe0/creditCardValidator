package main

import (
	b "bufio"
	f "fmt"
	"os"
	r "regexp"
	s "strings"
)

// Function to validate credit card number bases on the pattern
func isCreditCardValid(ccn string) bool {
	// Regex pattern to validate credit card number (16 digits or 4 groups of 4 digits separated by space)
	pattern := r.MustCompile(`^(\d{16})$|^(\d{4}\s\d{4}\s\d{4}\s\d{4})$`)

	// Rerun true if the credit card number matches the pattern, else return false
	return pattern.MatchString(ccn)
}

// Function to validate credit card number using Luhn algorithm
func luhnAlgorithm(ccn string) bool {
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
	ans := sum % 10 == 0
	return ans
}

func main() {
	f.Println("This program will validate your credit card number")
	f.Print("Please enter your Credit Card number: ")
	reader := b.NewReader(os.Stdin)
	ccn, _ := reader.ReadString('\n')

	// Removing symbol '\n' from the end of the string
	ccn = s.TrimSpace(ccn)

	if !isCreditCardValid(ccn) {
		f.Println("------------------------------------")
		f.Println("Your credit card number is NOT valid")
		f.Println("------------------------------------")
		f.Println()
		return
	}

	option := -1
	for option != 0 {
		f.Println("Options:")
		f.Println("1. Check if the credit card number is valid using Luhn algorithm")
		f.Println("0. Exit")

		f.Print("Please, choose an option: ")
		f.Scanln(&option)

		switch option {
		case 1:
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
		case 0:
			f.Println("--------")
			f.Println("Goodbye!")
			f.Println("--------")
		default:
			f.Println("--------------")
			f.Println("Unknown option")
			f.Println("--------------")
			f.Println()
		}
	}
}
