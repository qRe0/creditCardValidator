package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Function to validate credit card number bases on the pattern
func isCreditCardValid(ccn string) bool {
	// Remove \n from the end of the string
	ccn = strings.TrimSpace(ccn)
	// Regex pattern to validate credit card number (16 digits or 4 groups of 4 digits separated by space)
	pattern := regexp.MustCompile(`^(\d{16})$|^(\d{4}\s\d{4}\s\d{4}\s\d{4})$`)

	// Rerun true if the credit card number matches the pattern, else return false
	return pattern.MatchString(ccn)
}

// Function to validate credit card number using Luhn algorithm
func luhnAlgorithm(ccn string) bool {
	ccn = strings.TrimSpace(ccn)

	sum := 0
	alternate := false

	// Remove all spaces in the credit card number and reverse it to use Luhn algorithm
	ccn = strings.ReplaceAll(ccn, " ", "")
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
func readCCNFromFile(fileName string) []bool {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to open file", fileName)
		return nil
	}

	lines := strings.Split(string(content), "\n")
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
		if isValid {
			fmt.Println(i+1, "is valid")
		} else {
			fmt.Println(i+1, "is NOT valid")
		}
	}
}
