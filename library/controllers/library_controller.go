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
	fmt.Println("\n\n*************************************************\n\n")

	fmt.Print("Enter the member's name: ")
	name, _ := lc.reader.ReadString('\n')
	name = strings.TrimSpace(name)

	member := models.Member{ID: 0, Name: name, BorrowedBooks: []models.Book{}}
	lc.service.AddMember(member)

	fmt.Println("Member added.")

	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) AddBook() {
	fmt.Println("\n\n*************************************************\n\n")

	fmt.Print("Enter book title: ")
	title, _ := lc.reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := lc.reader.ReadString('\n')
	author = strings.TrimSpace(author)

	book := models.Book{ID: 0, Title: title, Author: author, Status: "Available"}
	lc.service.AddBook(book)
	fmt.Println("Book added.")

	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) RemoveBook() {
	fmt.Println("\n\n*************************************************\n\n")

	fmt.Print("Enter book ID to be Removed: ")
	bookIDs, _ := lc.reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDs))

	if !lc.service.RemoveBook(bookID) {
		fmt.Println("You can not remove borrowed book.")

	} else {
		fmt.Println("Book removed.")
	}

	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) BorrowBook() {
	fmt.Println("\n\n*************************************************\n\n")

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
	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) ReturnBook() {
	fmt.Println("\n\n*************************************************\n\n")

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

	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) ListAvailableBooks() {
	books := lc.service.ListAvailableBooks()
	fmt.Println("\n\n*************************************************\n\n")
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s , **%s** \n \n", book.ID, book.Title, book.Author, book.Status)
	}
	fmt.Println("\n\n*************************************************\n\n")

}

func (lc *LibraryController) ListBorrowedBooks() {
	fmt.Println("\n\n*************************************************\n\n")

	fmt.Print("Enter member ID to list borrowed books: ")
	memberIDStr, _ := lc.reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	books, _ := lc.service.ListBorrowedBooks(memberID)

	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
	fmt.Println("\n\n*************************************************\n\n")

}
