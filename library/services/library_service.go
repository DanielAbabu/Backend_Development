package services

import (
	"errors"
	"library/models"
)

var (
	BookIDCounter   int = 1
	MemberIDCounter int = 1
)

type Library struct {
	books   map[int]models.Book
	members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

func (l *Library) AddMember(member models.Member) {
	member.ID = MemberIDCounter
	MemberIDCounter += 1
	l.members[member.ID] = member
}

func (l *Library) AddBook(book models.Book) {
	book.ID = BookIDCounter
	BookIDCounter += 1
	l.books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) bool {
	book, _ := l.books[bookID]
	if book.Status == "Borrowed" {
		return false
	}
	delete(l.books, bookID)
	return true
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exist := l.books[bookID]

	if !exist {
		return errors.New("Book doesn't exist!")
	}

	if book.Status == "Borrowed" {
		return errors.New("Book is not Available")
	}

	member, is := l.members[memberID]

	if !is {
		return errors.New("Member doesn't exist")
	}

	book.Status = "Borrowed"
	l.books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member

	return nil

}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exist := l.books[bookID]

	if !exist {
		return errors.New("Book doesn't exist!")
	}

	member, is := l.members[memberID]
	if !is {
		return errors.New("Member doesn't exist")
	}

	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book.Status = "Available"
			l.books[bookID] = book
			l.members[memberID] = member
			return nil
		}
	}
	return errors.New("Member did not borrow the book")
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}

	for _, book := range l.books {
		if book.Status == "Available" {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberId int) ([]models.Book, error) {
	member, exist := l.members[memberId]
	if !exist {
		return []models.Book{}, errors.New("Member doesn't exist")
	}

	return member.BorrowedBooks, nil
}
