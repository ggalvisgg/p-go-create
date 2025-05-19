package controllers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    //"fmt"

    //"github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/controllers"
)

// ---------------------- MOCK DEL SERVICIO ----------------------


type MockBookService struct {
    mock.Mock
}

func (m *MockBookService) AddBook(book models.Book) (*models.Book, error) {
    args := m.Called(book)
    return args.Get(0).(*models.Book), args.Error(1)
}

// ---------------------- TESTS ----------------------


func TestCreateBook_Success(t *testing.T) {
    mockService := new(MockBookService)
    controller := controllers.NewBookController(mockService)

    book := models.Book{
        Title:  "Go Programming",
        Author: "John Doe",
        ISBN:   "1234567890",
    }

    body, _ := json.Marshal(book)
    req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    mockService.On("AddBook", mock.MatchedBy(func(b models.Book) bool {
        return b.Title == book.Title && b.Author == book.Author && b.ISBN == book.ISBN
    })).Return(&models.Book{
        ID:     primitive.NewObjectID(),
        Title:  book.Title,
        Author: book.Author,
        ISBN:   book.ISBN,
    }, nil)

    controller.CreateBook(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
    mockService.AssertExpectations(t)
}

func TestCreateBook_InvalidData(t *testing.T) {
    mockService := new(MockBookService)
    controller := controllers.NewBookController(mockService)

    req := httptest.NewRequest("POST", "/books", bytes.NewReader([]byte("invalid json")))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    controller.CreateBook(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}
