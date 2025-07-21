package main

import (
    "encoding/json"
    "log"
    "os"
)

// book struct represents a book in the library.
type book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

// Global variables to hold the library of books and the file path.
var (
    libraryFile = "library.json"
    library     []book
)

// loadLibrary loads the library from a JSON file.
func loadLibrary() {
    _, err := os.Stat(libraryFile)
    if err == nil {
        data, err := os.ReadFile(libraryFile)
        if err != nil {
            log.Fatalf("Error reading library file! - %v", err)
        }
        json.Unmarshal(data, &library)
    }
}

// saveLibrary saves the current library to a JSON file.
func saveLibrary() {
    data, err := json.MarshalIndent(library, "", "  ")
    if err != nil {
        log.Fatalf("Error saving library file! - %v", err)
    }
    os.WriteFile(libraryFile, data, 0644)
}

// deleteBook removes a book from the library by its index.
func deleteBook(library []book, index int) ([]book, bool) {
    if index < 0 || index >= len(library) {
        return library, false
    }
    return append(library[:index], library[index+1:]...), true
}