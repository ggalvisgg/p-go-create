package services

import (
    "fmt"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

type BookServiceInterface interface {
    AddBook(book models.Book) (*models.Book, error)
}

type BookService struct {
    repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
    return &BookService{repo}
}

func (s *BookService) AddBook(book models.Book) (*models.Book, error) {
    return s.repo.CreateBook(book)
}
