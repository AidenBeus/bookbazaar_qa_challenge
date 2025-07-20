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
}

// This slice holds the list of books in the library.
var books = []book{
	{ID: "1", Title: "Warbreaker", Author: "Brandon Sanderson"},
	{ID: "2", Title: "Educated", Author: "Tara Westover"},
	{ID: "3", Title: "Ranger's Apprentice", Author: "John Flanagan"},
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

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

// main function initializes the Gin router and sets up the routes.
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookByID)
	router.POST("/books", postBook)
	router.DELETE("/delete/:id", deleteBook)
	router.Run("localhost:1001")
}
