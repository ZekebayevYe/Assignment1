package main

import (
	"Assignment1/Library"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	lib := library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for{
		fmt.Println("Choose:")
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")

		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option{
		case "1":
			fmt.Println("Enter book info (ID, Title, Author):")
			scanner.Scan()
			input := strings.Split(scanner.Text(),",")
			if len(input)!=3{
				fmt.Println("Invalid input")
				continue
			}
			lib.AddBook(library.Book{
				ID:         strings.TrimSpace(input[0]),
				Title:      strings.TrimSpace(input[1]),
				Author:     strings.TrimSpace(input[2]),
				IsBorrowed: false,
			})
		case "2":
			fmt.Println("Enter id to borrow")
			scanner.Scan()
			lib.BorrowBook(strings.TrimSpace(scanner.Text()))
		case "3":
			fmt.Println("Enter id to return:")
			scanner.Scan()
			lib.ReturnBook(strings.TrimSpace(scanner.Text()))
		case "4":
			lib.ListBooks()
		case "5":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
		
}

