package controllers

import (
    "encoding/json"
    "net/http"
    //"github.com/gorilla/mux"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
    //"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookController struct {
    Service services.BookServiceInterface
}

func NewBookController(service services.BookServiceInterface) *BookController {
    return &BookController{Service: service}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    newBook, err := c.Service.AddBook(book)
    if err != nil {
        http.Error(w, "Error al insertar libro", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newBook)
}
