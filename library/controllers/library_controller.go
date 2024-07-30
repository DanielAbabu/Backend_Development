// controllers/library_controller.go
package controllers

import (
	"bufio"
	"fmt"
	"library/models"
	"library/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	service services.LibraryManager
	reader  *bufio.Reader
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (lc *LibraryController) AddMember() {
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter the member's name: ")
	name, _ := lc.reader.ReadString('\n')
	name = strings.TrimSpace(name)

	member := models.Member{ID: 0, Name: name, BorrowedBooks: []models.Book{}}
	lc.service.AddMember(member)

	fmt.Println("Member added.")

	fmt.Println("\n", strings.Repeat("-", 73), "\n")

}

func (lc *LibraryController) AddBook() {
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter book title: ")
	title, _ := lc.reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := lc.reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{ID: 0, Title: title, Author: author, Status: "Available"}
	lc.service.AddBook(book)
	fmt.Println("Book added.")

	fmt.Println("\n", strings.Repeat("-", 73), "\n")

}

func (lc *LibraryController) RemoveBook() {
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter book ID to be Removed: ")
	bookIDs, _ := lc.reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDs))

	if !lc.service.RemoveBook(bookID) {
		fmt.Println("You can not remove borrowed book.")

	} else {
		fmt.Println("Book removed.")
	}

	fmt.Println("\n", strings.Repeat("-", 73), "\n")

}

func (lc *LibraryController) BorrowBook() {

	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := lc.reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := lc.reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	if err := lc.service.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error borrowing book:", err)
	} else {
		fmt.Println("Book borrowed.")
	}
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

}

func (lc *LibraryController) ReturnBook() {
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter book ID to return: ")
	bookIDStr, _ := lc.reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter member ID: ")
	memberIDStr, _ := lc.reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	if err := lc.service.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error returning book:", err)
	} else {
		fmt.Println("Book returned.")
	}

	fmt.Println("\n", strings.Repeat("-", 73), "\n")

}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.service.ListAvailableBooks()
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	if len(books) == 0 {
		fmt.Println(strings.Repeat(" ", 33), "NO BOOKS AVAILABLE!")
		fmt.Println("\n", strings.Repeat("-", 73), "\n")

	} else {

		fmt.Println("Available Books:")
		fmt.Println(strings.Repeat("-", 73))

		fmt.Printf("| %-5s | %-30s | %-15s | %-10s |\n", "ID", "Title", "Author", "Status")
		fmt.Println(strings.Repeat("-", 73))
		for _, book := range books {
			fmt.Printf("| %-5d | %-30s | %-15s | %-10s |\n", book.ID, book.Title, book.Author, book.Status)
		}
		fmt.Println(strings.Repeat("-", 73))
	}
}

func (lc *LibraryController) ListBorrowedBooks() {
	fmt.Println("\n", strings.Repeat("-", 73), "\n")

	fmt.Print("Enter member ID to list borrowed books: ")
	memberIDStr, _ := lc.reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	books, _ := lc.service.ListBorrowedBooks(memberID)

	if len(books) == 0 {
		fmt.Println("\n\n", strings.Repeat(" ", 33), "NO BOOKS AVAILABLE!")
		fmt.Println("\n", strings.Repeat("-", 73), "\n")

	} else {

		fmt.Println("Borrowed Books:")
		fmt.Println(strings.Repeat("-", 73))
		fmt.Printf("| %-5s | %-43s | %-15s |\n", "ID", "Title", "Author")
		fmt.Println(strings.Repeat("-", 73))

		for _, book := range books {
			fmt.Printf("| %-5d | %-43s | %-15s |\n", book.ID, book.Title, book.Author)
			fmt.Println(strings.Repeat("-", 73))
		}

	}

}
