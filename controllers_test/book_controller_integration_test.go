package controllers_test

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-mongo-app/controllers"
	"example.com/go-mongo-app/models"
	"example.com/go-mongo-app/repositories"
	"example.com/go-mongo-app/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	repo := repositories.NewBookRepository()
	service := services.NewBookService(repo)
	controller := controllers.NewBookController(service)

	r := mux.NewRouter()
	r.HandleFunc("/books", controller.CreateBook).Methods("POST")

	return r
}

func TestBookCRUDIntegration(t *testing.T) {
	router := setupRouter()

	// 1. Create a book
	book := models.Book{
		ISBN:   "123-456-789",
		Title:  "Test Book",
		Author: "Test Author",
	}
	body, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Get the ID of the created book
	bodyResp, _ := ioutil.ReadAll(rr.Body)
	var createdBook models.Book
	json.Unmarshal(bodyResp, &createdBook)
	assert.NotEmpty(t, createdBook.ID)

}
