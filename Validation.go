package main

import (
	"bufio"
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
func readCCNFromFile(fileName string) map[string]map[string]bool {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Unable to open file", fileName)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validity := make(map[string]map[string]bool)

	for scanner.Scan() {
		ccn := scanner.Text()
		if isCreditCardValid(ccn) {

			//{ Data structure
			//	"ccn": {
			//	"check option": value (true or false)
			//  },
			//	"4003 0241 0084 7010": {
			//	"luhn": true
			//  }
			//}

			// Store the validity of the credit card number based on Luhn algorithm
			// If the credit card number is valid by luhn algorithm, we have ccn as key, "luhn" as key and true as value
			validity[ccn] = make(map[string]bool)
			validity[ccn]["Luhn"] = luhnAlgorithm(ccn)
		} else {
			// Store the validity of the credit card number based on card format
			// If the credit card number is valid by card format, we have ccn as key, "card format" as key and true as value
			validity[ccn] = make(map[string]bool)
			validity[ccn]["Card format"] = false
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return validity
}

// Function to print the validity list
func printValidityList(list map[string]map[string]bool) {
	totalCards := 0
	validCards := 0
	invalidCards := 0

	for ccn, validity := range list {
		totalCards++

		fmt.Printf("Credit Card Number: %s\n", ccn)
		for check, result := range validity {
			var isValid string
			if result {
				isValid = "is valid"
				validCards++
			} else {
				isValid = "is NOT valid"
				invalidCards++
			}
			// Where check is "luhn" or "card format", isValid is "is valid" or "is NOT valid"
			fmt.Printf("- %s check: %s\n", check, isValid)
		}
		fmt.Println()
	}

	// P.S. %d\n means print the number as decimal and go to the next line
	fmt.Printf("Total Cards: %d\n", totalCards)
	fmt.Printf("Valid Cards: %d\n", validCards)
	fmt.Printf("Invalid Cards: %d\n", invalidCards)
}
