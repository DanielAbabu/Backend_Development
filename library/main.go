package main

import (
	"fmt"
	"library/controllers"
	"library/services"
)

func main() {
	library := services.NewLibrary()
	controller := controllers.NewLibraryController(library)

	for {
		fmt.Println("Library Management System")
		fmt.Println("0. Add Member")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {

		case "0":
			controller.AddMember()

		case "1":
			controller.AddBook()

		case "2":
			controller.RemoveBook()

		case "3":
			controller.BorrowBook()

		case "4":
			controller.ReturnBook()

		case "5":
			controller.ListAvailableBooks()

		case "6":
			controller.ListBorrowedBooks()

		case "7":
			fmt.Println("Good bye")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
