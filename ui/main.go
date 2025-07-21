package main
//used tview to create the UI for managing a book library

import (
    "fmt"
    "github.com/rivo/tview"
    "strconv"
)
// main function initializes the tview application and sets up the UI.
func main() {
    app := tview.NewApplication()
    loadLibrary()
    libraryList := tview.NewTextView().
        SetDynamicColors(true).SetWordWrap(true)
    libraryList.SetBorder(true).SetTitle("Library Books")
	// Function to refresh the library list display.
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
// Input fields for adding and deleting books and quitting the application.
    bookNameInput := tview.NewInputField().SetLabel("Book Name: ")
    bookAuthorInput := tview.NewInputField().SetLabel("Book Author: ")
    bookIDInput := tview.NewInputField().SetLabel("Book ID (delete only): ")

    form := tview.NewForm().
        AddFormItem(bookNameInput).
        AddFormItem(bookAuthorInput).
        AddFormItem(bookIDInput).
        AddButton("Add Book", func() { // This function adds a new book to the library.
            name := bookNameInput.GetText()
            author := bookAuthorInput.GetText()
            id := bookIDInput.GetText()
			// Check if all fields are filled before adding a new book.
            if name == "" || author == ""{
                fmt.Fprintln(libraryList, "Please fill in all fields.")
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
        AddButton("Delete Book", func() { // This function deletes a book from the library by its ID.
            indexStr := bookIDInput.GetText()
            index, err := strconv.Atoi(indexStr)
            if err != nil || index < 1 || index > len(library) {
                fmt.Fprintln(libraryList, "Invalid book index")
                return
            }
            library, _ = deleteBook(library, index - 1) // Convert to zero-based index
            saveLibrary()
            refreshList()
            bookIDInput.SetText("")
        }).
        AddButton("Quit", func() { // This function quits the application.	
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