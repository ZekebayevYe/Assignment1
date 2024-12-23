package bank

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
	fmt.Printf("Deposited %.2f. New Balance: %.2f\n", amount, b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	if b.Balance >= amount {
		b.Balance -= amount
		fmt.Printf("Withdrawn %.2f. New Balance: %.2f\n", amount, b.Balance)
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func (b *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, transaction := range transactions {
		if transaction > 0 {
			account.Deposit(transaction)
		} else {
			account.Withdraw(-transaction)
		}
	}
}
