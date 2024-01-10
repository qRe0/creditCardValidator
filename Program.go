package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("This program will validate your credit card number")
	fmt.Println("Loading...")

	time.Sleep(1500 * time.Millisecond)

	var fileName string
	var validity map[string]map[string]bool

	option := -1
	for option != 0 {
		fmt.Println("Options:")
		fmt.Println("1. Check if the credit card number is valid using Luhn algorithm")
		fmt.Println("2. Check validity of CCN list (from file)")
		fmt.Println("0. Exit")

		fmt.Print("Please, choose an option: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Print("Please enter your Credit Card number: ")
			reader := bufio.NewReader(os.Stdin)
			ccn, _ := reader.ReadString('\n')

			if !isCreditCardValid(ccn) {
				fmt.Println("------------------------------------")
				fmt.Println("Your credit card number is NOT valid (wrong format)")
				fmt.Println("------------------------------------")
				fmt.Println()
			}

			if luhnAlgorithm(ccn) {
				fmt.Println("-------------------------------------------------")
				fmt.Println("Your credit card number is valid (Luhn algorithm)")
				fmt.Println("-------------------------------------------------")
			} else {
				fmt.Println("------------------------------------")
				fmt.Println("Your credit card number is NOT valid")
				fmt.Println("------------------------------------")
			}
			fmt.Println()
		case 2:
			fmt.Print("Please enter the file name: ")
			fmt.Scanln(&fileName)

			validity = readCCNFromFile(fileName)
			fmt.Println("------------------------------------")
			fmt.Println("CCN validity (Luhn algorithm):")
			printValidityList(validity)
			fmt.Println()
		case 0:
			fmt.Println("--------")
			fmt.Println("Goodbye!")
			fmt.Println("--------")
		default:
			fmt.Println("--------------")
			fmt.Println("Unknown option")
			fmt.Println("--------------")
			fmt.Println()
		}
	}
}
