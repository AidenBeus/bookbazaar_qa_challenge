package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

// This struct represents a book in the library system.
type book struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Price        float64 `json:"price"`
	NumAvailable int     `json:"num_available"`
}

// This slice holds the list of books in the library.
var books = []book{
	{ID: "1", Title: "Warbreaker", Author: "Brandon Sanderson", Price: 22, NumAvailable: 4},
	{ID: "2", Title: "Educated", Author: "Tara Westover", Price: 7.97, NumAvailable: 1},
	{ID: "3", Title: "Ranger's Apprentice", Author: "John Flanagan", Price: 5.32, NumAvailable: 3},
}

// getBooks handles the GET request to retrieve all books.
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// bookByID handles the GET request to retrieve a book by its ID.
func bookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// getBookByID retrieves a book by its ID from the slice of books.
func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// postBook handles the POST request to add a new book to the library.
func postBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// main function initializes the Gin router and sets up the routes.
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", postBook)
	router.Run("localhost:1001")
}
