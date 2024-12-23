package library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
	fmt.Println("book added")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("book not found")
		return
	}
	if book.IsBorrowed {
		fmt.Println("book alredy borrowed")
		return
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("book borrewed")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("book not found")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("book not borrowed")
		return
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("book returned")
}

func (l *Library) ListBooks() {
	fmt.Println("available books")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}
