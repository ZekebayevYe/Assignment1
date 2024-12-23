package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Assignment1/Bank"
)

func main() {
	account := bank.BankAccount{AccountNumber: "2006", HolderName: "Yerassyl", Balance: 667.0}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")

		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Println("Enter deposit amount:")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			account.Deposit(amount)
		case "2":
			fmt.Println("Enter withdrawal amount:")
			scanner.Scan()
			amount, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			account.Withdraw(amount)
		case "3":
			account.GetBalance()
		case "4":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
