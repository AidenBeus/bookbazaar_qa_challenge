package main

import (
    "encoding/json"
    "log"
    "os"
)

type book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var (
    libraryFile = "library.json"
    library     []book
)

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

func saveLibrary() {
    data, err := json.MarshalIndent(library, "", "  ")
    if err != nil {
        log.Fatalf("Error saving library file! - %v", err)
    }
    os.WriteFile(libraryFile, data, 0644)
}

func deleteBook(library []book, index int) ([]book, bool) {
    if index < 0 || index >= len(library) {
        return library, false
    }
    return append(library[:index], library[index+1:]...), true
}