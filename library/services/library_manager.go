package services

import "library/models"

type LibraryManager interface {
	AddMember(member models.Member)
	AddBook(book models.Book)
	RemoveBook(bookID int) bool
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}
