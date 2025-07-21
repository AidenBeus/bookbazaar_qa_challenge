package main

import (
    "fmt"
    "github.com/rivo/tview"
    "encoding/json"
    "log"
    "os"
    "strconv"
)

type book struct {
    ID     string `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var (
    library     = []book{}
    libraryFile = "library.json"
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

func deleteBook(index int) {
    if index < 0 || index >= len(library) {
        fmt.Println("Invalid book index")
        return
    }
    library = append(library[:index], library[index+1:]...)
    saveLibrary()
}

func main() {
    app := tview.NewApplication()
    loadLibrary()
    libraryList := tview.NewTextView().
        SetDynamicColors(true).SetWordWrap(true)
    libraryList.SetBorder(true).SetTitle("Library Books")

    refreshList := func() {
        libraryList.Clear()
        if len(library) == 0 {
            fmt.Fprintln(libraryList, "No books in the library.")
            return
        }
        for i, book := range library {
            fmt.Fprintf(libraryList, "[%d] %s by %s\n", i+1, book.Title, book.Author)
        }
    }

    bookNameInput := tview.NewInputField().SetLabel("Book Name: ")
    bookAuthorInput := tview.NewInputField().SetLabel("Book Author: ")
    bookIDInput := tview.NewInputField().SetLabel("Book ID (for delete): ")

    form := tview.NewForm().
        AddFormItem(bookNameInput).
        AddFormItem(bookAuthorInput).
        AddFormItem(bookIDInput).
        AddButton("Add Book", func() {
            name := bookNameInput.GetText()
            author := bookAuthorInput.GetText()
            id := bookIDInput.GetText()
            if name == "" || author == "" || id == "" {
                fmt.Println("Please fill in all fields.")
                return
            }
            newBook := book{ID: id, Title: name, Author: author}
            library = append(library, newBook)
            saveLibrary()
            refreshList()
            bookNameInput.SetText("")
            bookAuthorInput.SetText("")
            bookIDInput.SetText("")
        }).
        AddButton("Delete Book", func() {
            indexStr := bookIDInput.GetText()
            index, err := strconv.Atoi(indexStr)
            if err != nil || index < 1 || index > len(library) {
                fmt.Println("Invalid book index")
                return
            }
            deleteBook(index - 1)
            refreshList()
            bookIDInput.SetText("")
        }).
        AddButton("Quit", func() {
            app.Stop()
        })

    form.SetBorder(true).SetTitle("Book Management").SetTitleAlign(tview.AlignLeft)

    flex := tview.NewFlex().
        AddItem(libraryList, 0, 1, false).
        AddItem(form, 0, 1, true)

    refreshList()
    if err := app.SetRoot(flex, true).Run(); err != nil {
        panic(err)
    }
}