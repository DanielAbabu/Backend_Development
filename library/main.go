package main

import (
	"fmt"
	"library/controllers"
	"library/services"
	"strings"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	for {
		fmt.Println("\n\n", strings.Repeat("=", 60), "\n")
		fmt.Println("                 Library Management System")
		fmt.Println("\n", strings.Repeat("=", 60), "\n")
		fmt.Println("  1. Add Member")
		fmt.Println("  2. Add Book")
		fmt.Println("  3. Remove Book")
		fmt.Println("  4. Borrow Book")
		fmt.Println("  5. Return Book")
		fmt.Println("  6. List Available Books")
		fmt.Println("  7. List Borrowed Books")
		fmt.Println("  8. Exit")
		fmt.Println(strings.Repeat("=", 60))
		fmt.Print("  Choose an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			controller.AddMember()
		case "2":
			controller.AddBook()
		case "3":
			controller.RemoveBook()
		case "4":
			controller.BorrowBook()
		case "5":
			controller.ReturnBook()
		case "6":
			controller.ListAvailableBooks()
		case "7":
			controller.ListBorrowedBooks()
		case "8":
			fmt.Println("\nGoodbye!")
			return
		default:
			fmt.Println("\nInvalid choice. Please try again.")
		}
	}
}
